// Очистка текущей строки ввода
export function clearCurrentLine() {
    term.write('\x1b[2K\r');
    term.write(state.inputBuffer);
    moveCursor(state.cursorPos);
}

// Перемещение курсора
export function moveCursor(pos) {
    term.write(`\r`);
    if (pos > 0) {
        term.write(`\x1b[${pos}C`);
    }
    state.cursorPos = pos;
}
