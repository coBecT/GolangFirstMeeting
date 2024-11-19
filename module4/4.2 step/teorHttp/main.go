// package main
//
// import (
//
//	"fmt"
//	"io"
//	"net/http"
//
// )
//
//	func main() {
//		conn, err := http.Get("https://www.golang.org")
//		if err != nil {
//			fmt.Println("fatal error http", err)
//		}
//		defer conn.Body.Close()
//
//		data, err := io.ReadAll(conn.Body)
//		if err != nil {
//			fmt.Println("fatal error http data", err)
//		}
//
//		fmt.Println(string(data))
//
// }
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Todo - структура для представления объекта Todo
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Создаем экземпляр структуры Todo
	todo := Todo{
		UserID:    1,
		ID:        2,
		Title:     "наш title",
		Completed: true,
	}

	// Кодируем структуру Todo в формат JSON
	jsonReq, err := json.Marshal(todo)
	if err != nil {
		log.Println(err)
		return
	}

	// URL сервера
	baseURL := "https://jsonplaceholder.typicode.com/posts/1"

	// Создаем новый HTTP-запрос с методом POST
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println("Ошибка при создании запроса:", err)
		return
	}

	// Устанавливаем заголовки запроса
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Отправляем запрос
	client := &http.Client{}        // создаем http клиент
	response, err := client.Do(req) // передаем выше подготовленный запрос на отправку
	if err != nil {
		log.Println("Ошибка при выполнении запроса: ", err)
		return
	}

	defer response.Body.Close() // не забываем закрыть тело

	// Читаем и конвертируем тело ответа в байты
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	// Конвертируем тело ответа в строку и выводим
	bodyString := string(bodyBytes)
	fmt.Printf("API ответ в форме строки: %s\n", bodyString)

	// Конвертируем тело ответа в Todo struct
	var todoStruct Todo
	err = json.Unmarshal(bodyBytes, &todoStruct)
	if err != nil {
		log.Println(err)
	}

	// Выводим структуру Todo
	fmt.Printf("API ответ в форме struct:\n%+v\n", todoStruct)

	// Вывод статуса ответа (если 200 - то успешный)
	fmt.Println("Статус ответа:", response.Status)
}
