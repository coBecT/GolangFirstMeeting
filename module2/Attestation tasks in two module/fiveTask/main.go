package main

import (
	"fmt"
	"math"
)

var k, p, v = 1296.0, 6.0, 6.0

func main() {
	fmt.Println(T())

}
func T() float64 {
	t := 6 / W()
	return t
}
func W() float64 {
	w := math.Sqrt(k / M())
	return w
}
func M() float64 {
	m := p * v
	return m
}
