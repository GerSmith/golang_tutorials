package main

import "fmt"

func main() {
	const (
		exists    string = "Существует"
		notExists string = "Не существует"
	)

	var a, b, c int

	_, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	if c < a+b && b < a+c && a < b+c {
		fmt.Println(exists)
	} else {
		fmt.Println(notExists)
	}
}
