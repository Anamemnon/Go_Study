package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, myChan chan Page) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	myChan <- Page{URL: url, Size: len(body)}
}

func main() {
	myChannel := make(chan Page)

	go responseSize("https://example.com", myChannel)
	go responseSize("https://golang.org/doc", myChannel)
	go responseSize("https://golang.org", myChannel)

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
