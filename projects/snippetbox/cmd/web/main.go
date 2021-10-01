package main

import (
	"fmt"
	"log"
	"net/http"
)
const (
	PORT = 4000
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("Запуск веб-сервера на http://127.0.0.1:%d\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
	log.Fatal(err)
}
