package main

func main() {

}
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	mk := make(chan int)
	select {
	case v := <-firstChan:
		mk <- v * v
	case v := <-secondChan:
		mk <- v * 3
	case <-stopChan:
		break

	}

	return mk

}
