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

### Nginx & Systemd

Needed tools:

1. Docker
2. Nginx

Configuration is static (may be improved in the future). Used ports:

1. 8851 - for server part (run as system service)
2. 8803 - listening port for Nginx

To install Werminal use script `./scripts/install.sh`. Run it as root.

## Used tools

1. `Golang` (server) with `pty` package
2. `JS` (client) with `xterm.js` library
3. `WebSocket` for client-server connection
4. `Docker compose` for app deployment
