/*
	Программирование на Golang. Отображения (map)
*/

package main

import (
	"fmt"
	"time"
)

func work(x int) int {
	if x > 3 {
		time.Sleep(time.Millisecond * 500)
		return x + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return x - 1
	}
}

func main() {
	var n int
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)

		_, ok := m[n]

		if !ok {
			m[n] = work(n)
		}

		fmt.Print(m[n], " ")
	}
}
