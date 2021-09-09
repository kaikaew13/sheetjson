package sheet

import "fmt"

func NewSheet() {
	cols := []Col{}
	for i := 0; i < 26; i++ {
		col := newCol(5, string(byte(65+i)))
		cols = append(cols, *col)
	}

	draw(&cols)
}

func draw(cols *[]Col) {
	for j := 0; j < 5; j++ {
		if j == 0 {
			for k := 0; k < len(*cols); k++ {
				s := ""
				for i := 0; i < (*cols)[k].maxDataLength; i++ {
					s += "-"
				}
				if k == 0 {
					fmt.Printf("+-%s-+", s)
				} else {
					fmt.Printf("-%s-+", s)
				}
			}
			fmt.Println()
			for k := 0; k < len(*cols); k++ {
				s := (*cols)[k].colLabel
				tmp := len(s)
				for i := 0; i < (*cols)[k].maxDataLength-tmp; i++ {
					s += " "
				}
				if k == 0 {
					fmt.Printf("| %s |", s)
				} else {
					fmt.Printf(" %s |", s)
				}
			}
			fmt.Println()
			for k := 0; k < len(*cols); k++ {
				s := ""
				for i := 0; i < (*cols)[k].maxDataLength; i++ {
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
			for k := 0; k < len(*cols); k++ {
				if k == 0 {
					fmt.Printf("| %s |", (*cols)[k].cells[j].dataString())
				} else {
					fmt.Printf(" %s |", (*cols)[k].cells[j].dataString())
				}
			}
			fmt.Println()
			for k := 0; k < len(*cols); k++ {
				s := ""
				for i := 0; i < (*cols)[k].maxDataLength; i++ {
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
