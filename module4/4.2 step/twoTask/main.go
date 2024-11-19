package main

import (
	"fmt"
	"io"
	"net/http"
)

// Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get и напечатайте ответ сервера (только тело).
func main() {
	path := "http://127.0.0.1:5555/get"
	ht, err := http.Get(path)
	if err != nil {
		fmt.Println("error http", err)
	}

	defer ht.Body.Close()

	resp, err := io.ReadAll(ht.Body)
	if err != nil {
		fmt.Println("error io", err)
	}

	fmt.Println(string(resp))
}
