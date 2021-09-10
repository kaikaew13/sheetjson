package sheet

import "fmt"

const (
	stringType string = "string"
	numberType string = "number"
	otherType  string = "other"

	nestedJSON string = "<nested>"

	maxLen int = 25

	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
	colorBlue   string = "\033[34m"
	colorPurple string = "\033[35m"
	colorCyan   string = "\033[36m"
	colorWhite  string = "\033[37m"
)

type Cell struct {
	data      *Data
	row       int
	col       int
	isHeader  bool
	isIndexer bool
}

type Data struct {
	d     interface{}
	dType string
	dLen  int
}

func (d *Data) String() string {
	return fmt.Sprintf("%v", d.d)
}

func newCell(d interface{}, row, col int, isHeader, isIndexer bool) *Cell {
	dt := Data{
		d: d,
	}
	dt.dLen = len(dt.String())
	dt.dType = getDataType(dt.d)

	c := &Cell{
		data:      &dt,
		row:       row,
		col:       col,
		isHeader:  isHeader,
		isIndexer: isIndexer,
	}

	c.fmtData()

	if c.isIndexer {
		c.data.d = fmt.Sprintf("%s%s%s", colorGreen, c.data.String(), colorReset)
	} else if c.isHeader {
		c.data.d = fmt.Sprintf("%s%s%s", colorCyan, c.data.String(), colorReset)
	}

	return c
}

func getDataType(data interface{}) string {
	switch data.(type) {
	case string:
		return stringType
	case int, int8, int16, int32, int64, float32, float64:
		return numberType
	default:
		return otherType
	}
}

func (c *Cell) fmtData() {
	switch c.data.dType {
	case stringType:
		if c.data.dLen > maxLen {
			c.data.d = c.data.String()[:maxLen-3] + "..."
			c.data.dLen = maxLen
		}
	case numberType:
		if c.data.dLen > maxLen {
			c.data.d = c.data.String()[:maxLen]
			c.data.dLen = maxLen
		}
		c.data.d = fmt.Sprintf("%s%s%s", colorYellow, c.data.String(), colorReset)
	case otherType:
		c.data.d = fmt.Sprintf("%s%s%s", colorPurple, nestedJSON, colorReset)
		c.data.dLen = len(nestedJSON)
	}
}
