package main

import "fmt"

const dva = 2

func main() {
	var N, vr, ostatok int
	fmt.Scan(&N)
	sl1 := []int{}
	vr = N
	for {
		ostatok = vr % dva
		vr = vr / dva

		sl1 = append(sl1, ostatok)
		if vr == 1 || vr == 0 {
			sl1 = append(sl1, 1)
			break
		}
	}
	sl2 := make([]int, len(sl1))
	idx := len(sl1) - 1
	for i, _ := range sl1 {
		sl2[i] = sl1[idx]
		idx--
	}
	for _, i := range sl2 {
		fmt.Printf("%d", i)
	}

}
