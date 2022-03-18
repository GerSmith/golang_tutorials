/*
	Программирование на Golang. Строки. Задания

	На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(runed []rune) bool {
	i, j := 0, len(runed)-1
	for i < j {
		if runed[i] == runed[j] {
			i++
			j--
		} else {
			return false
		}

	}
	return true
}

func main() {
	// ввод
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// обработку ошибок пропускаем
	// создаём срез из рун
	rs := []rune(text)
	// обрезка
	rs = rs[:len(rs)-2]

	if isPalindrome(rs) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
