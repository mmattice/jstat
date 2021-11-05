package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Filedata struct {
	Filename string
	Size int64
	Mtime time.Time
}

func main() {
	files, err := os.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data := []Filedata{}

	for _, fdirentry := range files {
		statr, _ := os.Stat(fdirentry.Name())
		fdata := Filedata{ Filename: fdirentry.Name(), Size: statr.Size(), Mtime: statr.ModTime()}
		data = append(data, fdata)
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(append(b, '\n'))
}
