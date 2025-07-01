# Werminal

Terminal in web-interface (for linux).

## `ENV` variables for server that may be specified

```dotenv
SERVER_PORT=8080
SERVER_SHUTDOWN_TIMEOUT=5s
```

## Deploy

### Docker

```shell
cd ./Werminal/deployment
docker compose up --build
```

### Nginx & Server process

#### Preparation

> You should set up user/passwd for nginx `basic auth` before app installing.
>
> Follow the next guide

##### Install tool for password hash generating

```shell
sudo apt install apache2-utils
```

##### Create password for user (user1 in the samople below)

```shell
htpasswd -с /etc/nginx/.htpasswd user1
```

#### App Installation

Needed tools:

1. Docker
2. Nginx

Configuration is static. Used ports:

1. 8851 - for server part (run as process)
2. 8092 - listening port for Nginx

To install Werminal use script `./scripts/install.sh`. Run it as root.
To uninstall Werminal use script `./scripts/uninstall.sh`. Run it as root.

## Used tools

1. `Golang` (server) with `pty` package
2. `JS` (client) with `xterm.js` library
3. `WebSocket` for client-server connection
4. `Docker` for app deployment
