/*
	Программирование на Golang. Глава 2. Тесты.

	Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var str string

	_, err := fmt.Scan(&str)
	if err != nil {
		log.Println("error", err)
		return
	}

	var max int
	for _, el := range strings.Split(str, "") {
		el, _ := strconv.Atoi(el)
		if err != nil {
			log.Println("error", err)
			return
		}
		if max < el {
			max = el
		}
	}
	fmt.Println(max)
}
