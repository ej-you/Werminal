import { Terminal } from "@xterm/xterm";

// обработка ошибки от сервера
export function SetupOnError(socket: WebSocket, term: Terminal) {
    socket.onerror = (event: Event) => {
        const errorEvent = event as ErrorEvent;
        term.writeln(`\x1b[31mconnection error: ${errorEvent.message}\x1b[0m`);
        console.error(`Websocket connection error: ${errorEvent.message}`);
    };
}