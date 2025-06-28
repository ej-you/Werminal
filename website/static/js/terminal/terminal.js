// import { Terminal } from "@xterm/xterm";
// import { FitAddon } from "@xterm/addon-fit";

// @param: elem - div элемент для терминала
// @return: Terminal
function NewTerminal(elem /* HTMLElement */) /* Terminal */ {
    // инициализация терминала
    const term = new Terminal({
        cursorBlink: true,
        fontSize: 16,
        theme: {
            background: "#222222",
            foreground: "#FFFFFF",
            cursor: "#FFFFFF",
            cursorAccent: "#222222",
        }
    });
    // почему-то класс FitAddon загружается в window.FitAddon
    const fitAddon = new window.FitAddon.FitAddon();

    // автоподгонка под размер элемента
    term.loadAddon(fitAddon);
    term.open(elem);
    fitAddon.fit();

    // автоматическое изменение размера терминала
    window.addEventListener("resize", () => fitAddon.fit());

    return term
}

// настройка фокуса на терминал
// @param: elem - div элемент для терминала
// @param: term - Terminal
function SetupTerminalFocus(elem /* HTMLElement */, term /* Terminal */) {
    // фокус на терминал при клике
    elem.addEventListener("click", () => {
        term.focus();
    });
    // фокус при загрузке страницы
    term.focus();
}
