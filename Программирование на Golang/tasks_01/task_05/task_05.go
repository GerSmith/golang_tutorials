/*
	Программирование на Golang. Глава 1. Тесты. Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
*/

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	// fmt.Println(a, b, c)

	condition := a+b > c && a+c > b && b+c > a

	if condition {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
