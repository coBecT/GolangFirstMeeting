package main

import "fmt"

func main() {
	k, h, m := 0, 0, 0
	fmt.Scan(&k)
	if k > 0 && k < 86399 {
		//m = (k % 3600) / 60
		m = k / 60
		h = m / 60
		for {
			if m > 60 {
				m = m - 60
			} else {
				break
			}
		}

		fmt.Printf("It is %d hours %d minutes.", h, m)

	}
}

//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	var k int
//	fmt.Scan(&k)
//
//	hours := k / 3600
//	minutes := (k % 3600) / 60
//
//	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
//}
