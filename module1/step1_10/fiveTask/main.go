package main

import "fmt"

func main() {

	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	if a < 10000 {
		a = a / 100 / 10
		fmt.Println(a)
	} else {
		fmt.Println("Вы превысили число, не должно превышать 10000")
	}
	//const one = 10000
	//var n, c, d int
	//fmt.Scan(&n, &c, &d)
	//if n < one && c < one && d < one {
	//	for i := 1; i <= n; i++ {
	//		if i%c == 0 && i%d != 0 {
	//			fmt.Print(i)
	//			break
	//		}
	//	}
	//}
}

//package main
//
//import "fmt"
//
//func main() {
//	var n, c, d int
//	fmt.Scan(&n, &c, &d)
//	if n < 10000 && c < 10000 && d < 10000 {
//		for i := 1; i <= n; i++ {
//			if i%c == 0 && i%d != 0 {
//				fmt.Println(i)
//				break
//			} else {
//			}
//		}
//	}
//}
