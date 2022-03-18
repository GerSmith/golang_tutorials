/*
	Программирование на Golang. Строки. Задания

	На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// читаем строку из STDIN
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strings.TrimSpace(text)
	//fmt.Printf("%v - %T", text, text)
	// преобразуем в срез рун
	r := []rune(text)
	//
	//fmt.Println(r[0], r[len(r)-1], r[len(r)-2])
	//fmt.Println(string(r[0]), string(r[len(r)-1]), string(r[len(r)-2]))
	//for idx, elem := range r {
	//	fmt.Println(idx, elem, string(elem))
	//}
	// провекра условия в соответсвие с задание
	if unicode.IsUpper(r[0]) && r[len(r)-3] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
