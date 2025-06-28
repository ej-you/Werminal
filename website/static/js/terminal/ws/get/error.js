// import { Terminal } from "@xterm/xterm";

// обработка ошибки от сервера
function SetupOnError(socket /* WebSocket */, term /* Terminal */) {
    socket.onerror = (errorEvent /* Event as ErrorEvent*/) => {
        term.writeln(`\x1b[31mconnection error: ${errorEvent.message}\x1b[0m`);
        console.error(`Websocket connection error: ${errorEvent.message}`);
    };
}