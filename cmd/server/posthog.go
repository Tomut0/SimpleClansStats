package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type posthogTelemetry struct {
	key, host, version, appURL string
	installationID             string
	client                     *http.Client
	mu                         sync.Mutex
	lastSent                   time.Time
}

func newPosthogTelemetry(installationID string) *posthogTelemetry {
	key := os.Getenv("POSTHOG_API_KEY")
	if key == "" || installationID == "" {
		return nil
	}
	host := strings.TrimRight(os.Getenv("POSTHOG_HOST"), "/")
	if host == "" {
		host = "https://us.i.posthog.com"
	}
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "unknown"
	}
	return &posthogTelemetry{key: key, host: host, version: version, appURL: strings.TrimRight(os.Getenv("APP_URL"), "/"), installationID: installationID, client: &http.Client{Timeout: 5 * time.Second}}
}

func (p *posthogTelemetry) middleware(next http.Handler, allowed func(*http.Request) bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowed(r) {
			p.capture(r)
		}
		next.ServeHTTP(w, r)
	})
}

func (p *posthogTelemetry) capture(r *http.Request) {
	appURL := p.appURL
	if appURL == "" {
		appURL = r.Header.Get("Origin")
	}
	if appURL == "" {
		scheme := r.Header.Get("X-Forwarded-Proto")
		if scheme != "https" {
			scheme = "http"
		}
		appURL = scheme + "://" + r.Host
	}
	now := time.Now().UTC()
	p.mu.Lock()
	// shortcut: Sends at most one heartbeat per five minutes, sufficient for last-activity reporting; reduce the interval if finer activity history is needed.
	if now.Sub(p.lastSent) < 5*time.Minute {
		p.mu.Unlock()
		return
	}
	p.lastSent = now
	p.mu.Unlock()

	go func() {
		payload, err := json.Marshal(map[string]any{
			"api_key": p.key, "event": "app_activity", "distinct_id": p.installationID,
			"properties": map[string]any{"app_url": appURL, "app_version": p.version, "last_activity": now.Format(time.RFC3339), "$process_person_profile": false},
		})
		if err != nil {
			return
		}
		request, err := http.NewRequest(http.MethodPost, p.host+"/capture/", bytes.NewReader(payload))
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/json")
		response, err := p.client.Do(request)
		if err == nil {
			response.Body.Close()
		}
	}()
}
