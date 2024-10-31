package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, sum int
	fmt.Scan(&a)
	if a < 100000000 {
		str := strconv.Itoa(a)
		for i := 0; i < len(str); i++ {
			y, e := strconv.Atoi(string(str[i]))
			sum += y
			if e == nil {
			}
		}
	}
	sum2 := 0
	b := ""
	b = strconv.Itoa(sum)
	for i := 0; i < len(b); i++ {
		y, e := strconv.Atoi(string(b[i]))
		if e != nil {
		}
		sum2 += y
	}
	fmt.Println(sum2)
}
