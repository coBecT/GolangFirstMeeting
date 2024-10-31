package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print(maxCh())
}

func maxCh() int {
	t, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	sl := make([]int, 0)
	rn := []rune(t)
	max := 0
	for _, item := range rn {
		num, err := strconv.Atoi(string(item))
		if err != nil {
		}
		sl = append(sl, num)
	}
	for _, it := range sl {
		if max < it {
			max = it
		}
	}

	return max
}

//1112221112
//for _, item := range rn {
//	if max < int(item) {
//		fmt.Println(int(item))
//		max = int(item)
//	} else {
//		continue
//	}
//}
