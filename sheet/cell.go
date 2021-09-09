package sheet

import "fmt"

type Cell struct {
	data       interface{}
	dataLength int
	row        int
	col        int
	isHeader   bool
	isIndexer  bool
}

func newCell(data string, row int, col int, isHeader, isIndexer bool) *Cell {
	c := &Cell{
		data:       data,
		dataLength: len(data),
		row:        row,
		col:        col,
		isHeader:   isHeader,
		isIndexer:  isIndexer,
	}

	if c.isIndexer {
		c.data = fmt.Sprintf("\033[32m%s\033[0m", data)
	}

	return c
}

func (c *Cell) dataString() string {
	return c.data.(string)
}
