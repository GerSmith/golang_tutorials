/*
	Программирование на Golang. Глава 1. Тесты.	
	Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
*/

package main

import "fmt"

func digitalRoot(number int) int {
	if number < 10 {
		return number
	} else {
		var sum int = 0

		for number > 0 {
			// fmt.Println(number % 10)
			sum = sum + number%10
			number = number / 10
		}

		if sum > 9 {
			return digitalRoot(sum)
		} else {
			return sum
		}
	}
}

func main() {
	var num int

	fmt.Scan(&num)
	//fmt.Println(num)
	fmt.Println(digitalRoot(num))

}
