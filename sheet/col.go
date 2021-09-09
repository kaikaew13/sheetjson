package sheet

import (
	"fmt"
)

type Col struct {
	cells      []*Cell
	maxDataLen int
	colId      string
}

func newCol(rows int, colId, header string, body []interface{}) *Col {
	col := &Col{
		colId: colId,
	}

	indexCell := newCell(colId, 0, colIDToNum(colId), false, true)
	col.cells = append(col.cells, indexCell)
	col.setMaxLen(indexCell.dataLen)

	headerCell := newCell(header, 1, colIDToNum(colId), true, false)
	col.cells = append(col.cells, headerCell)
	col.setMaxLen(headerCell.dataLen)

	i := 2
	for _, b := range body {
		c := newCell(b, i, colIDToNum(colId), false, false)
		col.cells = append(col.cells, c)
		col.setMaxLen(c.dataLen)
		i++
	}

	return col
}

func newIndexCol(rows int) *Col {
	col := &Col{}

	indexCell := newCell("", 0, 0, false, true)
	col.cells = append(col.cells, indexCell)

	for i := 1; i <= rows; i++ {
		c := newCell(fmt.Sprintf("%d", i), i, 0, false, true)
		col.cells = append(col.cells, c)
		col.setMaxLen(c.dataLen)
	}

	return col
}

func (col *Col) setMaxLen(dataLength int) {
	if col.maxDataLen < dataLength {
		col.maxDataLen = dataLength
	}
}

func colIDToNum(colId string) int {
	n, base := 0, 1
	for _, r := range colId {
		n += (int(r) - asciiA + 1) * base
		base *= 26
	}

	return n
}
