package main

import (
	"fmt"
)

func main() {
	str := 0
	fmt.Scan(&str)

	sum := 0

	for i := 0; i < 3; i++ {
		if i == 0 {
			sum += str % 10
			//fmt.Println(str % 10)
		} else if i == 1 {
			sum += (str / 10) % 10
			//fmt.Println((str / 10) % 10)
		} else if i == 2 {
			sum += (str / 100) % 10
			//fmt.Println((str / 100) % 10)
		}
	}
	fmt.Println(sum)
}
