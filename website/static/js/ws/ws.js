const hostAddr = "127.0.0.1:8803";

// открытие WebSocket соединения
function NewWebSocket(rows /* number */, cols /* number */) /* WebSocket */ {
    const addr = `ws://${hostAddr}/api/v1/ws/terminal/?rows=${rows}&cols=${cols}`
    return new WebSocket(addr);
}
