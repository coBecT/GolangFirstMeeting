package main

import "fmt"

func main() {
	//groupCity := map[int][]string{
	//	10:   []string{"Огни", "Дербент", "Каспийск"},         // города с населением 10-99 тыс. человек
	//	100:  []string{"Махачкала", "Грозный", "Ставрополь"},  // города с населением 100-999 тыс. человек
	//	1000: []string{"Москва", "Санкт-Петербург", "Казань"}, // города с населением 1000 тыс. человек и более
	//
	//}
	groupCity := map[int][]string{
		10:   []string{"Огни", "Дербент", "Каспийск"},
		100:  []string{"Махачкала", "Грозный", "Ставрополь"},
		1000: []string{"Москва", "Санкт-Петербург", "Казань"},
	}
	cityPopulation := map[string]int{
		"Moscaov":     54,
		"Klinton":     65,
		"kirovsk":     789,
		"CPB":         345,
		"NVZ":         548,
		"Briansk":     579,
		"KIpr":        65468,
		"Harkov":      78645,
		"Novosibirsk": 57498,
	}
	//for k, v := range cityPopulation {
	//	if v >= 100 && v <= 999 {
	//		continue
	//	} else {
	//		delete(cityPopulation, k)
	//	}
	//}
	//fmt.Print(cityPopulation)

	for key := range cityPopulation {
		delete(cityPopulation, key)
	}
	for _, v := range groupCity[100] {
		cityPopulation[v] = 100
	}
	fmt.Println(len(groupCity))
	fmt.Println(cityPopulation)
}
