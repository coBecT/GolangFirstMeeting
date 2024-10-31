package main

import (
	"fmt"
	"strconv"
)

// 1 переводим строчный символ в инт
// 2 получаем квадрат
// 3 переводим полученное число в строку
func main() {
	//	fmt.Print(powCh())
	powCh()
}

func powCh() {
	var a string
	_, _ = fmt.Scan(&a)
	run := []rune(a)
	vv := ""
	for i := 0; i < len(run); i++ {
		ln, err := strconv.Atoi(string(run[i]))
		if err != nil {
		}
		sum := ln * ln
		vv += strconv.Itoa(sum)
	}
	fmt.Println(vv)
}
