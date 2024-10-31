package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	sl := make([]int, a)
	for i := 0; i < a; i++ {
		b := 0
		fmt.Scan(&b)
		sl[i] = b
	}
	//	minCount := sl[0]
	minZh := 9223372036854775807
	for i := 1; i < a; i++ {
		if sl[i] < minZh {
			minZh = sl[i]
		}
	}
	min := 0
	for _, item := range sl {
		if item == minZh {
			min++
		}
	}

	//for _, item := range sl {
	//	if item <= minCount {
	//		if item < minCount {
	//			min = 0
	//		}
	//		min++
	//	}
	//}
	fmt.Println(min)
}

//test 6
//6
//-2
//4
//1
//-3
//-2
//-2
//test 7
//3
//5
//4
//4
