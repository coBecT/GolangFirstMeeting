package main

import "fmt"

func main() {
	var lnSl int
	fmt.Scan(&lnSl)

	if lnSl >= 4 {
		sl := make([]int, lnSl, lnSl)
		for i := 0; i < len(sl); i++ {
			var item int
			fmt.Scan(&item)
			sl[i] = item
		}
		fmt.Printf("%d", sl[3])
	} else {
		fmt.Println("ERROR")
	}
}
