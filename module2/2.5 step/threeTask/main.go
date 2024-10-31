package main

import (
	"fmt"
	"strings"
)

func main() {
	x, s := "", ""
	fmt.Scan(&x, &s)
	//st := bufio.NewReader(os.Stdin)
	//x, _ := st.ReadString('\n')
	//s, _ := st.ReadString('\n')

	str := []rune(x)
	str2 := []rune(s)
	//	sl := []int{}
	//	ln := len(str2)
	idex := 0

	if strings.Contains(x, s) {
		for idx, i := range str {
			if i == str2[0] && str[idx+1] == str2[1] {
				idex = idx
				break
			}
		}
	}
	if idex != 0 {
		fmt.Println(idex)
	} else {
		fmt.Println("-1")
	}

	//for i := 0; i < len; i++ {
	//	for _, d := range str {
	//		if d == str[i] {
	//			sl = append(sl, int(d))
	//		}
	//	}
	//}
}
