package sheet

import (
	"fmt"
)

const (
	asciiA        int = 65
	noOfAlphabets int = 26
)

type Col struct {
	cells      []*Cell
	maxDataLen int
	colID      string
}

func newCol(rows int, colID, header string, body []interface{}) *Col {
	col := &Col{
		colID: colID,
	}

	indexCell := newCell(colID, 0, colIDToInt(colID), false, true)
	col.cells = append(col.cells, indexCell)
	col.setColLen(indexCell.data.dLen)

	headerCell := newCell(header, 1, colIDToInt(colID), true, false)
	col.cells = append(col.cells, headerCell)
	col.setColLen(headerCell.data.dLen)

	for i, b := range body {
		c := newCell(b, i+2, colIDToInt(colID), false, false)
		col.cells = append(col.cells, c)
		col.setColLen(c.data.dLen)
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
		col.setColLen(c.data.dLen)
	}

	return col
}

func (col *Col) setColLen(dataLength int) {
	if col.maxDataLen < dataLength {
		col.maxDataLen = dataLength
	}
}

// A == 1, 2 == B, ..., 27 == AA
// like in the spreadsheet
func colIDToInt(colID string) int {
	colID = reverseString(colID)

	n, base := 0, 1
	for _, r := range colID {
		n += (int(r) - asciiA + 1) * base
		base *= noOfAlphabets
	}

	return n
}

// 1 == A, 2 == B, ..., 27 == AA
// like in the spreadsheet
func intToColID(n int) string {
	s := ""
	for n > 0 {
		r := (n - 1) % noOfAlphabets
		s += string(byte(asciiA + r))
		if (n % noOfAlphabets) == 0 {
			n /= 26
			n--
		} else {
			n /= noOfAlphabets
		}
	}

	return reverseString(s)
}

func reverseString(s string) string {
	b := []byte(s)

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
