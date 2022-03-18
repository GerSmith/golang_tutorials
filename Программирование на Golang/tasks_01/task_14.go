/*
	Программирование на Golang. Глава 1. Тесты. Дано натуральное число N. Выведите его представление в двоичном виде.
*/

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Printf("%b", n)
}
