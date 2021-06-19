# SimpleClansStats

![изображение](https://user-images.githubusercontent.com/42711312/122625148-437ab680-d0ac-11eb-863f-1d3b9d98ac42.png)

Laravel based web interface for SimpleClans plugin. <br>
Laravel version: 6.20.28

* Laravel [documentation](https://laravel.com/docs)
* SimpleClans [plugin](https://github.com/RoinujNosde/SimpleClans)
* The project idea is taken directly from [SimpleClansStats2](https://github.com/lpostiglione/SimpleClansStats2)

## Getting Started
### Web server
To install SimpleClansStats interface, please follow the next simple steps:
1. Download the lateast release from [this link](https://github.com/Tomut0/SimpleClansStats/releases)
2. Install [PHP](https://www.php.net/downloads) >= 7.4
3. Make sure you have [NGINX](https://www.nginx.com/) or [Apache](https://httpd.apache.org/) as one of your web server.
4. Unpack zip archieve with any ZIP explorer program (such as [WinRar](https://www.win-rar.com/download.html?&L=0) or [7-Zip](https://www.7-zip.org/download.html)) to web server's home folder.
5. Run `composer install`
5. Make copy of `.env.example` to `.env` and set up it.
```shell
cp .env.example .env
php artisan key:generate
```
> **Note** <br>
> The database is mandatory and must have tables `sc_players` and `sc_clans` to work fine.
6. Profit! 😃
### Local
WIP.
## Useful links
1. Using [Laravel mix](https://github.com/JeffreyWay/laravel-mix/) to compile SCSS (not only)
2. [SCSS documentation](https://sass-scss.ru/documentation/)
## License

This project is under [Apache2 licence](https://github.com/Tomut0/SimpleClansStats/blob/master/LICENSE).

[![Crowdin](https://badges.crowdin.net/simpleclansstats/localized.svg)](https://crowdin.com/project/simpleclansstats)
