package main

import "fmt"

func main() {
	a, b, res := 0, 0, 0
	fmt.Scan(&a, &b)
	if a <= 100 && b <= 100 && a < b {
		for i := a; i <= b; i++ {
			res += i
			if i == b {
				fmt.Print(res)
			}
		}
	}
}

//package main
//
//import "fmt"
//
//func main() {
//	a, b, res := 0, 0, 0
//	fmt.Scan(&a, &b)
//	if a < 100 && b < 100 && a < b {
//		for ; a <= b; a++ {
//			res += a
//		}
//
//	} else {
//		fmt.Print()
//	}
//	fmt.Print(res)
//}
