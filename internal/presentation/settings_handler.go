package presentation

import (
	"bytes"
	"context"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/Tomut0/SimpleClansStats/internal/application"
	"github.com/Tomut0/SimpleClansStats/internal/domain"
)

// 1 MB
const maxLogoBytes = 1 << 20
const MaxSettingsRequestBytes = maxLogoBytes + 4096

func (s *ServerHandler) GetAppSettings(ctx context.Context, _ GetAppSettingsRequestObject) (GetAppSettingsResponseObject, error) {
	settings, err := s.settingsService.Get(ctx)
	if err != nil {
		log.Printf("read app settings: %v", err)
		return GetAppSettings500Response{}, nil
	}
	return GetAppSettings200JSONResponse(appSettingsResponse(settings)), nil
}

func (s *ServerHandler) UpdateAppSettings(ctx context.Context, request UpdateAppSettingsRequestObject) (UpdateAppSettingsResponseObject, error) {
	switch s.adminStatus(ctx) {
	case http.StatusUnauthorized:
		return UpdateAppSettings401Response{}, nil
	case http.StatusForbidden:
		return UpdateAppSettings403Response{}, nil
	}

	form, err := request.Body.ReadForm(MaxSettingsRequestBytes)
	if err != nil {
		var limit *http.MaxBytesError
		if errors.As(err, &limit) || errors.Is(err, multipart.ErrMessageTooLarge) {
			return UpdateAppSettings413Response{}, nil
		}
		return UpdateAppSettings400Response{}, nil
	}
	defer form.RemoveAll()

	params, ok := parseAppSettings(form)
	if !ok {
		return UpdateAppSettings400Response{}, nil
	}
	logos := form.File["logo"]
	if len(logos) > 1 || len(form.Value["logo"]) != 0 {
		return UpdateAppSettings400Response{}, nil
	}
	if len(logos) == 1 {
		if logos[0].Size > maxLogoBytes {
			return UpdateAppSettings413Response{}, nil
		}
		params.Logo, err = readLogo(logos[0])
		if err != nil {
			return UpdateAppSettings400Response{}, nil
		}
	}

	saved, err := s.settingsService.Update(ctx, params)
	if err != nil {
		if errors.Is(err, application.ErrInvalidSettings) {
			return UpdateAppSettings400Response{}, nil
		}
		log.Printf("save app settings: %v", err)
		return UpdateAppSettings500Response{}, nil
	}
	return UpdateAppSettings200JSONResponse(appSettingsResponse(saved)), nil
}

func (s *ServerHandler) GetAppLogo(ctx context.Context, _ GetAppLogoRequestObject) (GetAppLogoResponseObject, error) {
	logo, err := s.settingsService.Logo(ctx)
	if err != nil {
		log.Printf("read app logo: %v", err)
		return GetAppLogo500Response{}, nil
	}
	switch http.DetectContentType(logo) {
	case "image/png":
		return GetAppLogo200ImagepngResponse{Body: bytes.NewReader(logo), ContentLength: int64(len(logo))}, nil
	case "image/jpeg":
		return GetAppLogo200ImagejpegResponse{Body: bytes.NewReader(logo), ContentLength: int64(len(logo))}, nil
	default:
		return GetAppLogo404Response{}, nil
	}
}

func parseAppSettings(form *multipart.Form) (domain.AppSettingsUpdate, bool) {
	var params domain.AppSettingsUpdate
	for _, name := range []string{"appName", "removeAuthor", "timeZone", "dateFormat", "timeFormat", "playerAvatarUrl"} {
		if len(form.Value[name]) != 1 {
			return params, false
		}
	}
	params.AppName = form.Value["appName"][0]
	params.TimeZone = form.Value["timeZone"][0]
	params.DateFormat = form.Value["dateFormat"][0]
	params.TimeFormat = form.Value["timeFormat"][0]
	params.PlayerAvatarURL = form.Value["playerAvatarUrl"][0]
	removeAuthor := form.Value["removeAuthor"][0]
	params.RemoveAuthor = removeAuthor == "true"
	var ok bool
	params.ServerOnlineTimeoutMinutes, ok = optionalSettingsInt(form.Value["serverOnlineTimeoutMinutes"])
	if !ok {
		return params, false
	}
	params.IngestRetentionDays, ok = optionalSettingsInt(form.Value["ingestRetentionDays"])
	if !ok {
		return params, false
	}

	if removeAuthor != "true" && removeAuthor != "false" {
		return params, false
	}
	return params, true
}

func optionalSettingsInt(values []string) (*int64, bool) {
	if len(values) == 0 {
		return nil, true
	}
	if len(values) != 1 {
		return nil, false
	}
	value, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return nil, false
	}
	return &value, true
}

func readLogo(header *multipart.FileHeader) ([]byte, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	logo, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	config, format, err := image.DecodeConfig(bytes.NewReader(logo))
	if err != nil || (format != "png" && format != "jpeg") ||
		config.Width < 1 || config.Width > 2048 || config.Height < 1 || config.Height > 2048 {
		return nil, errors.New("invalid logo format or dimensions")
	}
	if _, _, err := image.Decode(bytes.NewReader(logo)); err != nil {
		return nil, err
	}
	return logo, nil
}

func appSettingsResponse(row domain.AppSettings) AppSettings {
	return AppSettings{
		AppName:                    row.AppName,
		RemoveAuthor:               row.RemoveAuthor,
		TimeZone:                   row.TimeZone,
		DateFormat:                 AppSettingsDateFormat(row.DateFormat),
		TimeFormat:                 AppSettingsTimeFormat(row.TimeFormat),
		PlayerAvatarUrl:            row.PlayerAvatarURL,
		LogoVersion:                row.LogoVersion,
		ServerOnlineTimeoutMinutes: row.ServerOnlineTimeoutMinutes,
		IngestRetentionDays:        row.IngestRetentionDays,
	}
}
