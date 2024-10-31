package main

import "fmt"

type battery struct {
	capacity string
}

func (b battery) String() string {
	res := "["
	count := 0
	for _, v := range b.capacity {
		if string(v) == "1" {
			count++
		} else {
			res += " "
		}
	}
	for i := 0; i < count; i++ {
		res += "X"
	}
	res += "]"
	return res
}
func main() {
	a := ""
	fmt.Scan(&a)
	batteryForTest := battery{capacity: a}

	// Ваша функция, принимающая Battery на стандартный вывод
	funThatPrintsBattery(batteryForTest)
}
func funThatPrintsBattery(b fmt.Stringer) {
	fmt.Println(b.String())
}
