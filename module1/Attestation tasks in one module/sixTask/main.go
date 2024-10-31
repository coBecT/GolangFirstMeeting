package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var f float64
	f = float64(a+b) / 2.0
	fmt.Print(f)
}
