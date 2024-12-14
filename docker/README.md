# Prerequisites

* Docker should be installed on your system.
  https://docs.docker.com/engine/install/

# Installation through Composer

1. Make a directory on your dedicated server:

```bash
mkdir -p /home/$USER/scstats
cd /home/$USER/scstats
```

1. Create a [.env.prod](https://github.com/Tomut0/SimpleClansStats/blob/master/docker/.env.prod). This file
   contains all application settings you need to work with. Feel free to change it.
2. Create a [nginx.conf](https://github.com/Tomut0/SimpleClansStats/blob/master/docker/nginx.conf). This is Nginx
   web config file.
3. Create a [compose.yaml](https://github.com/Tomut0/SimpleClansStats/blob/master/docker/compose.yaml). With this
   file Docker will create 3 services: SSStats fpm-server, Nginx and the MySQL Database.
4. Run the application:

```bash
docker compose up -d
```

Wait until all services are up and running:
```
docker compose ps
```

Browse to `http://{ip_or_domain_name}:8000` after started.
# Building locally
1. Clone from master branch or download from [release](https://github.com/Tomut0/SimpleClansStats/releases) page.
2. **Only if you know what you do**: 
make sure your BuildKit cache is empty and volumes are clear:
```shell
docker builder prune -af
docker compose -f .\docker\Dockerfile down -v
```
3. Then from the project's root:
```shell
docker compose -f .\docker\Dockerfile up -d --build --force-recreate
```
