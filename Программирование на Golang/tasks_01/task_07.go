/*
	Программирование на Golang. Глава 1. Тесты. По данным числам, определите количество чисел, которые равны нулю.  
*/

package main

import "fmt"

func main() {
	var amt, num int

	zero_cnt := 0

	for fmt.Scan(&amt); amt > 0; amt-- {
		if fmt.Scan(&num); num == 0 {
			zero_cnt++
		}
	}

	fmt.Println(zero_cnt)
}
