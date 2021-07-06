package main

import "fmt"

func main() {

	var v int = 12

	switch {
	case v > 10:
		fmt.Println("v > 10")
		fallthrough
	case v == 12:
		fmt.Println("v == 12? но case сработал только потому, что есть fallthrough")
	default:
		fmt.Println("а этого условия может и не быть")
	}

	var x interface{} = 12

	switch i := x.(type) { // проверим тип интерфейса
	case int:
		fmt.Println("в х действительно лежит число", i)
	case string:
		fmt.Println("а не строка")

	}
}
