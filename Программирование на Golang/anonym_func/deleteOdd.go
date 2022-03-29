package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in int = 727178
	// var out int = 28

	fmt.Println("Before: ", in)

	var sum []int
	var d int

	for ; in > 0; in /= 10 {
		d = in % 10
		fmt.Println("d = ", d)
		if d%2 != 0 || d == 0 {
			continue
		} else {
			sum = append(sum, d)
			fmt.Println("sum = ", sum)
		}
	}

	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}

	var sum_str string
	for _, elem := range sum {
		sum_str += strconv.Itoa(elem)
	}

	fmt.Println("After: ", sum_str)
	// fmt.Println("in ==  out: ", sum = out)
}
