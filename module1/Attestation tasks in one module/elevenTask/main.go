package main

import "fmt"

const (
	kr1 = "korova"
	kr2 = "korovy"
	kr3 = "korov"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 && n < 100 {
		if n == 1 || n%10 == 1 {
			fmt.Println(n, kr1)
		} else if (n%10 >= 2 && n%10 <= 4) && (n%100 < 12 || n%100 > 14) {
			fmt.Println(n, kr2)
		} else {
			fmt.Println(n, kr3)
		}
	}
}
