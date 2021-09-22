package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const FILE_NAME_R = "./temp/file.txt"
const FILE_NAME_W = "./temp/new_file.txt"

func main() {
	f, err := os.Open(FILE_NAME_R)

	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile(FILE_NAME_W, os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		panic(err)
	}

	io.WriteString(f, "File has been created.")
	f.Close()
}
