/*
	Программирование на Golang. Глава 1. Тесты. Дано трехзначное число. Найдите сумму его цифр.
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	// тестовый вывод
	// fmt.Println(num)
	// решение в лоб
	first_digit := num % 10
	second_digit := num / 10 % 10
	third_digit := num / 100 % 10
	fmt.Println(first_digit + second_digit + third_digit)
	// решение через цикл
	sum := 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	fmt.Println(sum)
	// решение через scanf
	var a, b, c int
	fmt.Scanf("%1d%1d%1d", &a, &b, &c)
	fmt.Print(c + b + a)
}
