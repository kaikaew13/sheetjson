package main

import (
	"log"
	"os"

	"github.com/kaikaew13/sheetjson/json"
	"github.com/kaikaew13/sheetjson/sheet"
)

func main() {
	f, err := os.Open("ex2.json")
	if err != nil {
		log.Fatalln("fail to open file")
	}
	defer f.Close()

	om, err := json.Setup(f)
	if err != nil {
		log.Fatalln(err)
	}

	sh := sheet.NewSheet("Testing", om)
	sh.Display()
}
