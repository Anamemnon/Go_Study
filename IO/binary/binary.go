package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// Запись бинарных значений
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}

	const GREETING = "Hello"

	if err := binary.Write(buf, binary.BigEndian, []byte(GREETING)); err != nil {
		panic(err)
	}

	// Чтение записанных значений
	var number float64

	if err := binary.Read(buf, binary.BigEndian, &number); err != nil {
		panic(err)
	}

	fmt.Printf("%T %f\n", number, number)

	hello := make([]byte, len(GREETING))

	if err := binary.Read(buf, binary.BigEndian, &hello); err != nil {
		panic(err)
	}

	fmt.Printf("%T %s\n", string(hello), string(hello))

}
