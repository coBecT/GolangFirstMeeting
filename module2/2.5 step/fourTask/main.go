package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	read, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	len := utf8.RuneCountInString(read)

	rn := []rune(read)
	ad := []rune{}

	for i := 0; i < len; i++ {
		if i%2 == 1 {
			ad = append(ad, rn[i])
		}
	}
	fmt.Print(string(ad))
}
