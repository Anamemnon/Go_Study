package main

import (
	"fmt"
	"github.com/Anamemnon/Go_Study/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2010)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(date.Year(), date.Month(), date.Day())
}
