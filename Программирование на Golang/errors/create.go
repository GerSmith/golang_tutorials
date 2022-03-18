/*
	Программирование на Golang. Обработка ошибок
*/

package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	err := errors.New("my error")
	fmt.Println("", err)

	d, err := divide(2, 0)
	if err != nil {
		fmt.Println("error:", err) // Функция вернула непустую ошибку
	} else {
		fmt.Println(d) // Деление прошло без ошибок
	}
}
