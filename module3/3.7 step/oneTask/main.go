package main

import (
	"fmt"
	"time"
)

const unixDate = "Mon Jan _2 15:04:05 MST 2006"

func main() {
	//data, err := bufio.NewReader(os.Stdin).ReadString('\n')

	var data string
	//if err != nil {
	//	fmt.Println("Error data ->", err.Error())
	//}
	fmt.Scan(&data)
	lay := time.RFC3339

	out, err := time.Parse(lay, data)
	if err != nil {
		fmt.Println("Error errTime ->", err.Error())
	}

	vv := out.Format(unixDate)
	fmt.Println(vv)
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// Строка, представляющая дату и время
//	//	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
//	//	input := "1986-04-16T05:20:00+06:00"
//	var input string
//	fmt.Scan(&input)
//	// Параметры для парсинга строки даты
//	layout := time.RFC3339
//
//	// Парсим строку в тип Time
//	t, err := time.Parse(layout, input)
//	if err != nil {
//		fmt.Println("Ошибка при парсинге времени:", err)
//		return
//	}
//
//	// Форматируем дату в нужный вид
//	unixDateFormat := "Mon Jan 02 15:04:05 -0700 2006"
//	output := t.Format(unixDateFormat)
//
//	// Печатаем результат
//	fmt.Println(output)
//}
