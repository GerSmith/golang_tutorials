package main

/*
 Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
 Решение через строки
*/

import "fmt"

func main() {
	var a string
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Println("Последнее число")
	fmt.Println(string(a[len(a)-1]))
	fmt.Println("Предпоследнее число")
	fmt.Println(string(a[len(a)-2]))
}
