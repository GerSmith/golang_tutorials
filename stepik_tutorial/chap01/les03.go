package main

import "fmt"

// iota идентификатор Go используется в объявлениях констант для упрощения определений увеличивающихся чисел

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6
}
