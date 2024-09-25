package main

import "fmt"

func main() {
	year := 0
	fmt.Scan(&year)

	switch {
	case year%400 == 0:
		fmt.Println("YES")
	case year%4 == 0 && year%100 != 0:
		fmt.Println("YES")
	default:
		fmt.Println("NO")

	}
}

//package main
//
//import "fmt"
//
//func main() {
//	s := 0
//	fmt.Scan(&s)
//	if s <= 10000 && s > 0 {
//		if s%400 == 0 || s%4 == 0 && s%100 != 0 {
//			fmt.Println("YES")
//		} else {
//			fmt.Println("NO")
//		}
//	} else {
//		fmt.Println()
//	}
//}
