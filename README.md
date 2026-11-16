# SimpleClansStats

A web dashboard for viewing and exploring [SimpleClans](https://github.com/RoinujNosde/SimpleClans) data across one or more Minecraft servers written in Golang and SvelteKit.

Use the dashboard to monitor clan rosters, wars, balances, kills, online players, and recent activity. Search for a clan or player to view detailed information, or customize your dashboard widgets to focus on what matters most to your server. Manage multiple servers from a single installation.

The interface is available in English and Russian (contributions for additional translations are welcome!).

[![Docker build](https://github.com/Tomut0/SimpleClansStats/actions/workflows/docker-publish.yml/badge.svg?branch=master)](https://github.com/Tomut0/SimpleClansStats/actions/workflows/docker-publish.yml) [![License](https://img.shields.io/github/license/Tomut0/SimpleClansStats)](LICENSE) [![Latest tag](https://img.shields.io/github/v/tag/Tomut0/SimpleClansStats?label=version)](https://github.com/Tomut0/SimpleClansStats/tags)

## Screenshots

| Dashboard | Add widget | Edit dashboard |
| --- | --- | --- |
| <img src="docs/images/dashboard.png" width="300" alt="Dashboard"> | <img src="docs/images/add-widget.png" width="300" alt="Add widget dialog"> | <img src="docs/images/edit-dashboard.png" width="300" alt="Edit dashboard"> |

| Player details | Settings | Settings |
| --- | --- | --- |
| <img src="docs/images/player-modal.png" width="300" alt="Player details dialog"> | <img src="docs/images/settings-1.png" width="300" alt="Settings page"> | <img src="docs/images/settings-2.png" width="300" alt="Settings page, continued"> |

## Features

- Customizable Dashboard: Administrators can add, remove and reposition widgets to create their preferred dashboard layout.
- Real-time Updates: Server data and activity are updated in real time, keeping your dashboard up to date.
- Clan & Player Search: Find clans and players and explore their detailed statistics and information.
- Multilingual Support: Let users choose their preferred language or enforce a single application-wide language.
- Multi-Server Support: Connect and monitor multiple Minecraft servers from a single installation.
- Theme Customization: Choose from different themes powered by shadcn-svelte.
- Custom Branding: Personalize your installation with a custom application name and logo.
- Light & Dark Mode: Switch between light and dark appearances.
- API & Web Modes: Run the application with its built-in web interface or in API-only mode to build your own frontend.

## 📦 How to Install

### Standalone
For the very basic installation:
1. Download the binary file from [latest release](https://github.com/Tomut0/SimpleClansStats/releases/latest).
```bash
curl -fL "https://github.com/Tomut0/SimpleClansStats/releases/latest/download/simpleclansstats-linux-$( [[ "$(uname -m)" == "x86_64" ]] && echo "amd64" || echo "arm64" )" -o scstats
```
2. Make it executable and create a persistent folder:
```bash
chmod +x scstats
mkdir -p data
```
3. Copy [.env.example](https://github.com/Tomut0/SimpleClansStats/blob/master/.env.example) and rename it to `.env`, open the file and set variables.
4. Run! :)
```
./scstats --mode=web
```

Now visit `http://localhost:4000/` to see your web dashboard.

### Docker
```bash
docker run -d --name simpleclansstats `
  -p 4000:4000 `
  -v simpleclansstats-data:/data `
  -e BOOTSTRAP_ADMIN_EMAIL=admin@example.com `
  -e BOOTSTRAP_ADMIN_USERNAME=admin `
  -e BOOTSTRAP_ADMIN_PASSWORD='admin' `
  -e CORS_ALLOWED_ORIGINS=http://localhost:4000 `
  -e APP_URL=https://stats.example.com `
  -e APP_VERSION=2.0.0 `
  -e POSTHOG_API_KEY=phc_wboArzoxaHWqADcRiazJLa4otEs7KViZNgvzzsjMsdXX `
  -e POSTHOG_HOST=https://us.i.posthog.com `
  docker.io/minat0/scstats:latest
```

Now visit `http://localhost:4000/` to see your web dashboard.

## Support

If you have faced any issues, I'll answer on any questions in official
SimpleClans [discord](https://discord.gg/CkNwgdE).

## ✨ Sponsored By

<img align="left" width=125 style="margin-right: 12px" src="https://mikasa.host/images/favicon.png" alt="image">

### mikasa.host

It is my game hosting, the perfect solution for launching your own Minecraft server. \
For just $1 in a month, you can create a powerful server featuring top-tier components like Ryzen 9 5950X
processors, DDR4 RAM, and NVMe disks. Mikasa.host offers convenient and speedy registration via Discord and supports hosting servers for a variety of games, not just Minecraft.

<br clear="left">

## 😎 Thank you especially 

<a href="https://github.com/RoinujNosde"><img src="https://github.com/RoinujNosde.png" width="50px" alt="RoinujNosde" /></a>
<a href="https://github.com/ToxicDuke"><img src="https://github.com/ToxicDuke.png" width="50px" alt="ToxicDuke" /></a>
<a href="https://github.com/mehmet-27"><img src="https://github.com/mehmet-27.png" width="50px" alt="mehmet-27" /></a>

## License

Licensed under the [Apache License 2.0](LICENSE).
