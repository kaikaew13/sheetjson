package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/kaikaew13/sheetjson/sheet"
)

func main() {
	f, err := os.Open("ex2.json")
	if err != nil {
		log.Fatalln("fail to open")
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("fail to readall")
	}

	a := []map[string]interface{}{}

	if err = json.Unmarshal(bytes, &a); err != nil {
		log.Fatalln(err)
	}

	m := make(map[string][]interface{})

	for i := range a {
		for k, v := range a[i] {
			m[k] = append(m[k], v)
		}
	}

	sh := sheet.NewSheet("Testing", m, len(a)+1)
	sh.Display()
}
