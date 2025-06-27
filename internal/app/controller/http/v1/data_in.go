package v1

import (
	"errors"
	"strconv"

	gowebsocket "github.com/gofiber/websocket/v2"
)

const (
	_defaultTermRows = 35  // default pseudo-terminal rows
	_defaultTermCols = 180 // default pseudo-terminal columns
)

// terminalIn is an input query params for terminal websocket connection.
type terminalIn struct {
	Rows uint16 `query:"rows"`
	Cols uint16 `query:"cols"`
}

// parseTerminalIn parses query params to terminalIn struct.
func parseTerminalIn(conn *gowebsocket.Conn) (*terminalIn, error) {
	var (
		err        error
		rows, cols uint64
	)

	// parse rows
	rowsQuery := conn.Query("rows", "")
	if rowsQuery == "" {
		rows = _defaultTermRows
	} else {
		// convert rows
		rows, err = strconv.ParseUint(rowsQuery, 10, 16)
		if err != nil {
			return nil, errors.New("invalid rows value")
		}
	}

	colsQuery := conn.Query("cols", "")
	if rowsQuery == "" {
		cols = _defaultTermCols
	} else {
		// convert cols
		cols, err = strconv.ParseUint(colsQuery, 10, 16)
		if err != nil {
			return nil, errors.New("invalid rows value")
		}
	}

	return &terminalIn{
		Rows: uint16(rows),
		Cols: uint16(cols),
	}, nil
}
