/*
	Программирование на Golang. Строки. Задания

	На вход подается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
*/

package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	for idx, elem := range str {
		if idx%2 != 0 {
			fmt.Print(string(elem))
		}
	}
}
