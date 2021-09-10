package sheet

import (
	"fmt"
)

const (
	asciiA       int = 65
	noOfAlphabet int = 26
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

	indexCell := newCell(colID, 0, colIDToNum(colID), false, true)
	col.cells = append(col.cells, indexCell)
	col.setMaxLen(indexCell.data.dLen)

	headerCell := newCell(header, 1, colIDToNum(colID), true, false)
	col.cells = append(col.cells, headerCell)
	col.setMaxLen(headerCell.data.dLen)

	i := 2
	for _, b := range body {
		c := newCell(b, i, colIDToNum(colID), false, false)
		col.cells = append(col.cells, c)
		col.setMaxLen(c.data.dLen)
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
		col.setMaxLen(c.data.dLen)
	}

	return col
}

func (col *Col) setMaxLen(dataLength int) {
	if col.maxDataLen < dataLength {
		col.maxDataLen = dataLength
	}
}

func colIDToNum(colID string) int {
	n, base := 0, 1
	for _, r := range colID {
		n += (int(r) - asciiA + 1) * base
		base *= noOfAlphabet
	}

	return n
}

// 1 = A, 2 = B, ..., 27 = AA
// like in the spreadsheet
func toColID(n int) string {
	s := []byte{}
	for n > 0 {
		r := (n - 1) % noOfAlphabet
		s = append(s, byte(asciiA+r))
		if (n % noOfAlphabet) == 0 {
			n /= 26
			n--
		} else {
			n /= noOfAlphabet
		}
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
