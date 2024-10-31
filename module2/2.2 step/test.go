package main

import "fmt"

func printNumber(number int) {
	number = g(number)
	fmt.Println(number*3 - 10)
}
