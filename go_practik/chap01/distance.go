/*
	Напишите программу,
	которая рассчитает евклидово расстояние
	между точками на плоскости (x1, y1) и (x2, y2)
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		x1, y1, x2, y2 float64
	)

	fmt.Scan(&x1, &y1, &x2, &y2)

	fmt.Printf("x1 = %.1f, y1 = %.1f, x2 = %.1f, y2 = %.1f\n", x1, y1, x2, y2)

	length := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

	fmt.Println(length)
}
