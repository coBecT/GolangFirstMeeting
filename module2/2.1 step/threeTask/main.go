package main

import "fmt"

func main() {
	fmt.Println(vote(0, 0, 1))
}
func vote(x int, y int, z int) int {
	sl := []int{x, y, z}
	nol, odin := 0, 0
	for _, v := range sl {
		if v == 0 {
			nol++
		} else {
			odin++
		}
	}
	if nol > odin {
		return 0
	} else {
		return 1
	}
}
