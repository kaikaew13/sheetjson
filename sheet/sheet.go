package sheet

import "fmt"

const asciiA int = 65

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

	i := 0
	for k, v := range m {
		col := newCol(sh.rows, string(rune(asciiA+i)), k, v)
		sh.cols = append(sh.cols, col)
		i++
	}

	// for i := 0; i < len(a[0]); i++ {
	// 	col := newCol(sh.rows, string(rune(A+i)))
	// 	sh.cols = append(sh.cols, col)
	// }

	return sh
}

func (sh *Sheet) Display() {
	for i := 0; i <= sh.rows; i++ {
		if i == 0 {
			printBorder(sh.cols)
		}

		printData(sh.cols, i)
		printBorder(sh.cols)
	}
}

func printBorder(cols []*Col) {
	for j := 0; j < len(cols); j++ {
		s := formatBorder(cols[j].maxDataLen)
		if j == 0 {
			fmt.Printf("+-%s-+", s)
		} else {
			fmt.Printf("-%s-+", s)
		}
	}

	fmt.Println()
}

func printData(cols []*Col, i int) {
	for j := 0; j < len(cols); j++ {
		col, cell := cols[j], cols[j].cells[i]
		s := formatPadding(cell.String(), cell.dataLen, col.maxDataLen)
		if j == 0 {
			fmt.Printf("| %s |", s)
		} else {
			fmt.Printf(" %s |", s)
		}
	}

	fmt.Println()
}

func formatBorder(maxDataLen int) string {
	s := ""
	for i := 0; i < maxDataLen; i++ {
		s += "-"
	}

	return s
}

func formatPadding(data string, dataLen, maxDataLen int) string {
	s := data
	for i := 0; i < maxDataLen-dataLen; i++ {
		s += " "
	}

	return s
}
