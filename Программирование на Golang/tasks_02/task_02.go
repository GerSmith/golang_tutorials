/*
	Программирование на Golang. Глава 2. Тесты.

	Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var str string

	_, err := fmt.Scan(&str)
	if err != nil {
		log.Println("error:", err)
		return
	}
	
	fmt.Println(strings.Join(strings.Split(str, ""), "*"))
}
