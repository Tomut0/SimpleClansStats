# SimpleClansStats

<img src="https://mikasa.host/images/scstats.png" alt="image">

Laravel based web interface for SimpleClans plugin. \
Laravel version: 10.15.0

Watch demo of the application [here](https://scstats.roinujnosde.me/).

* Laravel [documentation](https://laravel.com/docs)
* SimpleClans [plugin](https://github.com/RoinujNosde/SimpleClans)
---
## Features

* **Powerful Leaderboard**: \
  Easily see which players have the best combat performance with sorting options for KDR (Kill-Death Ratio), Balance,
  and Member count.
* **Updates clan positions daily, weekly, and monthly**: \
  Track your progress and see how your clan ranks over different time periods.
* **Multi Language Support**: \
  Seamless language switching without page reload. \
  Instantly translate the page into supported languages for a smoother user experience. \
  Currently supported languages are: English, Portuguese (Brazil), Russian.
* **Comprehensive Statistics**: \
  Displays latest kills between clan players, provides detailed kills by types, shows average clan balance, \
  and tracks the total number of clans on the server.
* **Extremely Fast and Configurable**: \
  Built with Vue and Inertia for a smooth and dynamic user experience. \
  Easily deployable with Docker, allowing you to quickly set up and configure your environment to meet your needs.
* **Adaptive Design**: \
  Responsive layout that adjusts seamlessly to different screen sizes and devices, \
  ensuring an optimal viewing experience on desktops and mobile phones.
## Support

If you have faced any issues, I'll answer on any questions in official
SimpleClans [discord](https://discord.gg/CkNwgdE).

## 🚀 Getting Started

### Docker

To run SimpleClansStats under Docker, you have to use `minat0/scstats:latest` docker image in your compose file. \
Example is located [here](https://github.com/Tomut0/SimpleClansStats/tree/master/docker).

### Web server

#### Prerequisites

* Any http server. Most popular are Nginx or Apache.
* PHP v8.1 or higher
* Composer v2.7.2 or higher
* Node.js v20.12.2 or higher
* npm v10.8.2 or higher

#### Step by step

1. Clone from master branch or download from [release](https://github.com/Tomut0/SimpleClansStats/releases) page.
2. Install php dependencies through composer:

```bash
composer install --no-dev --optimize-autoloader
```

3. Install and build node.js dependencies through npm:

```bash
npm install
npm run build
```

4. Make copy of `.env.example` to `.env` and set up it.

```bash
cp .env.example .env
php artisan key:generate
```

You have to change `DB_`, `SC_` settings to connect to the databases. \
Make sure `SC_` database has `sc_players`, `sc_clans`, `sc_kills` tables to work finely.

5. Add task scheduling rule to crontab.

```bash 
crontab -e 
* * * * * cd /path-to-your-project && php artisan schedule:run >> /dev/null 2>&1
```

6. Configure your http server to process the SCStats. \
   You may look an example of Nginx
   configuration [here](https://github.com/Tomut0/SimpleClansStats/blob/master/docker/nginx.conf). \
   Don't forget to change `fastcgi_pass scstats:9000;` to your php-fpm server.

7. Profit!

## ✨ Sponsored By

<img align="left" width=125 style="margin-right: 12px" src="https://mikasa.host/images/favicon.png" alt="image">

### mikasa.host

It is my game hosting, the perfect solution for launching your own Minecraft server. \
For just $1 in a month, you can create a powerful server featuring top-tier components like Ryzen 9 5950X
processors, DDR4 RAM, and NVMe disks. Mikasa.host offers convenient and speedy registration via Discord and supports
hosting servers for a variety of games, not just Minecraft.

## 😎 Thank you especially 

<a href="https://github.com/RoinujNosde"><img src="https://github.com/RoinujNosde.png" width="50px" alt="RoinujNosde" /></a>
<a href="https://github.com/ToxicDuke"><img src="https://github.com/ToxicDuke.png" width="50px" alt="ToxicDuke" /></a>
<a href="https://github.com/mehmet-27"><img src="https://github.com/mehmet-27.png" width="50px" alt="mehmet-27" /></a>

## License

This project is under [Apache2](https://github.com/Tomut0/SimpleClansStats/blob/master/LICENSE) license.
