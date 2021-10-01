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

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, err := w.Write([]byte("Привет"))

	if err != nil {
		panic(err)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Отображение заметки..."))

	if err != nil {
		panic(err)
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, fmt.Sprintf("%s-метод запрещен", r.Method), http.StatusMethodNotAllowed)
		return
	}
	_, err := w.Write([]byte("Форма для создания новой заметки..."))

	if err != nil {
		panic(err)
	}
}