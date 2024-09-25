package main

import "fmt"

//Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
//
//
//
//Формат входных данных
//
//На вход подается номер билета - одно шестизначное  число.
//
//Формат выходных данных
//Выведите "YES", если билет счастливый, в противном случае - "NO".
//
//Sample Input:
//
//613244
//Sample Output:
//
//YES

func main() {
	sum, oneCh, twoCh := 0, 0, 0
	fmt.Scan(&sum)
	oneCh = sum / 1000
	twoCh = sum % 1000

	fmt.Println(oneCh)
	fmt.Println(twoCh)

	sumOne := oneCh/100 + (oneCh%100)/10 + oneCh%10
	sumTwo := twoCh/100 + (twoCh%100)/10 + twoCh%10
	switch sumOne == sumTwo {
	case true:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}

//package main
//
//import "fmt"
//
//func main() {
//	s, ch, res := 0, 0, 0
//	fmt.Scan(&s)
//	ch = s / 1000
//	pr, tw, thr := ch/100, ch/10, ch/1
//	tw = tw % 10
//	thr = thr % 10
//	//fmt.Println(pr, tw, thr)
//	res = pr + tw + thr
//	//fmt.Println(res)
//	ch = s % 1000
//
//	pr, tw, thr = ch/100, ch/10, ch/1
//	tw = tw % 10
//	thr = thr % 10
//	res2 := pr + tw + thr
//	//fmt.Println(res2)
//	if res == res2 {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}
