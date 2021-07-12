/*
	🤔 Счетчик цифр
	Напишите программу, которая считает, сколько раз каждая цифра встречается в числе.
	Гарантируется, что на вход подаются только положительные целые числа,
	не выходящие за диапазон int.
*/

package main

import "fmt"

func main() {
	var number string
	fmt.Scanf("%s", &number)
	fmt.Println(number)

	num_sl := make([]string, len(number))
	for i := 0; i < len(number); i++ {
		num_sl = append(num_sl, number[i])
	}

	fmt.Println(num_sl)

	counter := make(map[int]int)
	fmt.Println(counter)

}
