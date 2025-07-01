const host = "wss://fredcv.ru:8092";

// открытие WebSocket соединения
function NewWebSocket(rows /* number */, cols /* number */) /* WebSocket */ {
    const addr = `${host}/api/v1/ws/terminal/?rows=${rows}&cols=${cols}`
    return new WebSocket(addr);
}
