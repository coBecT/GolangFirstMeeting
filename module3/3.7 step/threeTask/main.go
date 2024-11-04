package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const layout = "02.01.2006 15:04:05"

func main() {
	r := bufio.NewReader(os.Stdin)
	str1, _ := r.ReadString(',')
	str1 = strings.TrimSuffix(str1, ",")

	str2, _ := r.ReadString('\n')
	str2 = strings.TrimSuffix(str2, "\n")

	//  13.03.2018 14:00:15,12.03.2018 14:00:15
	time1, err := time.Parse(layout, str1)

	if err != nil {
		fmt.Printf("Error time1 parsing -> %s\n", err.Error())
		panic(err)
		return
	}
	time2, err := time.Parse(layout, str2)
	if err != nil {
		fmt.Printf("Error time2 parsing -> %s\n", err.Error())
		return
	}

	if time1.After(time2) {
		//Если time1 произошло после time2
		difference := time1.Sub(time2)
		fmt.Print(difference)
	} else if time1.Before(time2) {
		//Если time1 произошло до time2
		difference := time2.Sub(time1)
		fmt.Print(difference)
	} else if time1.Equal(time2) {
		//Если time1 равно time2
		fmt.Print("")
	}

}
