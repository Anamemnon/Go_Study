package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}

func main() {
	const FILE_NAME = "data.xml"

	f, err := os.Open(FILE_NAME)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	decoder := xml.NewDecoder(f)

	// Чтение book по частям
	books := make([]Book, 0)

	for {
		token, err := decoder.Token()

		if err != nil {
			panic(err)
		}

		if token == nil {
			break
		}

		switch tp := token.(type) {
		case xml.StartElement:
			if tp.Name.Local == "book" {
				// Декодирование элемента в структуру
				var b Book
				decoder.DecodeElement(&b, &tp)
				books = append(books, b)
			}
		}
	}
	fmt.Println(books)
}
