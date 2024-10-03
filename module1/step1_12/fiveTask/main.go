package main

import "fmt"

func main() {
	lnArr := 0
	fmt.Scan(&lnArr)
	item := 0
	if 1 <= lnArr && lnArr <= 100 {
		a := 0
		arr := make([]int, lnArr)
		for i := 0; i < lnArr; i++ {
			fmt.Scan(&a)
			arr[i] = a
		}
		for i := 0; i < lnArr; i++ {
			if arr[i] > 0 {
				item++
			}
		}
		fmt.Println(item)
	}
}
