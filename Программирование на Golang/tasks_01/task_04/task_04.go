/*
	Программирование на Golang. Глава 1. Тесты. Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	condition := a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a
	if condition {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
