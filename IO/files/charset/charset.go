package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

const FILE_NAME = "example.txt"

func main() {
	// Запись строки в кодировке Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, err := encoder.String("This is sample text with runes Š")

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(FILE_NAME, []byte(s), os.ModePerm)

	f, err := os.Open(FILE_NAME)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)

	b, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
