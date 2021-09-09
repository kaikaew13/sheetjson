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
	for _, col := range *cols {
		for _, c := range col.cells {
			fmt.Printf("%s", c.dataString())
		}
	}
}
