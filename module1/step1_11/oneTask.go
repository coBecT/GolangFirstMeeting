package main

import "fmt"

func main() {
	var fl float64
	fmt.Scan(&fl)
	if fl > 10000 {
		fmt.Printf("%e", fl)
	} else if fl <= 0 {
		fmt.Printf("число %2.2f не подходит ", fl)
	} else {
		fl *= fl
		fmt.Printf("%.4f", fl)
	}
}
