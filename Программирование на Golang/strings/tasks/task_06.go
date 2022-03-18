/*
	Программирование на Golang. Строки. Задания

	Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isPwdValid(rsPwd []rune) bool {
	if len(rsPwd) >= 5 {
		for _, r := range rsPwd {
			if unicode.Is(unicode.Latin, r) || unicode.IsDigit(r) {
				continue
			} else {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func main() {
	// ввод данных с STDIN
	pwd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rsPwd := []rune(pwd)
	// обрезка
	rsPwd = rsPwd[:len(rsPwd)-2]
	// тестовый вывод
	//fmt.Println(string(rsPwd))
	//fmt.Println(len(rsPwd))
	if isPwdValid(rsPwd) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
