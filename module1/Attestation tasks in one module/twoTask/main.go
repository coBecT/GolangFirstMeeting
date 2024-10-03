package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := 0
	fmt.Scan(&ch)
	if ch%10 == 0 {

	} else {
		for i := 2; i >= 0; i-- {
			str := strconv.Itoa(ch)

			fmt.Print(string(str[i]))
		}

	}
}
