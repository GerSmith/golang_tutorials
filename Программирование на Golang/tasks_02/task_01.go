/*
	Программирование на Golang. Глава 2. Тесты.

	На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
*/

package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	a, b int
}

func (t *Triangle) hypotenuse() float64 {
	return math.Sqrt(math.Pow(float64(t.a), 2.0) + math.Pow(float64(t.b), 2.0))
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	triangle := Triangle{a: a, b: b}
	fmt.Println(triangle.hypotenuse())
}
