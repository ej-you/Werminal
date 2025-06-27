"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SetupOnError = SetupOnError;
// обработка ошибки от сервера
function SetupOnError(socket, term) {
    socket.onerror = (event) => {
        const errorEvent = event;
        term.writeln(`\x1b[31mconnection error: ${errorEvent.message}\x1b[0m`);
        console.error(`Websocket connection error: ${errorEvent.message}`);
    };
}
