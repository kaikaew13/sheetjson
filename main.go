package main

import "fmt"

var cell = `+---+
|   |
+---+`

func main() {

	for j := 0; j < 5; j++ {
		if j == 0 {
			for k := 0; k < 5; k++ {
				if k == 0 {
					fmt.Print("+---+")
				} else {
					fmt.Print("---+")
				}
			}
			fmt.Println()
			for k := 0; k < 5; k++ {
				if k == 0 {
					fmt.Print("|   |")
				} else {
					fmt.Print("   |")
				}
			}
			fmt.Println()
			for k := 0; k < 5; k++ {
				if k == 0 {
					fmt.Print("+---+")
				} else {
					fmt.Print("---+")
				}
			}
		} else {
			fmt.Println()
			for k := 0; k < 5; k++ {
				if k == 0 {
					fmt.Print("|   |")
				} else {
					fmt.Print("   |")
				}
			}
			fmt.Println()
			for k := 0; k < 5; k++ {
				if k == 0 {
					fmt.Print("+---+")
				} else {
					fmt.Print("---+")
				}
			}
		}
	}

}

var design = `

+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+
|   |   |   |   |   |
+---+---+---+---+---+

`
