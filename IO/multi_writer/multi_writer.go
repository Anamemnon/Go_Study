package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buf1 := bytes.NewBuffer([]byte{})
	buf2 := bytes.NewBuffer([]byte{})

	wr := io.MultiWriter(buf1, buf2)
	_, err := io.WriteString(wr, "hello world")

	if err != nil {
		panic(err)
	}

	fmt.Println("buf1 =", buf1)
	fmt.Println("buf2 =", buf2)
}
