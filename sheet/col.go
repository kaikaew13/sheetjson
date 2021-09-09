package sheet

import "fmt"

type Col struct {
	cells         []*Cell
	maxDataLength int
	id            string
}

func newCol(row int, colLabel byte, indexCol bool) *Col {
	col := &Col{
		id: string(colLabel),
	}

	indexCell := newCell(string(colLabel), 1, int(colLabel)-65, false, true)
	col.cells = append(col.cells, indexCell)

	for i := 1; i <= row; i++ {
		var c *Cell
		if indexCol {
			c = newCell(fmt.Sprintf("%d", i), 1, int(colLabel)-65, false, true)
		} else {
			c = newCell("ABC", 1, int(colLabel)-65, false, false)
		}
		col.cells = append(col.cells, c)
		col.maxLen(c.dataLength)
	}

	if col.maxDataLength == 10 {
		for _, v := range col.cells {
			fmt.Println(v.dataLength)
		}
	}
	return col
}

func (col *Col) maxLen(dataLength int) {
	if col.maxDataLength < dataLength {
		col.maxDataLength = dataLength
	}
}
