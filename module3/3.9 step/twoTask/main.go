package main

import "sync"

func main() {
	var wg sync.WaitGroup
	//	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// func work() don`t release
			//			work()
			defer wg.Done()
		}()
	}
	wg.Wait()
}
