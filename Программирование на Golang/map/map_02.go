/*
	Программирование на Golang. Отображения (map)
*/

package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"город1", "город2", "город3"}, // города с населением 10-99 тыс. человек
		100:  []string{"город4", "город5", "город6"}, // города с населением 100-999 тыс. человек
		1000: []string{"город7", "город8", "город9"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"город4": 101,
		"город5": 102,
		"город6": 103,
		"город1": 1,
		"город9": 1000,
	}

	for k, _ := range cityPopulation {
		isConsist := false
		for _, v := range groupCity[100] {
			if k == v {
				isConsist = true
				break
			}
		}
		if isConsist == false {
			delete(cityPopulation, k)
		}
	}

	fmt.Println(cityPopulation)
}
