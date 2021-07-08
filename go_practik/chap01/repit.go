/*
	Программа принимает на вход строку source и число times. Требуется склеить source саму с собой times раз и вернуть результат:
source = x, times = 3 → xxx
source = omm, times = 2 → ommomm
*/

package main

import "fmt"

func main() {
	var (
		source, result string
		times          int
	)

	fmt.Scan(&source, &times)

	for i := 0; i < times; i++ {
		result += source
	}

	fmt.Println(result)
}
