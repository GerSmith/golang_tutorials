/*
	Программирование на Golang. Анонимные функции
*/

package main

import (
	"fmt"
	"strconv"
)

func deleteOdd(in uint) uint {
	var sum []uint
	var d uint

	for ; in > 0; in /= 10 {
		d = in % 10
		if d%2 != 0 || d == 0 {
			continue
		} else {
			sum = append(sum, d)
		}
	}

	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}

	var sum_str string
	for _, elem := range sum {
		sum_str += strconv.Itoa(int(elem))
	}

	res, _ := strconv.ParseUint(sum_str, 10, 64)

	if res == 0 {
		return uint(100)
	} else {
		return uint(res)
	}
}

func main() {
	var in uint = 727178
	var out uint = 28
	fmt.Println(deleteOdd(in))
	fmt.Println(deleteOdd(in) == out)
}
