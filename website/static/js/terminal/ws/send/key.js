// import { Terminal } from "@xterm/xterm";
// import { TermState } from "../../state.js";

// возвращает true если введён число-буквенно-символьный (печатаемый) символ
// @return: boolean
function isPrintable(keyEvent /* KeyboardEvent */) /* boolean */ {
    if (keyEvent.ctrlKey || keyEvent.altKey || keyEvent.metaKey) {
        return false
    }
    const key = keyEvent.key;
    if (["Backspace", "Tab", "Insert", "Delete"].includes(key)) {
        return false
    }
    if (["PageUp", "PageDown", "Home", "End"].includes(key)) {
        return false
    }
    if (["Enter", "Escape"].includes(key)) {
        return false
    }
    if (["ArrowLeft", "ArrowUp", "ArrowRight", "ArrowDown"].includes(key)) {
        return false
    }
    const fRegex = /^F[123456789][012]?$/; // Fn keys check
    if (key.match(fRegex) != null) {
        return false
    }
    return true
}

function printable(ws /* WebSocket */, term /* Terminal */, termState /* TermState */, key /* string */) {
    console.log("print:", key)
    ws.send(key);

    console.log("termState 1:", termState)

    // высчитываем позицию курсора и вставляем элемент в нужное место буфера
    const cursor = termState.cursorPos;
    termState.inputBuffer = termState.inputBuffer.slice(0, cursor) + key + termState.inputBuffer.slice(cursor);
    termState.moveCursorRight()

    console.log("termState 2:", termState)

    // вставка символа
    term.write(key);
    const newCursor = termState.cursorPos;
    // переписываем остаток строки если вставка в середину
    if (newCursor < termState.inputBuffer.length - 1) {
        term.write(termState.inputBuffer.slice(newCursor));
    }
}

function arrowLeft(ws /* WebSocket */, term /* Terminal */, termState /* TermState */) {
    term.write("\x1b[2K");
    return

    const cursorBefore = termState.cursorPos
    termState.moveCursorLeft()

    // если курсор сдвинулся
    if (cursorBefore != termState.cursorPos) {
        term.write("\x1b[D");
        ws.send("\x1b[D");
    }
}

function arrowRight(ws /* WebSocket */, term /* Terminal */, termState /* TermState */) {
    const cursorBefore = termState.cursorPos
    termState.moveCursorRight()

    // если курсор сдвинулся
    if (cursorBefore != termState.cursorPos) {
        term.write("\x1b[C");
        ws.send("\x1b[C");
    }
}

// обработка ввода с клавиатуры
function SetupOnKey(ws /* WebSocket */, term /* Terminal */, termState /* TermState */) {
    term.onKey(e => {
        // if (keyEvent.ctrlKey || keyEvent.altKey || keyEvent.metaKey) {
        //     return
        // }
        // console.log("keyEvent: ", keyEvent)

        const keyEvent = e.domEvent;
        const keyName = keyEvent.key;

        // const termOutput = keyEvent.key; // строка для записи в элемент терминала
        // const wsOutput = keyEvent.key; // строка для отправки на сервер

        switch (true) {
            // любая печатаемая цифра/буква/символ
            case isPrintable(keyEvent):
                printable(ws, term, termState, keyName)
                break;
            // case keyName == "ArrowLeft":
            //     arrowLeft(ws, term, termState)
            //     break;
            // case keyName == "ArrowRight":
            //     arrowRight(ws, term, termState)
            //     break;
            default:
                console.log("NOT PRINTABLE:", e.domEvent)
        }
    });
    // term.onData(raw => {
    //     let data = raw;
    //     if (data == "\r") {
    //         data = "\r\n"
    //     }

    //     console.log(`"${data}" | lenght: `, data.length)
    //     term.write(data);
    //     ws.send(data);
    // });
}

// term.onKey(e => {
//     // const key = e.domEvent.key;

//     // // Обработка Enter
//     // if (key === 'Enter') {
//     //     socket.send('\r\n');
//     //     term.write('\r\n');
//     // }
//     // // Обработка Backspace
//     // else if (key === 'Backspace') {
//     //     socket.send('\b \b');
//     //     term.write('\b \b');
//     // }
//     // // Остальные символы
//     // else if (key.length === 1) {
//     //     socket.send(key);
//     //     term.write(key);
//     // }
//     // else if (key === 'Tab') {
//     //     socket.send('\t');
//     //     term.write('\t');
//     // }

//     // // Стрелки влево/вправо
//     // if (key === 'ArrowLeft') {
//     //     if (state.cursorPos > state.prompt.length) {
//     //         state.cursorPos--;
//     //         term.write('\x1b[D');
//     //     }
//     //     return;
//     // }

//     // if (key === 'ArrowRight') {
//     //     if (state.cursorPos < state.prompt.length + state.inputBuffer.length) {
//     //         state.cursorPos++;
//     //         term.write('\x1b[C');
//     //     }
//     //     return;
//     // }

//     // // Стрелки вверх/вниз (история)
//     // if (key === 'ArrowUp' || key === 'ArrowDown') {
//     //     if (state.history.length === 0) return;

//     //     if (key === 'ArrowUp' && state.historyIndex < state.history.length - 1) {
//     //         state.historyIndex++;
//     //     } else if (key === 'ArrowDown' && state.historyIndex > 0) {
//     //         state.historyIndex--;
//     //     } else if (key === 'ArrowDown' && state.historyIndex === 0) {
//     //         state.historyIndex = -1;
//     //         state.inputBuffer = '';
//     //         clearCurrentLine();
//     //         return;
//     //     }

//     //     if (state.historyIndex >= 0) {
//     //         state.inputBuffer = state.history[state.historyIndex];
//     //         state.cursorPos = state.prompt.length + state.inputBuffer.length;
//     //         clearCurrentLine();
//     //     }
//     //     return;
//     // }

//     const key = e.domEvent.key;
//     const printable = !e.domEvent.ctrlKey && !e.domEvent.altKey && !e.domEvent.metaKey;

//     // Enter
//     if (key === 'Enter') {
//         const command = state.inputBuffer.trim();
//         if (command) {
//             state.history.unshift(command);
//             state.historyIndex = -1;
//         }
//         state.inputBuffer = '';
//         state.cursorPos = 0;

//         socket.send(command + '\n');
//         term.write('\r\n');
//         state.waitingForPrompt = true;
//         return;
//     }

//     // Backspace - только для введенного текста
//     if (key === 'Backspace') {
//         if (state.cursorPos > 0) {
//             const pos = state.cursorPos - 1;
//             state.inputBuffer = state.inputBuffer.slice(0, pos) + state.inputBuffer.slice(pos + 1);
//             state.cursorPos = pos;

//             // Удаляем символ правильно с учетом Unicode
//             term.write('\b \b');
//             // Переписываем остаток строки если удаление в середине
//             if (pos < state.inputBuffer.length) {
//                 const rest = state.inputBuffer.slice(pos);
//                 term.write(rest + ' ');
//                 moveCursor(pos);
//             }
//         }
//         return;
//     }

//     // Стрелки влево/вправо
//     if (key === 'ArrowLeft') {
//         if (state.cursorPos > 0) {
//             state.cursorPos--;
//             term.write('\x1b[D');
//         }
//         return;
//     }

//     if (key === 'ArrowRight') {
//         if (state.cursorPos < state.inputBuffer.length) {
//             state.cursorPos++;
//             term.write('\x1b[C');
//         }
//         return;
//     }

//     // Стрелки вверх/вниз (история)
//     if (key === 'ArrowUp' || key === 'ArrowDown') {
//         if (state.history.length === 0) return;

//         if (key === 'ArrowUp' && state.historyIndex < state.history.length - 1) {
//             state.historyIndex++;
//         } else if (key === 'ArrowDown' && state.historyIndex > 0) {
//             state.historyIndex--;
//         } else if (key === 'ArrowDown' && state.historyIndex === 0) {
//             state.historyIndex = -1;
//             state.inputBuffer = '';
//             clearCurrentLine();
//             return;
//         }

//         if (state.historyIndex >= 0) {
//             state.inputBuffer = state.history[state.historyIndex];
//             state.cursorPos = state.inputBuffer.length;
//             clearCurrentLine();
//         }
//         return;
//     }

//     // Tab (автодополнение)
//     if (key === 'Tab') {
//         // Реализация автодополнения
//         return;
//     }

//     // Обычные символы
//     if (printable && key.length === 1) {
//         const pos = state.cursorPos;
//         state.inputBuffer = state.inputBuffer.slice(0, pos) + key + state.inputBuffer.slice(pos);
//         state.cursorPos++;

//         // Вставка символа
//         term.write(key);
//         // Переписываем остаток строки если вставка в середину
//         if (pos < state.inputBuffer.length - 1) {
//             term.write(state.inputBuffer.slice(pos + 1));
//         }
//     }
// });
