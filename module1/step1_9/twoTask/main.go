package main

import "fmt"

// По данному трехзначному числу определите, все ли его цифры различны.
// Формат входных данных
// На вход подается одно натуральное трехзначное число.
// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".
// Sample Input 1:
// 237
// Sample Output 1:
// YES
// Sample Input 2:
// 117
// Sample Output 2:
// NO
func main() {
	var numOne = 0
	fmt.Scan(&numOne)

	if numOne/100 == (numOne%100)/10 || numOne%10 == numOne/100 || (numOne%100)/10 == numOne/100 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
