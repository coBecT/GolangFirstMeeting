package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := ""
	fmt.Scan(&a)
	b := strings.ToLower(a)
	sl := []rune(b)
	parol := 0
	if len(b) > 5 {
		for _, item := range sl {
			if !unicode.Is(unicode.Latin, item) {
				break
			} else {
				if unicode.Is(unicode.Number, item) || unicode.Is(unicode.Latin, item) {
					parol = 1
				}
			}
		}
	}
	if parol > 0 {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}

//fdsghdfgjsdDD1
//package main
//
//import "fmt"
//
//func main() {
//	a := ""
//	fmt.Scan(&a)
//	if isValidPassword(a) {
//		fmt.Println("Ok")
//	} else {
//		fmt.Println("Wrong password")
//	}
//}
//
//func isValidPassword(password string) bool {
//	// Проверка длины пароля
//	if len(password) < 5 {
//		return false
//	}
//
//	// Проверка на соответствие требованиям (только буквы и цифры)
//	for _, char := range password {
//		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
//			return false
//		}
//	}
//
//	return true
//}
