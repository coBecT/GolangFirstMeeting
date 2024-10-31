package main

import "fmt"

func main() {
	//groupCity := map[int][]string{
	//	10:   []string{"Огни", "Дербент", "Каспийск"},         // города с населением 10-99 тыс. человек
	//	100:  []string{"Махачкала", "Грозный", "Ставрополь"},  // города с населением 100-999 тыс. человек
	//	1000: []string{"Москва", "Санкт-Петербург", "Казань"}, // города с населением 1000 тыс. человек и более
	//
	//}

	cityPopulation := map[string]int{
		"Махачкала":  900,
		"Каспийск":   60,
		"Грозный":    400,
		"Дербент":    90,
		"Ставрополь": 300,
	}
	for k, v := range cityPopulation {
		if v >= 100 && v <= 999 {
			continue
		} else {
			delete(cityPopulation, k)
		}
	}
	fmt.Print(cityPopulation)
}
