/*
	Программирование на Golang. Преобразование типов данных
*/

package main

import (
	"fmt"
)

func main() {
	a_str := "str"
	b_str := []byte(a_str)
	c_str := string(b_str)

	fmt.Println(a_str) // str
	fmt.Println(b_str) // [115 116 114] - побайтовый срез
	fmt.Println(c_str) // str

	a_r := "строка"
	b_r := []rune(a_r) // срез рун
	c_r := string(b_r)

	fmt.Println(a_r) // строка
	fmt.Println(b_r) // [1089 1090 1088 1086 1082 1072] - срез рун
	fmt.Println(c_r) // строка
}
