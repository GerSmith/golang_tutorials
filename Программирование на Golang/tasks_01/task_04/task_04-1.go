package main

import "fmt"

func main() {
	val, a, b, c := map[bool]string{
		true:  "Прямоугольный",
		false: "Непрямоугольный",
	}, 0, 0, 0
	fmt.Scan(&a, &b, &c)
	fmt.Println(val[a*a+b*b == c*c])
}
