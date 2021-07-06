package main

/*
   Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/

import "fmt"

func main() {
	var count, max, n int = 0, 0, 0

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			count = 1
		} else if n == max {
			count += 1
		}
	}

	fmt.Println(count)
}
