package sheet

import "fmt"

const (
	A int = 65

	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
	colorBlue   string = "\033[34m"
	colorPurple string = "\033[35m"
	colorCyan   string = "\033[36m"
	colorWhite  string = "\033[37m"
)

type Sheet struct {
	cols []*Col
	rows int
	name string
}

func NewSheet(name string) *Sheet {
	sh := &Sheet{
		name: name,
		rows: 5,
	}

	indexCol := newIndexCol(sh.rows)
	sh.cols = append(sh.cols, indexCol)

	for i := 0; i < 26; i++ {
		col := newCol(sh.rows, string(rune(A+i)))
		sh.cols = append(sh.cols, col)
	}

	return sh
}

func (sh *Sheet) Display() {
	for i := 0; i < sh.rows; i++ {
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
		s := formatPadding(cell.dataToString(), cell.dataLen, col.maxDataLen)
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
