package main

import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(u uint) uint {
		var res uint
		var re uint64
		var tm string

		str := strconv.FormatUint(uint64(u), 10)
		run := []rune(str)
		for _, v := range run {
			intik, err := strconv.Atoi(string(v))
			if err != nil {
			}
			if string(v) == "0" || intik%2 != 0 {
				continue
			} else {
				//	re, err = strconv.ParseUint(string(v), 10, 64)
				tm += string(v)

				//if err != nil {
				//}
				//
				//res += uint(re)
			}
		}
		re, err := strconv.ParseUint(tm, 10, 64)
		if err != nil {
		}
		res = uint(re)

		if res == 0 {
			u = 100
		} else if tm == "" {
			u = 100
		} else {
			u = res
		}
		return u
	}

	fmt.Println(fn)
}
