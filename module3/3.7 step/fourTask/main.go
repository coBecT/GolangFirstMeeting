package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const now = 1589570165

func main() {
	//Считываем 2 числа минуты и секунды в 2 разные переменные
	buf := bufio.NewReader(os.Stdin)

	minD, err := buf.ReadString('.')
	if err != nil {
		fmt.Println("error min")
	}
	minD = strings.TrimSuffix(minD, ".")
	secD, err := buf.ReadString('.')
	secD = strings.TrimSuffix(secD, ".")
	rn1 := []rune(minD)
	rn2 := []rune(secD)

	str1 := ""
	for _, v := range rn1 {
		if unicode.IsDigit(v) {
			str1 += string(v)
		}
	}
	min, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("error min atoi")
	}
	str2 := ""
	for _, v := range rn2 {
		if unicode.IsDigit(v) {
			str2 += string(v)
		}
	}
	sec, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("error sec atoi")
	}
	if err != nil {
		fmt.Print("error sec")
	}
	base := time.Unix(now, 0).UTC()
	minInSec := min * 60
	secRes := sec + minInSec
	res := strconv.Itoa(secRes) + "s"
	var dr time.Duration
	dr, err = time.ParseDuration(res)
	result := base.Add(dr)
	fmt.Println(result.Format(time.UnixDate))
}

//if err != nil {
//	fmt.Print("Error reading ->", err)
//}
//min, err :=
//dur, err := time.ParseDuration(buf)
//if err != nil {
//	fmt.Print("Error parsing to duration ->", err)
//}
//baseTime := time.Unix(now, 0)
//12 мин. 13 сек.
