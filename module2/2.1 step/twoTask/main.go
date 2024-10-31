package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	a, b, c, d, mi := 0, 0, 0, 0, 8349506
	fmt.Scan(&a, &b, &c, &d)
	sl := []int{a, b, c, d}
	for i := 0; i < 4; i++ {
		if mi > sl[i] {
			mi = sl[i]
		}
	}
	return mi
}
