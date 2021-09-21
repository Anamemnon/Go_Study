package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Stdout\n")
	io.WriteString(os.Stderr, "Stderr\n")

	buf := []byte{0xC2, 0xBE, 0x0020, 0xC6, 0xB5, 0x0020}
	for i := 0; i < 3; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// Пакет fmt также можно использовать
	fmt.Fprintln(os.Stdout, "\nFprintln Stdout")
}
