"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SetupOnClose = SetupOnClose;
// обработка закрытия websocket соединения
function SetupOnClose(socket, term) {
    socket.onclose = (event) => {
        term.clear();
        term.writeln('\x1b[31mGoodbye!\x1b[0m');
        console.error(`Websocket connection closed: ${event.code} ${event.reason}`);
    };
}
