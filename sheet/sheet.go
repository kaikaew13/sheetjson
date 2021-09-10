package sheet

import "fmt"

type Sheet struct {
	cols []*Col
	rows int
	name string
}

func NewSheet(name string, m map[string][]interface{}, rows int) *Sheet {
	sh := &Sheet{
		name: name,
		rows: rows,
	}

	indexCol := newIndexCol(sh.rows)
	sh.cols = append(sh.cols, indexCol)

	i := 1
	for k, v := range m {
		col := newCol(sh.rows, toColID(i), k, v)
		sh.cols = append(sh.cols, col)
		i++
	}

	return sh
}

func (sh *Sheet) Display() {
	for i := 0; i <= sh.rows; i++ {
		if i == 0 {
			sh.printBorder()
		}

		sh.printData(i)
		sh.printBorder()
	}
}

func (sh *Sheet) printBorder() {
	for j := 0; j < len(sh.cols); j++ {
		s := fmtBorder(sh.cols[j].maxDataLen)
		if j == 0 {
			fmt.Printf("+-%s-+", s)
		} else {
			fmt.Printf("-%s-+", s)
		}
	}

	fmt.Println()
}

func (sh *Sheet) printData(i int) {
	for j := 0; j < len(sh.cols); j++ {
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
