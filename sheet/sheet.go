package sheet

import (
	"fmt"

	"github.com/kaikaew13/sheetjson/json"
)

type Sheet struct {
	cols []*Col
	rows int
	name string
}

func NewSheet(name string, om *json.OrderedMap) *Sheet {
	sh := &Sheet{
		name: name,
		rows: om.ValLen,
	}
	sh.cols = make([]*Col, om.KeyLen+1)

	indexCol := newIndexCol(sh.rows)
	sh.cols[0] = indexCol

	for i := 1; ; i++ {
		k, ok := om.It.Next()
		if !ok {
			break
		}

		col := newCol(sh.rows, intToColID(i), k, om.M[k])
		sh.cols[i] = col
	}

	return sh
}

func (sh *Sheet) Display(maxX, maxY int) {
	fmt.Printf("sheet name: %s\n\n", sh.name)
	for i, lines := 0, 4; i <= sh.rows; i++ {
		if lines >= maxY {
			fmt.Printf("%d more rows...\n", sh.rows-i+1)
			break
		}

		if i == 0 {
			sh.printBorder(maxX)
			lines++
		}

		sh.printData(i, maxX)
		sh.printBorder(maxX)
		lines += 2
	}
}

func (sh *Sheet) printBorder(maxX int) {
	for j, chars := 0, 5; j < len(sh.cols); j++ {
		chars += 3 + sh.cols[j].maxDataLen
		if chars >= maxX {
			break
		}

		s := fmtBorder(sh.cols[j].maxDataLen)
		if j == 0 {
			fmt.Printf("+-%s-+", s)
		} else {
			fmt.Printf("-%s-+", s)
		}
	}

	fmt.Println()
}

func (sh *Sheet) printData(i, maxX int) {
	for j, chars := 0, 5; j < len(sh.cols); j++ {
		chars += 3 + sh.cols[j].maxDataLen
		if chars >= maxX {
			break
		}

		col, cell := sh.cols[j], sh.cols[j].cells[i]
		s := fmtPadding(cell.data.String(), cell.data.dLen, col.maxDataLen)
		if j == 0 {
			fmt.Printf("| %s |", s)
		} else {
			fmt.Printf(" %s |", s)
		}
	}

	fmt.Println()
}

func fmtBorder(maxDataLen int) string {
	s := ""
	for i := 0; i < maxDataLen; i++ {
		s += "-"
	}

	return s
}

func fmtPadding(data string, dataLen, maxDataLen int) string {
	s := data
	for i := 0; i < maxDataLen-dataLen; i++ {
		s += " "
	}

	return s
}
