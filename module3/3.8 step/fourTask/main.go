package main

import (
	"time"
)

func main() {
	ch, ch2 := make(chan string), make(chan string)
	go removeDuplicates(ch, ch2)
	time.Sleep(2 * time.Second)
}

func removeDuplicates(inputStream, outputStream chan string) {
	var str string

	for val := range inputStream {
		if val != str {
			outputStream <- val
			str = val
		}
	}

	close(outputStream)
}

//нам надо сравнивать прошлое введеное значение с новым и если они одинаковые то останавливать выполнение функции и закрывать канал

//

//Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
//Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд.
//Не забудьте закрыть канал ;)
//Функция должна называться removeDuplicates()
//Выводить или вводить ничего не нужно!
