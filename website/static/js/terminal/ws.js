// import { Terminal } from "@xterm/xterm";

// обработка закрытия websocket соединения
function SetupOnClose(socket /* WebSocket */, term /* Terminal */, closeFunc /* function */) {
    socket.onclose = (event /* CloseEvent */) => {
        term.clear()
        term.writeln('\x1b[31mGoodbye!\x1b[0m');
        console.error(`Websocket connection closed: ${event.code} ${event.reason}`);

        closeFunc()
    };
}

// обработка ошибки от сервера
function SetupOnError(socket /* WebSocket */, term /* Terminal */) {
    socket.onerror = (errorEvent /* Event as ErrorEvent*/) => {
        term.writeln(`\x1b[31mconnection error: ${errorEvent.message}\x1b[0m`);
        console.error(`Websocket connection error: ${errorEvent.message}`);
    };
}

// вывод данных сервера в терминал
function SetupOnMessage(socket /* WebSocket */, term /* Terminal */) {
    socket.onmessage = (event /* MessageEvent */) => {
        term.write(event.data);
    };
}

// обработка ввода с клавиатуры
function SetupOnKey(socket /* WebSocket */, term /* Terminal */) {
    term.onData(data => {
        // console.log(`"${data}" raw: ${JSON.stringify(data)} | lenght: `, data.length)
        socket.send(data);
    });
}
