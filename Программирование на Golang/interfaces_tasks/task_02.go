/*
	Программирование на Golang. Интерфейсы
*/

package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"strings"
)

type Battery struct {
	Charge string
}

func (b *Battery) String() string {
	return fmt.Sprintf(b.Charge)
}

func main() {
	// batteryForTest - не забывайте об имени
	// } Скобка, закрывающая функцию main() вам не видна, но она есть
	var text string
	_, _ = fmt.Scan(&text)

	chrgCntUp := strings.Count(text, "1")
	chrgCntDown := strings.Count(text, "0")
	chrg := "[" + strings.Repeat(" ", chrgCntDown) + strings.Repeat("X", chrgCntUp) + "]"

	batteryForTest := &Battery{Charge: chrg}
	fmt.Println(batteryForTest)

}
