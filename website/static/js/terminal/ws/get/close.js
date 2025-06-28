// import { Terminal } from "@xterm/xterm";

// обработка закрытия websocket соединения
function SetupOnClose(socket /* WebSocket */, term /* Terminal */) {
    socket.onclose = (event /* CloseEvent */) => {
        term.clear()
        term.writeln('\x1b[31mGoodbye!\x1b[0m');
        console.error(`Websocket connection closed: ${event.code} ${event.reason}`);
    };
}