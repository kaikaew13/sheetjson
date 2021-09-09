package sheet

type Cell struct {
	data      interface{}
	row       int
	col       string
	isHeader  bool
	isIndexer bool
}

func newCell(row int, col string, isHeader, isIndexer bool) *Cell {
	c := &Cell{
		data:      "ABC",
		row:       row,
		col:       col,
		isHeader:  isHeader,
		isIndexer: isIndexer,
	}

	return c
}

func (c *Cell) dataString() string {
	return c.data.(string)
}
