package sheet

import "fmt"

type Sheet struct {
	cols []*Col
	name string
}

func NewSheet() *Sheet {
	sh := &Sheet{
		name: "Default",
	}

	firstCol := newCol(5, byte('-'), true)
	sh.cols = append(sh.cols, firstCol)
	for i := 0; i < 26; i++ {
		col := newCol(5, byte(65+i), false)
		sh.cols = append(sh.cols, col)
	}

	return sh
}

func (sh *Sheet) Display() {
	cols := sh.cols

	for j := 0; j < 5; j++ {
		if j == 0 {
			for k := 0; k < len(cols); k++ {
				s := ""
				for i := 0; i < (cols)[k].maxDataLength; i++ {
					s += "-"
				}
				if k == 0 {
					fmt.Printf("+-%s-+", s)
				} else {
					fmt.Printf("-%s-+", s)
				}
			}
			fmt.Println()
			for k := 0; k < len(cols); k++ {
				s := (cols)[k].cells[j].dataString()
				tmp := (cols)[k].cells[j].dataLength
				for i := 0; i < (cols)[k].maxDataLength-tmp; i++ {
					s += " "
				}
				if k == 0 {
					fmt.Printf("| %s |", s)
				} else {
					fmt.Printf(" %s |", s)
				}
			}
			fmt.Println()
			for k := 0; k < len(cols); k++ {
				s := ""
				for i := 0; i < (cols)[k].maxDataLength; i++ {
					s += "-"
				}
				if k == 0 {
					fmt.Printf("+-%s-+", s)
				} else {
					fmt.Printf("-%s-+", s)
				}
			}
		} else {
			fmt.Println()
			for k := 0; k < len(cols); k++ {
				if k == 0 {
					fmt.Printf("| %s |", (cols)[k].cells[j].dataString())
				} else {
					fmt.Printf(" %s |", (cols)[k].cells[j].dataString())
				}
			}
			fmt.Println()
			for k := 0; k < len(cols); k++ {
				s := ""
				for i := 0; i < (cols)[k].maxDataLength; i++ {
					s += "-"
				}
				if k == 0 {
					fmt.Printf("+-%s-+", s)
				} else {
					fmt.Printf("-%s-+", s)
				}
			}
		}
	}
}
