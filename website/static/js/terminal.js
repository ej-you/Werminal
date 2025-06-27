document.addEventListener('DOMContentLoaded', function () {
    const terminalElement = document.getElementById('terminal');

    // Состояние терминала
    const state = {
        inputBuffer: '',
        cursorPos: 0,
        history: [],
        historyIndex: -1,
        currentPrefix: null
    };

    // Инициализация терминала
    const term = new Terminal({
        cursorBlink: true,
        fontSize: 14,
        theme: {
            background: '#222222',
            foreground: '#FFFFFF'
        }
    });
    const fitAddon = new FitAddon.FitAddon();

    term.loadAddon(fitAddon);
    term.open(terminalElement);
    fitAddon.fit(); // Автоподгонка под размер окна

    // WebSocket соединение
    const addr = `ws://127.0.0.1:8080/api/v1/ws/terminal/?rows=${term.rows}&cols=${term.cols}`
    const socket = new WebSocket(addr);

    // Очистка текущей строки ввода
    function clearCurrentLine() {
        term.write('\x1b[2K\r');
        term.write(state.inputBuffer);
        moveCursor(state.cursorPos);
    }

    // Перемещение курсора
    function moveCursor(pos) {
        term.write(`\r`);
        if (pos > 0) {
            term.write(`\x1b[${pos}C`);
        }
        state.cursorPos = pos;
    }

    // Обработка ввода с клавиатуры
    term.onKey(e => {
        // const key = e.domEvent.key;

        // // Обработка Enter
        // if (key === 'Enter') {
        //     socket.send('\r\n');
        //     term.write('\r\n');
        // }
        // // Обработка Backspace
        // else if (key === 'Backspace') {
        //     socket.send('\b \b');
        //     term.write('\b \b');
        // }
        // // Остальные символы
        // else if (key.length === 1) {
        //     socket.send(key);
        //     term.write(key);
        // }
        // else if (key === 'Tab') {
        //     socket.send('\t');
        //     term.write('\t');
        // }

        // // Стрелки влево/вправо
        // if (key === 'ArrowLeft') {
        //     if (state.cursorPos > state.prompt.length) {
        //         state.cursorPos--;
        //         term.write('\x1b[D');
        //     }
        //     return;
        // }

        // if (key === 'ArrowRight') {
        //     if (state.cursorPos < state.prompt.length + state.inputBuffer.length) {
        //         state.cursorPos++;
        //         term.write('\x1b[C');
        //     }
        //     return;
        // }

        // // Стрелки вверх/вниз (история)
        // if (key === 'ArrowUp' || key === 'ArrowDown') {
        //     if (state.history.length === 0) return;

        //     if (key === 'ArrowUp' && state.historyIndex < state.history.length - 1) {
        //         state.historyIndex++;
        //     } else if (key === 'ArrowDown' && state.historyIndex > 0) {
        //         state.historyIndex--;
        //     } else if (key === 'ArrowDown' && state.historyIndex === 0) {
        //         state.historyIndex = -1;
        //         state.inputBuffer = '';
        //         clearCurrentLine();
        //         return;
        //     }

        //     if (state.historyIndex >= 0) {
        //         state.inputBuffer = state.history[state.historyIndex];
        //         state.cursorPos = state.prompt.length + state.inputBuffer.length;
        //         clearCurrentLine();
        //     }
        //     return;
        // }
        const key = e.domEvent.key;
        const printable = !e.domEvent.ctrlKey && !e.domEvent.altKey && !e.domEvent.metaKey;

        // Enter
        if (key === 'Enter') {
            const command = state.inputBuffer.trim();
            if (command) {
                state.history.unshift(command);
                state.historyIndex = -1;
            }
            state.inputBuffer = '';
            state.cursorPos = 0;

            socket.send(command + '\n');
            term.write('\r\n');
            state.waitingForPrompt = true;
            return;
        }

        // Backspace - только для введенного текста
        if (key === 'Backspace') {
            if (state.cursorPos > 0) {
                const pos = state.cursorPos - 1;
                state.inputBuffer = state.inputBuffer.slice(0, pos) + state.inputBuffer.slice(pos + 1);
                state.cursorPos = pos;

                // Удаляем символ правильно с учетом Unicode
                term.write('\b \b');
                // Переписываем остаток строки если удаление в середине
                if (pos < state.inputBuffer.length) {
                    const rest = state.inputBuffer.slice(pos);
                    term.write(rest + ' ');
                    moveCursor(pos);
                }
            }
            return;
        }

        // Стрелки влево/вправо
        if (key === 'ArrowLeft') {
            if (state.cursorPos > 0) {
                state.cursorPos--;
                term.write('\x1b[D');
            }
            return;
        }

        if (key === 'ArrowRight') {
            if (state.cursorPos < state.inputBuffer.length) {
                state.cursorPos++;
                term.write('\x1b[C');
            }
            return;
        }

        // Стрелки вверх/вниз (история)
        if (key === 'ArrowUp' || key === 'ArrowDown') {
            if (state.history.length === 0) return;

            if (key === 'ArrowUp' && state.historyIndex < state.history.length - 1) {
                state.historyIndex++;
            } else if (key === 'ArrowDown' && state.historyIndex > 0) {
                state.historyIndex--;
            } else if (key === 'ArrowDown' && state.historyIndex === 0) {
                state.historyIndex = -1;
                state.inputBuffer = '';
                clearCurrentLine();
                return;
            }

            if (state.historyIndex >= 0) {
                state.inputBuffer = state.history[state.historyIndex];
                state.cursorPos = state.inputBuffer.length;
                clearCurrentLine();
            }
            return;
        }

        // Tab (автодополнение)
        if (key === 'Tab') {
            // Реализация автодополнения
            return;
        }

        // Обычные символы
        if (printable && key.length === 1) {
            const pos = state.cursorPos;
            state.inputBuffer = state.inputBuffer.slice(0, pos) + key + state.inputBuffer.slice(pos);
            state.cursorPos++;

            // Вставка символа
            term.write(key);
            // Переписываем остаток строки если вставка в середину
            if (pos < state.inputBuffer.length - 1) {
                term.write(state.inputBuffer.slice(pos + 1));
            }
        }
    });

    // Данные с сервера -> вывод в терминал
    socket.onmessage = (event) => {
        term.write(event.data);
    };

    // Обработка ошибок
    socket.onerror = (error) => {
        term.writeln(`\x1b[31mWebSocket error: ${error.message}\x1b[0m`);
    };

    socket.onclose = () => {
        term.writeln('\x1b[31mGoodbye!\x1b[0m');
    };

    // Автоматический ресайз
    window.addEventListener('resize', () => fitAddon.fit());

    // Фокус на терминал при клике
    terminalElement.addEventListener('click', () => {
        term.focus();
    });

    // Первоначальный фокус
    term.focus();
});
