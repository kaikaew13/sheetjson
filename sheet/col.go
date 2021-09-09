package sheet

import "fmt"

type Col struct {
	cells      []*Cell
	maxDataLen int
	colId      string
}

func newCol(row int, colId string) *Col {
	col := &Col{
		colId: colId,
	}

	indexCell := newCell(colId, 1, colIDToNum(colId), false, true)
	col.cells = append(col.cells, indexCell)

	for i := 1; i <= row; i++ {
		c := newCell("ABC", 1, colIDToNum(colId), true, false)
		col.cells = append(col.cells, c)
		col.setMaxLen(c.dataLen)
	}

	return col
}

func newIndexCol(row int) *Col {
	col := &Col{}

	indexCell := newCell("", 0, 0, false, true)
	col.cells = append(col.cells, indexCell)

	for i := 1; i <= row; i++ {
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
		n += (int(r) - A + 1) * base
		base *= 26
	}

	return n
}
