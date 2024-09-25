// package main
//
// import "fmt"
//
//	func main() {
//		fmt.Println("Hello World")
//
// }
package main

import (
	"fmt"
)

func main() {
	var N int
	N = 123
	if N < 10000 {
		fmt.Print(N % 10)

	}
}

//package main
//
//import "fmt"
//
//func main() {
//	a, hour, min := 0, 0, 0
//	fmt.Scan(&a)
//	for a > 0 {
//		if a >= 30 {
//			hour++
//			a -= 30
//		} else if a < 30 {
//			min = a + a
//			a = 0
//		}
//	}
//
//	fmt.Printf("It is %d hours %d minutes.", hour, min)
//}
