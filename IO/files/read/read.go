package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		// wr.WriteString(sc.Text())
		fmt.Println(sc.Text())
	}

	fmt.Println()
	fmt.Println("### ReadFile ###")

	// для более мелких файлов
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}
