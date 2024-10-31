package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	cratnoe := -1
	if a <= b {
		for i := a; i <= b; i++ {
			if i%7 == 0 {
				cratnoe = i
			}
		}
	}
	if cratnoe == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(cratnoe)
	}

}
