/*
	массивы
*/

package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 5

	numbers[0] = 87

	fmt.Println(numbers[0]) // 87
	fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for idx, elm := range numbers {
		fmt.Printf("Элемент с индексом %d: %d\n", idx, elm)
	}

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for _, elem := range a {
		elem = 100
		fmt.Println(elem)

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [1 2 3 4 5]

	for idx := range a {
		a[idx] = 100
		fmt.Println(a[idx])

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [100 100 100 100 100]
}
