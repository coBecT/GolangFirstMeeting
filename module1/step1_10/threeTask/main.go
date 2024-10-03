package main

import "fmt"

//Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
//
//Sample Input:
//
//5
//38 24 800 8 16
//Sample Output:
//
//40

func main() {
	var sum, oneVar, twoVar int
	fmt.Scan(&oneVar)
	for i := 1; oneVar >= i; i++ {
		fmt.Scan(&twoVar)
		if twoVar%8 == 0 && twoVar > 9 && twoVar < 100 {
			sum += twoVar
		}
	}
	fmt.Println(sum)
}

//package main
//
//import "fmt"
//
//func main() {
//	var a int
//	fmt.Scan(&a)
//	var res int
//	for i := 1; i <= a; i++ {
//		var pr int
//		fmt.Scan(&pr)
//		if pr/10 < 10 && pr/10 > 0 && pr%8 == 0 {
//			res += pr
//		} else {
//			res += 0
//		}
//
//	}
//
//	fmt.Println(res)
//}
