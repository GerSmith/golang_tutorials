/*
	Программирование на Golang. Интерфейсы
*/

package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/" //тут играемся с параметрами
}

func printError(value interface{}) {
	fmt.Printf("value=%v: %T", value, value)
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	switch value1.(type) {
	case float64:
	default:
		printError(value1)
		return
	}

	switch value2.(type) {
	case float64:
	default:
		printError(value2)
		return
	}

	switch operation.(type) {
	case string:
		switch operation.(string) {
		case "+":
			fmt.Printf("%.4f", value1.(float64)+value2.(float64))
		case "-":
			fmt.Printf("%.4f", value1.(float64)-value2.(float64))
		case "*":
			fmt.Printf("%.4f", value1.(float64)*value2.(float64))
		case "/":
			fmt.Printf("%.4f", value1.(float64)/value2.(float64))
		}
	default:
		fmt.Println("неизвестная операция")
		return
	}

}
