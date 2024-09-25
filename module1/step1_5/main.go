package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)

	if a > 0 && a <= 360 {
		hour := 0
		minute := 0
		if a%30 == 0 {
			hour = a / 30
			fmt.Printf("It is %d hours %d minutes.", hour, minute)
		} else if a%30 > 0 {

			hour = a / 30
			minute = (a % 30) * 2
			fmt.Printf("It is %d hours %d minutes.", hour, minute)
		}
	}
}
