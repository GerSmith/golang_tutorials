/*
	Тернарный оператор своими руками
*/

package main

import "fmt"

func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

func main() {
	var value int = 72

	var result = Ternary(value > 0, "positive", "negative")

	fmt.Println(result)
}
