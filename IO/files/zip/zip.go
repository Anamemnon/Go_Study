package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var buff bytes.Buffer

	zipW := zip.NewWriter(&buff)
	f, err := zipW.Create("new_file.txt")

	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte("Содержимое файла"))

	if err != nil {
		panic(err)
	}

	err = zipW.Close()

	if err != nil {
		panic(err)
	}

	// Запись данных в архив
	err = ioutil.WriteFile("data.zip", buff.Bytes(), os.ModePerm)

	if err != nil {
		panic(err)
	}

	// Распаковка содержимого архива
	zipR, err := zip.OpenReader("data.zip")

	if err != nil {
		panic(err)
	}

	for _, file := range zipR.File {
		fmt.Println("Файл " + file.Name + " содержит следующее:")
		r, err := file.Open()

		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(os.Stdout, r)

		if err != nil {
			panic(err)
		}

		err = r.Close()

		if err != nil {
			panic(err)
		}

		fmt.Println()
	}
}
