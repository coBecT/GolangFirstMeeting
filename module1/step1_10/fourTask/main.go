package main

import "fmt"

func main() {
	var ch, maxCh, maxList int
	for {
		fmt.Scan(&ch)
		if ch == 0 {
			break
		} else if ch > maxCh {
			//if maxCh < ch {
			//	maxList++
			//}
			maxCh = ch
			maxList = 1
		} else if ch == maxCh {
			maxList++
		}
	}
	fmt.Print(maxList)
}

//package main
//func main() {
//	var max, maxCount, chislo int
//	for fmt.Scan(&chislo); chislo != 0; fmt.Scan(&chislo) {
//
//		if max <= chislo {
//			if max < chislo {
//				maxCount = 0
//			}
//			max = chislo
//
//			if max == chislo {
//				maxCount++
//			}
//		}
//	}
//	fmt.Println(maxCount)
//}
