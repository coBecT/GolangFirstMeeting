package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)

	a, b := 1, 1
	idx := 2
	for b < A {
		a, b = b, a+b
		idx++

	}
	if b == A {
		fmt.Println(idx)
	} else {
		fmt.Println("-1")
	}
}
