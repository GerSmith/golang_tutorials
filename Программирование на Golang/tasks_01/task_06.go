/*
	Программирование на Golang. Глава 1. Тесты. Даны два числа. Найти их среднее арифметическое.
*/

package main

import "fmt"

func main() {
	var a, b int

	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	average := float32(a+b) / 2.0
	fmt.Println(average)

}
