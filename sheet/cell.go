package sheet

import "fmt"

type Cell struct {
	data      interface{}
	dataLen   int
	row       int
	col       int
	isHeader  bool
	isIndexer bool
}

func newCell(data string, row, col int, isHeader, isIndexer bool) *Cell {
	c := &Cell{
		data:      data,
		dataLen:   len(data),
		row:       row,
		col:       col,
		isHeader:  isHeader,
		isIndexer: isIndexer,
	}

	if c.isIndexer {
		c.data = fmt.Sprintf("%s%s%s", colorGreen, data, colorReset)
	}

	if c.isHeader {
		c.data = fmt.Sprintf("%s%s%s", colorBlue, data, colorReset)
	}

	return c
}

func (c *Cell) dataToString() string {
	return c.data.(string)
}
