"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TermState = void 0;
exports.NewTerminal = NewTerminal;
exports.SetupTerminalFocus = SetupTerminalFocus;
const xterm_1 = require("@xterm/xterm");
const addon_fit_1 = require("@xterm/addon-fit");
// @param: elem - div элемент для терминала
// @return: Terminal
function NewTerminal(elem) {
    // инициализация терминала
    const term = new xterm_1.Terminal({
        cursorBlink: true,
        fontSize: 14,
        theme: {
            background: '#222222',
            foreground: '#FFFFFF'
        }
    });
    const fitAddon = new addon_fit_1.FitAddon();
    // автоподгонка под размер элемента
    term.loadAddon(fitAddon);
    term.open(elem);
    fitAddon.fit();
    // автоматическое изменение размера терминала
    window.addEventListener('resize', () => fitAddon.fit());
    return term;
}
// настройка фокуса на терминал
// @param: elem - div элемент для терминала
// @param: term - Terminal
function SetupTerminalFocus(elem, term) {
    // фокус на терминал при клике
    elem.addEventListener('click', () => {
        term.focus();
    });
    // фокус при загрузке страницы
    term.focus();
}
// состояние терминала
class TermState {
    inputBuffer;
    cursorPos;
    history;
    historyIndex;
    currentPrefix;
    constructor() {
        this.inputBuffer = "";
        this.cursorPos = 0;
        this.history = [];
        this.historyIndex = -1;
        this.currentPrefix = null;
    }
}
exports.TermState = TermState;
