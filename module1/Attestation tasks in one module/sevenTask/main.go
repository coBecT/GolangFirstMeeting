package main

import "fmt"

func main() {
	a := 1
	idx := 0
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		b := 0
		fmt.Scan(&b)
		if b == 0 {
			idx++
		}
	}
	fmt.Println(idx)
}
