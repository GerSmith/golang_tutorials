/*
	Срезы и строки
*/

package main

import "fmt"

func main() {
	//Строку можно преобразовать в срез байт и обратно
	str := "го!"

	bytes := []byte(str)

	fmt.Println(bytes)
	// [208 179 208 190 33]

	fmt.Println(str == string(bytes))
	// true

	//Строку можно преобразовать в срез unicode-символов
	runes := []rune(str)

	fmt.Println(runes)
	// [1075 1086 33]

	fmt.Println(str == string(runes))
	// true
}
