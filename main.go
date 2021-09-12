package main

import (
	"log"
	"os"

	"github.com/kaikaew13/sheetjson/json"
	"github.com/kaikaew13/sheetjson/sheet"
	"golang.org/x/term"
)

var (
	terminalWidth  int
	terminalHeight int
)

func main() {
	terminalWidth, terminalHeight, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err)
	}

	if len(os.Args) != 2 {
		log.Fatalln("need to provide json file as second arg")
	}

	fileName := os.Args[1]

	if fileName[len(fileName)-4:] != "json" {
		log.Fatalln("require to be json file")
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("fail to open file")
	}
	defer f.Close()

	om, err := json.Setup(f)
	if err != nil {
		log.Fatalln(err)
	}

	sh := sheet.NewSheet("Testing", om)
	sh.Display(terminalWidth, terminalHeight)
}
