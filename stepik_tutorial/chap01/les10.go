package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(getFirstNumber(a))
}

func getFirstNumber(a int) int {
	if a < 10 {
		return a
	} else {
		return getFirstNumber(a / 10)
	}
}
