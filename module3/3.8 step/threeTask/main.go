package main

import (
	"fmt"
	"time"
)

//Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
//
//Функция должна называться task2().
//
//Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

func main() {
	ch := make(chan string, 5)
	go task2(ch, "Privet")
	time.Sleep(2 * time.Second)
	fmt.Println(<-ch)
}
func task2(ch chan string, str string) {

	for i := 0; i < 5; i++ {

		ch <- str + " "

	}

}
