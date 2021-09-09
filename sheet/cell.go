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
	data      interface{}
	dataType  string
	dataLen   int
	row       int
	col       int
	isHeader  bool
	isIndexer bool
}

func newCell(data interface{}, row, col int, isHeader, isIndexer bool) *Cell {
	c := &Cell{
		data:      data,
		dataType:  getDataType(data),
		row:       row,
		col:       col,
		isHeader:  isHeader,
		isIndexer: isIndexer,
	}
	c.dataLen = len(c.String())

	if c.dataType == stringType {
		if len(c.String()) > maxLen {
			c.data = c.String()[:maxLen-3] + "..."
			c.dataLen = maxLen
		}
	} else if c.dataType == numberType {
		if len(c.String()) > maxLen {
			c.data = c.String()[:maxLen]
			c.dataLen = maxLen
		}
		c.data = fmt.Sprintf("%s%s%s", colorYellow, c.String(), colorReset)
	} else if c.dataType == otherType {
		c.data = fmt.Sprintf("%s%s%s", colorPurple, nestedJSON, colorReset)
		c.dataLen = len(nestedJSON)
	}

	if c.isIndexer {
		c.data = fmt.Sprintf("%s%s%s", colorGreen, c.String(), colorReset)
	} else if c.isHeader {
		c.data = fmt.Sprintf("%s%s%s", colorBlue, c.String(), colorReset)
	}

	return c
}

func (c *Cell) String() string {
	return fmt.Sprintf("%v", c.data)
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
