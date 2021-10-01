package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = w.Write([]byte("Отображение заметки..."))

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
