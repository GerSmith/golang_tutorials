/*
	Программирование на Golang. Глава 1. Тесты. Найдите количество минимальных элементов в последовательности.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var len int

	fmt.Scan(&len)
	arr := make([]int, len, len)

	for i := 0; i < len; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)
	min_val, cnt := arr[0], 0

	for _, val := range arr {
		if val == min_val {
			cnt++
		}
	}

	fmt.Println(cnt)
}
