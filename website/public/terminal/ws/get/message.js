"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SetupOnMessage = SetupOnMessage;
// вывод данных сервера в терминал
function SetupOnMessage(socket, term) {
    socket.onmessage = (event) => {
        term.write(event.data);
    };
}
