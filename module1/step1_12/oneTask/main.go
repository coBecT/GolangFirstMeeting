package main

import "fmt"

func main() {
	var massCh, chetchik uint8
	var workArray [10]uint8
	//var workArray2 [10]uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&massCh)
		workArray[i] = massCh
	}
	//	fmt.Println(workArray)
	//	fmt.Println(len(workArray))
	var arrIdx [6]uint8
	for i := 0; i < 6; i++ {
		fmt.Scan(&chetchik)
		arrIdx[i] = chetchik
	}

	for i := 0; i < 6; i++ {
		index1 := arrIdx[i]
		index2 := arrIdx[i+1]

		i++

		workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}

//99 151 137 71 117 187 20 93 187 67
//1 2 3 5 7 8
//99 137 151 187 117 71 20 187 93 67
//uint8
//1 2 3 4 5 6 7 8 9 10
//1 2 3 5 7 8
