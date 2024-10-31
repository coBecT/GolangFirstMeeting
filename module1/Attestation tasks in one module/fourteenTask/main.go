package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	var del string
	fmt.Scan(&del)

	sl := make([]string, len(n))
	for i, char := range n {
		if string(char) != del {
			sl[i] = string(char)
		} else {
			continue
		}
	}
	for _, item := range sl {
		fmt.Print(item)
	}

}
