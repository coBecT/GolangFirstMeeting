package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {

	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
	fmt.Print(adding("%^8", "0h1"))
}
func adding(a, b string) int64 {
	var str1, str2 string
	r1 := []rune(a)
	r2 := []rune(b)
	for _, v := range r1 {
		if unicode.IsDigit(v) {
			str1 += string(v)
		} else {
			continue
		}
	}
	for _, v := range r2 {
		if unicode.IsDigit(v) {
			str2 += string(v)
		} else {
			continue
		}
	}
	res1, err := strconv.Atoi(str1)
	if err != nil {
	}
	res2, err := strconv.Atoi(str2)
	if err != nil {
	}
	res := int64(res1 + res2)
	return res
}
