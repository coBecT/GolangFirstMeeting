package main

import "fmt"

// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
//
// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.
//
// Формат выходных данных
// Выведите одно целое число - первую цифру заданного числа.
//
// Sample Input:
//
// 1234
// Sample Output:
//
// 1
func main() {
	var ch = 0
	fmt.Scan(&ch)
	if ch < 10000 {
		switch {
		case ch > 1000:
			fmt.Println(ch / 1000)
		case ch > 100:
			fmt.Println(ch / 100)
		case ch > 10:
			fmt.Println(ch / 10)

		}
	}
}
