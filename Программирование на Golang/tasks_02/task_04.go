/*
	Программирование на Golang. Глава 2. Тесты.

	На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(strings.Split(str, ""))
	for _, value := range strings.Split(str, "") {
		value, err := strconv.Atoi(value)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(math.Pow(float64(value), 2.0))
	}
}
