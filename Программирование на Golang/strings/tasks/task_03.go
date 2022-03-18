/*
	Программирование на Golang. Строки. Задания

	На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var base, es string
	fmt.Scan(&base)
	fmt.Scan(&es)

	fmt.Println(strings.Index(base, es))
}
