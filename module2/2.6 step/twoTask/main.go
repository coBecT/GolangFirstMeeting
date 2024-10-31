package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		sum, erro := divide(a, b)
		if erro != nil {
			fmt.Print(erro)
		} else {
			fmt.Print(sum)
		}
	}
}
func divide(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("ошибка")
	}
	sum := a / b
	return sum, nil
}

//func divide(a int, b int) (int, error) {
// if b == 0 {
//  return 0, errors.New("ошибка")
// }
// return a / b, nil
//}
//
//func main() {
// var a,b int
// fmt.Scan(&a,&b)
// result, err := divide(a, b)
// if err != nil {
//  fmt.Println(err)
// } else {
//  fmt.Println(result)
// }
//}
