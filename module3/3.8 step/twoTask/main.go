package main

import "time"

func main() {
	ch := make(chan int)
	go task(ch, 1)
	time.Sleep(2 * time.Second)
}
func task(ch chan int, n int) {
	ch <- n + 1
}
