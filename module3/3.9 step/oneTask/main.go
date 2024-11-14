package main

import "fmt"

func main() {
	sync := make(chan string)

	go func() {
		defer close(sync)
		sync <- work()
	}()
	fmt.Print(<-sync)
}
func work() string {
	return "Work!"
}
