// состояние терминала
class TermState {
    inputBuffer /* string */;
    cursorPos /* number */;
    history /* array */;
    historyIndex /* number */;
    currentPrefix /* string */;

    constructor() {
        this.inputBuffer = "";
        this.cursorPos = 0;
        this.history = [];
        this.historyIndex = -1;
        this.currentPrefix = null;
    }

    // передвигает курсор вправо на 1 позицию
    moveCursorRight() {
        console.log("AAAAAAA", this.cursorPos, this.inputBuffer.length)
        if (this.cursorPos < this.inputBuffer.length) {
            this.cursorPos++;
        }
    }

    // передвигает курсор влево на 1 позицию
    moveCursorLeft() {
        if (this.cursorPos > 0) {
            this.cursorPos--;
        }
    }
}
