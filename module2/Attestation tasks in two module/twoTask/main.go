package main

import (
	"fmt"
)

func main() {
	//	t, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var st string
	fmt.Scan(&st)

	rn := []rune(st)

	fmt.Println(len(rn))

	for i, r := range rn {

		if i != len(rn)-1 {
			fmt.Print(string(r), "*")
		} else {
			fmt.Print(string(r))
		}
	}
}
