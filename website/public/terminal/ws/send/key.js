"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SetupOnKey = SetupOnKey;
// Обработка ввода с клавиатуры
function SetupOnKey(ws, term, termState) {
    term.onKey(e => {
        console.log("key:", e.key);
        console.log("domEvent:", e.domEvent);
    });
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
