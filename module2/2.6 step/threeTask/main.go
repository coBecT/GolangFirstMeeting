package main

import "fmt"

func main() {
	var a uint32 = 5
	defer specPrint(a)
	a = 6
	fmt.Print(a)
}

func specPrint(s uint32) uint32 {
	fmt.Print(s)
	s++
	return s
}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(divide(15, 5))
//	fmt.Println(divide(4, 0))
//	fmt.Println("Program has been finished")
//}
//func divide(x, y float64) float64 {
//	if y == 0 {
//		panic("division by zero!")
//	}
//	return x / y
//}
