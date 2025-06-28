# Werminal

Terminal in web-interface (for linux).

## `ENV` variables for server that may be specified

```dotenv
SERVER_PORT=8080
SERVER_SHUTDOWN_TIMEOUT=5s
```

## Deploy

```shell
cd ./Werminal/deployment
docker compose up --build
```

## Used tools

1. `Golang` (server) with `pty` package
2. `JS` (client) with `xterm.js` library
3. `WebSocket` for client-server connection
4. `Docker compose` for app deployment
