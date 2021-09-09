package sheet

type Col struct {
	cells         []*Cell
	maxDataLength int
	colLabel      string
}

func newCol(n int, colLabel string) *Col {
	col := &Col{
		colLabel: colLabel,
	}

	indexCell := newCell(1, colLabel, false, true)
	col.cells = append(col.cells, indexCell)

	for i := 2; i <= n; i++ {
		c := newCell(1, colLabel, false, false)
		col.cells = append(col.cells, c)
		col.maxLen(c.dataString())
	}

	return col
}

func (col *Col) maxLen(data string) {
	if col.maxDataLength < len(data) {
		col.maxDataLength = len(data)
	}
}
