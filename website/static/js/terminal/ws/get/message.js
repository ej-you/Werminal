// import { Terminal } from "@xterm/xterm";

// вывод данных сервера в терминал
function SetupOnMessage(socket /* WebSocket */, term /* Terminal */) {
    socket.onmessage = (event /* MessageEvent */) => {
        term.write(event.data);
    };
}
