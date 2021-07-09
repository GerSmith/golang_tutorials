/*
	Напишите программу, которая определяет название языка по его коду. Правила:
		en → English
		fr → French
		ru или rus → Russian
		иначе → Unknown
*/

package main

import "fmt"

func main() {
	var code, lang string
	fmt.Scan(&code)

	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
