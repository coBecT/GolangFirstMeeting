package main

import "fmt"

func main() {
	var lnArr int
	fmt.Scan(&lnArr)
	if 1 <= lnArr && lnArr <= 100 {

		arr := make([]int, lnArr)

		for i := 0; i < lnArr; i++ {
			a := 0
			fmt.Scan(&a)
			arr[i] = a
		}
		for i := 0; i < lnArr; i++ {
			if i%2 == 0 {
				fmt.Print(arr[i], " ")
			}
		}
	}
}
