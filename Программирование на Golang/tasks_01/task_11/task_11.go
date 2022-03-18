/*
	Программирование на Golang. Глава 1. Тесты. По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
*/

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n == 1 {
		fmt.Printf("%d korova", n)
	} else if n > 1 && n <= 4 {
		fmt.Printf("%d korovy", n)
	} else if n > 4 && n <= 20 {
		fmt.Printf("%d korov", n)
	} else if n > 20 {
		if n%10 == 1 {
			fmt.Printf("%d korova", n)
		} else if n%10 == 2 || n%10 == 3 || n%10 == 4 {
			fmt.Printf("%d korovy", n)
		} else {
			fmt.Printf("%d korov", n)
		}
	}
}
