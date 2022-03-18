/*
	Программирование на Golang. Строки. Задания

	Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var strBase string
	fmt.Scan(&strBase)

	for _, elem := range strBase {
		if strings.Count(strBase, string(elem)) == 1 {
			fmt.Print(string(elem))
		} else {
			continue
		}
	}

}
