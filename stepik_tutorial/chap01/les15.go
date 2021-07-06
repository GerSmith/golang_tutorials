package main

import "fmt"

/*
	Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
*/

func main() {
	var A, B, sum int

	fmt.Scan(&A, &B)

	// подсчёт арифметической прогрессии
	fmt.Println((B - A + 1) * (A + B) / 2)

	for A <= B {
		sum += A
		A++
	}

	fmt.Println(sum)
}