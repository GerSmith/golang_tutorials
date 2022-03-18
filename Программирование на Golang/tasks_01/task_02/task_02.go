/*
	Программирование на Golang. Глава 1. Тесты.	
	Дано трехзначное число. Переверните его, а затем выведите.
*/

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanf("%1d%1d%1d", &a, &b, &c)
	fmt.Printf("%d%d%d", c, b, a)

}
