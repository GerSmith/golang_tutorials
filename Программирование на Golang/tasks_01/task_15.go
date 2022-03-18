/*
	Программирование на Golang. Глава 1. Тесты. Из натурального числа удалить заданную цифру.
*/

package main

import "fmt"

func reverseUints(input []uint) []uint {
	//reverse Uints slice
	i, j := 0, len(input)-1
	for i < j {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}
	return input
}

func main() {
	var num, skip uint

	numbers := make([]uint, 0)

	// fill the array
	for fmt.Scan(&num); num > 0; num /= 10 {
		numbers = append(numbers, num%10)
	}
	// reverse it
	reverseUints(numbers)

	fmt.Scan(&skip)

	for _, v := range numbers {
		if v != skip {
			fmt.Print(v)
		}
	}
}
