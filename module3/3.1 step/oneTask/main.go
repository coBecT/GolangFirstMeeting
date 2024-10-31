package main

import "fmt"

func main() {
	a := make([]int, 10)
	b := 0
	for i := 0; i < 10; i++ {
		fmt.Scan(&b)
		a[i] = b
	}
	cache := make(map[int]int)
	for i := 0; i < 10; i++ {
		num := a[i]

		if val, ok := cache[num]; ok {
			fmt.Print(val, " ")
		} else {
			res := work(num)
			cache[num] = res
			fmt.Print(res, " ")
		}
	}
}
func work(x int) int {
	return 0
}
