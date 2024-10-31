package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r := []rune(text)
	//ln := utf8.RuneCountInString(string(r))
	//fmt.Println(strings.HasSuffix(string(r), "."))
	//
	//if unicode.IsUpper(r[0]) {
	//	if strings.HasSuffix(text, ".") {
	//		fmt.Println("Right")
	//	}
	//}
	if unicode.IsUpper(r[0]) && r[len(r)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	//c := 0
	//for i := range r {
	//	if unicode.IsUpper(r[0]) {
	//		if r[i] == '.' {
	//			fmt.Println("Right")
	//			c++
	//		}
	//	}
	//}
	//if c != 1 {
	//	fmt.Println("Wrong")
	//}

}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"unicode"
//)
//
//func main() {
//
//	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//	r := []rune(text)
//	//ln := utf8.RuneCountInString(string(r))
//	//fmt.Println(strings.HasSuffix(string(r), "."))
//	//
//	//if unicode.IsUpper(r[0]) {
//	//	if strings.HasSuffix(string(r), ".") {
//	//
//	//	}
//	//}
//	//
//	c := 0
//	for i := range r {
//		if unicode.IsUpper(r[0]) {
//			if r[i] == '.' {
//				fmt.Println("Right")
//				c++
//			}
//		}
//	}
//	if c != 1 {
//		fmt.Println("Wrong")
//	}
//
//}
