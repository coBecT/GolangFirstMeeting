package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
func sumInt(a ...int) (int, int) {
	sum := 0
	idx := 0
	for _, i := range a {
		sum += i
		idx++
	}
	return idx, sum
}
