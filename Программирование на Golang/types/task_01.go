/*
	Программирование на Golang. Преобразование типов данных
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func adding(a, b string) int64 {
	a_runes := []rune(a)
	b_runes := []rune(b)

	var a_clear, b_clear string

	for _, v := range a_runes {
		if unicode.IsDigit(v) {
			a_clear += string(v)
		}
	}

	for _, v := range b_runes {
		if unicode.IsDigit(v) {
			b_clear += string(v)
		}
	}

	a_i, err := strconv.Atoi(a_clear)
	if err != nil {
		log.Println(err)
		return 0
	}

	b_i, err := strconv.Atoi(b_clear)
	if err != nil {
		log.Println(err)
		return 0
	}

	return int64(a_i + b_i)
}

func main() {
	a := "%^80"
	b := "hhhhh20&&&&nd"
	//test output
	//fmt.Printf("a = %v; a is %T\n", a, a)
	//fmt.Printf("b = %v; b is %T\n", b, b)

	result := adding(a, b)
	fmt.Printf("result = %v; result is %T", result, result)

}
