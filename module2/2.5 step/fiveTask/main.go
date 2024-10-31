// package main
//
// import (
//
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//
// )
//
//	func main() {
//		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//
//		sl := []rune(str)
//		ps := ""
//
//		for i := 0; i < len(sl)-1; i++ {
//			if sl[i] == sl[i+1] {
//				ps += string(sl[i])
//				fmt.Println(sl[i])
//			}
//		}
//
//		fmt.Println(ps)
//		fmt.Print(strings.Trim(str, ps))
//
// }
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
//		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//		result := removeDuplicateChars(str)
//		fmt.Println(result)
//	}
//
//	func removeDuplicateChars(s string) string {
//		counts := make(map[rune]int)
//
//		// Считаем количество вхождений каждого символа
//		for _, char := range s {
//			counts[char]++
//		}
//
//		// Формируем новую строку, добавляя только уникальные символы
//		var result []rune
//		for _, char := range s {
//			if counts[char] == 1 {
//				result = append(result, char)
//			}
//		}
//
//		return string(result)
//	}
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, newS string
	fmt.Scan(&s)

	for _, v := range s {
		if strings.Count(s, string(v)) == 1 {
			newS += string(v)
		}
	}
	fmt.Println(newS)
}
