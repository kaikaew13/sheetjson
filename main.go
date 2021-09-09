package main

import "github.com/kaikaew13/sheetjson/sheet"

func main() {
	sh := sheet.NewSheet("Testing")
	sh.Display()
}

var design = `

+---+---+---+---+---+
| A | B | C | D | E |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+

`
