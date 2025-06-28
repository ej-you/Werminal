let term /* Terminal */ = null;
let ws /* WebSocket */ = null;
// получение HTML элемента для терминала
const terminalElement = function () {
    let elem = document.getElementById('terminal');
    if (elem == null) {
        elem = document.createElement("div")
        document.body.appendChild(elem)
    }
    return elem
}()

function closeTerm() {
    terminalElement.style.display = "none"
    term.dispose()
    term = null

    document.getElementById('new-session').style.display = "block"
    document.getElementById('close-session').style.display = "none"
}

function StopTerminal() {
    console.log("stop terminal")
    if (term === null) {
        return
    }
    terminalElement.style.display = "none"

    // закрываем WebSocket соединение и терминал
    ws.close(1000)
    closeTerm()
}

function StartTerminal() {
    console.log("start terminal")
    if (term !== null) {
        return
    }
    terminalElement.style.display = "block"

    // создание терминала
    term = NewTerminal(terminalElement)
    SetupTerminalFocus(terminalElement, term)

    // открытие WebSocket соединения
    ws = NewWebSocket(term.rows, term.cols)

    // настройка обработчиков для ws и терминала
    SetupOnMessage(ws, term)
    SetupOnError(ws, term)
    SetupOnClose(ws, term, closeTerm)
    SetupOnKey(ws, term)

    document.getElementById('new-session').style.display = "none"
    document.getElementById('close-session').style.display = "block"
}

window.addEventListener('DOMContentLoaded', () => {
    document.getElementById('new-session').addEventListener('click', StartTerminal);
    document.getElementById('close-session').addEventListener('click', StopTerminal);
});