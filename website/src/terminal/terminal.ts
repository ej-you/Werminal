import { Terminal } from "@xterm/xterm";
import { FitAddon } from "@xterm/addon-fit";

// @param: elem - div элемент для терминала
// @return: Terminal
export function NewTerminal(elem: HTMLElement): Terminal {
    // инициализация терминала
    const term = new Terminal({
        cursorBlink: true,
        fontSize: 14,
        theme: {
            background: '#222222',
            foreground: '#FFFFFF'
        }
    });
    const fitAddon = new FitAddon();

    // автоподгонка под размер элемента
    term.loadAddon(fitAddon);
    term.open(elem);
    fitAddon.fit();

    // автоматическое изменение размера терминала
    window.addEventListener('resize', () => fitAddon.fit());

    return term
}

// настройка фокуса на терминал
// @param: elem - div элемент для терминала
// @param: term - Terminal
export function SetupTerminalFocus(elem: HTMLElement, term: Terminal) {
    // фокус на терминал при клике
    elem.addEventListener('click', () => {
        term.focus();
    });
    // фокус при загрузке страницы
    term.focus();
}

// состояние терминала
export class TermState {
    inputBuffer: string;
    cursorPos: 0;
    history: [];
    historyIndex: -1;
    currentPrefix: null;

    constructor() {
        this.inputBuffer = "";
        this.cursorPos = 0;
        this.history = [];
        this.historyIndex = -1;
        this.currentPrefix = null;
    }
}
