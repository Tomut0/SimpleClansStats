package main

import (
	"context"
	"log"
	"time"

	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
)

func startIngestCleanup(queries *generated.Queries) {
	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		for {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			if err := cleanupIngest(ctx, queries, time.Now()); err != nil {
				log.Printf("purge ingest history: %v", err)
			}
			cancel()
			<-ticker.C
		}
	}()
}

func cleanupIngest(ctx context.Context, queries *generated.Queries, now time.Time) error {
	days, err := queries.GetIngestRetentionDays(ctx)
	if err != nil || days == 0 {
		return err
	}
	cutoff := sqlite.FormatIngestTime(now.UTC().AddDate(0, 0, -int(days)))
	if err := queries.PurgeIngestSnapshots(ctx, cutoff); err != nil {
		return err
	}
	return queries.PurgeIngestEvents(ctx, cutoff)
}
