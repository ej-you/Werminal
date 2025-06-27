"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NewWebSocket = NewWebSocket;
// открытие WebSocket соединения
function NewWebSocket(rows, cols) {
    const addr = `ws://127.0.0.1:8080/api/v1/ws/terminal/?rows=${rows}&cols=${cols}`;
    return new WebSocket(addr);
}
