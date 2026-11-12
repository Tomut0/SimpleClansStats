# Repository-wide code simplification audit

## Repository Coverage

SimpleClansStats accepts collector snapshots and serves a customizable dashboard. The maintained application is a Go HTTP API with SQLite persistence and a SvelteKit frontend. The audit covers collector ingestion, browser sessions, and dashboard refresh and layout updates.

- Reviewed for discovery. Tracked source inventory, dependency manifests, root and frontend READMEs, build and deployment entrypoints, and code generation configuration.
- Reviewed Go startup, transport, application, domain and ports, SQLite adapters, password hashing, unit and HTTP integration tests, and SQL migrations and queries. Traced session, settings, ingest, retention, and dashboard flows across those layers.
- Reviewed SvelteKit routes, shared state, hooks, global layout and styles, application components, settings forms, dashboard grid and widgets, and Playwright specs. Checked English and Russian catalogs for matching leaf keys (160 each). Surveyed all shared UI primitive families by imports, exports and control flow, then inspected chart, select, button, input and field behavior in detail.
- Reviewed the active OpenAPI contract, migration and query sources, and Go and web generation, test, environment and deployment configuration. Compared the raw OpenAPI draft's operation and schema inventory with the active contract and checked for local references to the draft.
- Remaining review gaps. Individual constraints in the raw OpenAPI draft, class and prop-forwarding details in the thin shared UI wrappers, translation wording, live deployment and browser behavior.
- Excluded. Generated Go and TypeScript, dependencies and lock/checksum contents, local environment values and databases, build output, binary assets, and Git internals.
- The codebase-memory CLI declined the local coordination endpoint because of its Windows ACL. Coverage will use direct source inspection and reference searches.

## Findings

### Go HTTP and application

**[S1] Reuse the admin check across write handlers**

- **Priority:** Medium
- **Location:** `internal/presentation/servers_handler.go` (`serverAdminStatus`, `CreateServer`, `DeleteServer`), `internal/presentation/settings_handler.go` (`UpdateAppSettings`).
- **Problem:** Settings and server writes perform the same session and role lookup independently. Any change to the policy has to be kept in sync at both sites.
- **Change:** Rename `serverAdminStatus` to `adminStatus` and call it from all three handlers. Keep each handler's switch over its own generated response types.
- **Preserve:** An absent or incorrectly typed session user ID and a lookup error yield 401. A recognized non-admin yields 403. Only admins proceed, with authorization still preceding settings form reads and service calls.
- **Status:** Pending

**[S2] Separate widget options from layout geometry**

- **Priority:** Medium
- **Location:** `internal/application/dashboard_service.go` (`SaveLayout` and the widget ID and field tables).
- **Problem:** One loop mixes three ID rules, type-specific option checks, bounds, and duplicate detection. The untyped branch repeats a type check that its enclosing condition already guarantees.
- **Change:** Move per-widget ID, field, format and limit checks into one validator. Handle legacy untyped IDs first, then check the typed field allowlist, counter IDs and formats, and table limits with separate guards. Keep layout version, count, geometry, uniqueness and the repository call in `SaveLayout`.
- **Preserve:** Untyped legacy widgets require empty field and format and currently accept any limit. Typed counters require a valid format and either a matching legacy counter ID and field or the current counter pattern. Other typed widgets require the current widget pattern and allow empty or recognized formats. Tables accept limits 0, 5, 10 and 20. Other typed widgets accept only 0. Keep the exact regexes, allowlists, bounds, error sentinel and saved layout.
- **Status:** Pending

### SQLite persistence and retention

**[S3] Share the timestamp format used for SQLite comparisons**

- **Priority:** Medium
- **Location:** `internal/infrastructure/persistence/sqlite/ingest_repository.go` (`SaveIngest`, `takenIngestTime`), `cmd/ingest_cleanup.go` (`cleanupIngest`), `cmd/server/ingest_test.go` (retention cutoff), `sql/queries/ingest_snapshots.sql`.
- **Problem:** Inserts, ordering checks and retention depend on an identical fixed-width UTC encoding. Cleanup and ingest own separate copies of the encoding, and the two write queries each reformat the same timestamp and hash.
- **Change:** Expose the SQLite adapter's formatter as `sqlite.FormatIngestTime`. Reuse it for the cleanup cutoff and the integration test's computed cutoff. In `SaveIngest`, compute the formatted taken time and hex state hash once for both write queries. Keep fixed timestamp fixtures independent.
- **Preserve:** UTC with exactly nine fractional digits and a literal Z. Keep calendar-day subtraction in UTC, retention zero as disabled, strict cutoff comparison, and transaction and write order.
- **Status:** Pending

**[S4] Remove unused repository surface**

- **Priority:** Low
- **Location:** `internal/domain/ports/clan_repository.go` (`ClanRepository`), `internal/domain/ports/user_repository.go` (`All`), `internal/infrastructure/persistence/sqlite/user_repository.go` (`All`), `sql/queries/users.sql` (`ListUsers`), `internal/infrastructure/persistence/sqlite/sqlite.go` (`SQLite.Ctx`).
- **Problem:** The empty clan port, user enumeration path and stored background context have no consumers. The unused user query also produces a generated method.
- **Change:** Delete the empty port, the user enumeration method from the port and adapter, and its SQL query. Regenerate sqlc output. Remove `SQLite.Ctx` and its initialization.
- **Preserve:** Keep the authentication and bootstrap lookup methods, user row validation, schema, SQLite DSN, connection lifecycle and generated query construction.
- **Status:** Pending

### Dashboard UI

**[S5] Render row-based widgets directly with the existing table**

- **Priority:** Medium
- **Location:** `web/src/lib/components/Widgets/DashboardGrid.svelte` (`tableRows` and widget render branches), `RowsWidget.svelte`, `WarCenterWidget.svelte`, `ClanRosterWidget.svelte`, `PendingInvitesWidget.svelte`, `GainersLosersWidget.svelte`, and `ClaimsWidget.svelte` in the same directory.
- **Problem:** Five components forward props to `RowsWidget`. Four repeat the same row and column selection already present in `tableRows`, so editing a column can leave the configurable and persisted widgets out of sync.
- **Change:** Render `RowsWidget` in the five existing branches. Share the four matching row and column definitions with `tableRows`. Keep claims as a separate definition or render branch, and retain `tableRows`'s current fallback to balance changes. Remove the five forwarding components and imports.
- **Preserve:** Keep all five persisted type names, current titles, field order, row order, empty state and widget header. Keep the row cap on the configurable table alone. Leave catalog visibility, layout validation and persistence as they are.
- **Status:** Pending

**[S6] Reuse fixed widget titles**

- **Priority:** Low
- **Location:** `web/src/lib/components/Widgets/DashboardGrid.svelte` (`configuredTitle`, `configuredTypeTitle`, `configuredFieldTitle`).
- **Problem:** The same title translation is selected independently for the add catalog, rendered titles and remove-button labels. Several branches in the field title function are shadowed by its counter-field check.
- **Change:** Let `configuredTypeTitle` supply fixed titles in `configuredTitle` for peak online, war center, clan roster, economy, kill feed, pending invites, claims, clan card, player card, network, compare, heatmap and gainers/losers. Keep the existing specialized title branches and untyped fallback. Remove the unreachable `onlinePlayers` and `wars` branches from `configuredFieldTitle`.
- **Preserve:** Keep exact translation keys and when they are evaluated, the legacy-ID fallback, and field-dependent counter, table, distribution and time-series labels. In particular, the table field `wars` still resolves through the counter label.
- **Status:** Pending

### Shared UI

**[S7] Give both input branches one source for styling**

- **Priority:** Low
- **Location:** `web/src/lib/components/ui/input/input.svelte` (file and general input branches), `web/src/lib/components/LogoPreview.svelte` (file caller), `web/src/lib/components/LoginForm.svelte` and settings cards (general callers).
- **Problem:** The complete class merge is duplicated between branches. A focus, disabled or invalid style change has to be made twice.
- **Change:** Compute the merged class once with a reactive derived value and pass it to both existing input branches.
- **Preserve:** Keep both branches, their bindings and attribute order, including file-only `bind:files`, `bind:value`, the forwarded ref and the caller's reactive class overrides.
- **Status:** Pending

### Dashboard data

**[S8] Keep only the decoded snapshots used for balance deltas**

- **Priority:** Medium
- **Location:** `internal/application/dashboard_service.go` (`DashboardService.Data`), `internal/infrastructure/persistence/sqlite/dashboard_repository.go` (`GetDashboardData`).
- **Problem:** The history loop builds a second slice and retains up to 48 decoded collector maps. Only the two newest decoded maps are used after that loop.
- **Change:** Keep the existing reverse traversal and append history as each snapshot is decoded. Retain only the maps at source indexes 0 and 1, then use those for the delta if at least two snapshots were present. Remove the decoded slice.
- **Preserve:** Keep oldest-to-newest history order, first decode error and partial data returned on that error, latest-versus-previous delta order, nil snapshots and empty delta array on successful return where applicable. Use the number of snapshots for the delta guard because JSON null decodes into a nil map.
- **Status:** Pending

## Implementation Order

1. Apply S1 with the generated response types in place. Then apply S2 and S8 together in the dashboard service, where the two changes touch independent methods.
2. Apply S3 to the ingest and cleanup paths. Apply S4 and regenerate sqlc output from the updated query sources.
3. Apply S5 and S6 together in the dashboard grid, then S7 in the shared input. Run the browser checks after that batch.

## Verification

- The audit baseline passed with Go 1.27.1 using `go test -mod=readonly ./...` from the repository root. Rerun it after implementation. For S1, see `TestAppSettings` and `TestServerRegistrationAndIngest` for admin and non-admin requests. Also check that a failed user lookup yields 401 on each write endpoint.
- For S2, run `go test ./internal/application -run TestSaveDashboardLayout` and check omitted type, legacy limit, and typed table limit cases against the current behavior. For S8, run the dashboard history and balance-change tests and check three snapshots in newest-first input order, a bad JSON record midway through history, and a single snapshot.
- For S3, run `go test ./cmd/... -run 'TestCleanupUsesStoredRetention|TestServerRegistrationAndIngest'`. Keep the existing literal UTC assertion and verify a record exactly at the cutoff survives.
- For S4, run `go tool sqlc generate` from the root, verify generated changes match the removed query, and rerun the full Go tests.
- For S5 and S6, from `web` run `npm run check` and `npx playwright test dashboard.e2e.ts`. The existing spec covers widget add/remove, layout save/reload and live updates. Also load saved layouts with all five table-backed types and compare their headings, columns and row counts. For S7, check login inputs and the settings logo chooser, including keyboard focus, override classes, and selecting a replacement file.
- Web checks are implementation verification. The audit did not run them because the configured Svelte and Playwright workflows write generated files and test artifacts into the working tree.

## Implementation Progress

| Finding ID | Status | Notes |
| --- | --- | --- |
| S1 | Completed | Shared admin authorization across settings and server writes; full Go suite covers missing session, non-admin, and injected lookup failures on all three write endpoints. |
| S2 | Completed | Extracted widget option validation; layout tests cover configured widgets, legacy untyped limits, and accepted/rejected table limits. |
| S3 | Completed | Shared UTC formatter and reused formatted values; cleanup test confirms a record exactly at the strict cutoff survives, and ingest test retains its literal timestamp assertion. |
| S4 | Completed | Removed unused port and user enumeration query/method, removed SQLite context, and regenerated sqlc output; full Go suite passes. |
| S5 | Skipped | Reused RowsWidget and shared row definitions for all five persisted types; npm check passes. Playwright/browser layout checks could not be completed because its configured Vite server never became reachable on 127.0.0.1:4173. |
| S6 | Completed | Fixed titles now reuse configuredTypeTitle; npm check passes. |
| S7 | Skipped | Input styling now uses one reactive derived class; npm check passes. Playwright checks for file selection, focus, and overrides could not run because its configured Vite server never became reachable on 127.0.0.1:4173. |
| S8 | Completed | Retained only the newest two decoded snapshots for deltas; tests cover newest-first input, history order, partial history on bad JSON, and one null snapshot. |
