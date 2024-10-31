package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := ""
	fmt.Scan(&str)

	rn := []rune(str)
	ln := utf8.RuneCountInString(str)
	tr := false
	for i := 0; i < ln/2; i++ {
		//		fmt.Println(string(rn[i]), string(rn[ln-i-1]))
		if rn[i] != rn[ln-i-1] {
			tr = false
			break
		} else {
			tr = true
		}
	}

	if tr {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

//топот
//толос
