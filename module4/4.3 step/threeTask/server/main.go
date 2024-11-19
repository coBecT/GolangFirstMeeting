package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Напиши веб сервер (порт :3333) - счетчик который будет обрабатывать GET (/count) и POST (/count) запросы:
//
// GET:  возвращает счетчик
// POST: увеличивает ваш счетчик на значение  (с ключом "count") которое вы получаете из формы, но если пришло НЕ число то нужно ответить клиенту: "это не число" со статусом http.StatusBadRequest (400).

var count int = 0

func main() {

	http.HandleFunc("/count", handle)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
func handle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Write([]byte(string(count)))
	}
	if r.Method == "POST" {

		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("это не число " + http.StatusText(http.StatusInternalServerError)))
		}

		res, err := strconv.Atoi(string(data))
		if err != nil {
			w.Write([]byte("это не число " + http.StatusText(http.StatusInternalServerError)))
		}
		count += res
	}
}
