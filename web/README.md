# SimpleClansStats web client

The SvelteKit client uses the Go server for its API. During development, run `go run ./cmd --static` from the repository root and `npm ci` followed by `npm run dev` from `web/`, in separate terminals. Omit `--static` if another server, such as Nginx, serves the web client. If applying a preset changes npm dependencies while Vite is running, restart Vite.

## Themes

Applying a preset from **Settings → Theme** runs the full `shadcn-svelte apply <preset>` command (including dependency installation) and then `npm run build`. The full preset updates theme variables, fonts, `components.json`, and installed shadcn-svelte components. The settings page reloads after the build so the browser fetches the new production bundle. Make sure the web directory has its npm dependencies installed before running the Go server locally.

The Docker image installs the web dependencies, builds the initial static client, and starts with `--static`. To let Nginx serve the web client instead, override the container command with `--static=false`; Nginx must serve the same build directory if it should pick up theme rebuilds.

Build the static client manually with `npm run build`.
