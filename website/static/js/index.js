// import { NewWebSocket } from './ws/ws.js';
// import { NewTerminal, SetupTerminalFocus, TermState } from './terminal/terminal.js';
// import { TermState } from './terminal/state.js';
// import { SetupOnMessage } from './terminal/ws/get/message.js';
// import { SetupOnError } from './terminal/ws/get/error.js';
// import { SetupOnClose } from './terminal/ws/get/close.js';
// import { SetupOnKey } from './terminal/ws/send/key.js';

// получение HTML элемента для терминала
const terminalElement = function () {
    let elem = document.getElementById('terminal');
    if (elem == null) {
        elem = document.createElement("div")
        document.body.appendChild(elem)
    }
    return elem
}()

// создание терминала
const term = NewTerminal(terminalElement)
SetupTerminalFocus(terminalElement, term)
// инициализация состояния терминала
let termState = new TermState()

// открытие WebSocket соединения
const ws = NewWebSocket(term.rows, term.cols)

// настройка обработчиков для ws и терминала
SetupOnMessage(ws, term)
SetupOnError(ws, term)
SetupOnClose(ws, term)
SetupOnKey(ws, term, termState)
