// package main
//
// import (
//
//	"bufio"
//	"fmt"
//	"os"
//
// )
//
//	func main() {
//		reader := bufio.NewReader(os.Stdin)
//		input, _ := reader.ReadString('\n')
//		input = input[:len(input)-1]
//		var result string
//
//		for i := 0; i < len(input); i++ {
//			if i%2 == 1 { // Проверяем, является ли индекс нечетным
//				result += string(input[i])
//			}
//		}
//
//		fmt.Println(result)
//	}
package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println(a + 1)
	} else {
		fmt.Println(a + 2)
	}
}
