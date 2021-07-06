package main

import "fmt"

func split(num int) (left_num, right_num int) {
	/*
	   Функция разбиения шестизначного числа на левую и правую части (по три цифры).
	*/
	left_num = num / 1000
	right_num = num % 1000
	// тк возвращаемые значения именованы  - стр. 15 `(left_num, right_num int)`,
	// можно использовать пустой `return` (без аргументов),
	// что эквивалентно `return left_num, right_num`
	return
}

func sum_numbers(num int) (sum int) {
	/*
	   Функция вычисления суммы цифр числа.

	   - тк возвращаемое значение суммы именовано и является целочисленным,
	     то ему автоматически присваевается значение 0,
	     что позволяет опустить строчку инициализации `sum := 0`.
	*/
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	// тк возвращаемые значения именованы - стр. 17 `(sum int)`,
	// можно использовать пустой `return` (без аргументов),
	// что эквивалентно `return sum`
	return
}

func compare_nums(left_num, right_num int) {
	/*
	   Функция сравнения двух чисел на равенство.
	   Вывод функции "YES", если они равны, иначе "NO".
	*/
	if left_num == right_num {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var num int
	fmt.Scan(&num)

	left_num, right_num := split(num)
	left_sum, right_sum := sum_numbers(left_num), sum_numbers(right_num)
	compare_nums(left_sum, right_sum)
}
