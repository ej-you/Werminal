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

// открытие WebSocket соединения
const ws = NewWebSocket(term.rows, term.cols)

// настройка обработчиков для ws и терминала
SetupOnMessage(ws, term)
SetupOnError(ws, term)
SetupOnClose(ws, term)
SetupOnKey(ws, term)
