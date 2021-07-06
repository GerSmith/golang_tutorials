package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите значение а")
	fmt.Scan(&a)
	fmt.Println("Введите значение b")
	fmt.Scan(&b)

	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}

	switch a {
	case 0:
		fmt.Println("a равно 0")
	case 1:
		fmt.Println("a равно 1")
	case 2:
		fmt.Println("a равно 2")
	case 3:
		fmt.Println("a равно 3")
	case 4:
		fmt.Println("a равно 4")
	case 5:
		fmt.Println("a равно 5")
	case 6:
		fmt.Println("a равно 6")
	case 7:
		fmt.Println("a равно 7")
	case 8:
		fmt.Println("a равно 8")
	case 9:
		fmt.Println("a равно 9")
	default:
		fmt.Println("a больше 9")
	}
}
