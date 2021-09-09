package main

import "github.com/kaikaew13/sheetjson/sheet"

var cell = `+---+
|   |
+---+`

func main() {
	sh := sheet.NewSheet()
	sh.Display()
}

var design = `

+---+---+---+---+---+
| A | B | C | D | E |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+

`
