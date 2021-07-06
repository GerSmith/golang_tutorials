package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f, num int

	fmt.Scan(&num)

	a = num / 100000
	b = num / 10000 % 10
	c = num / 1000 % 10
	d = num / 100 % 10
	e = num / 10 % 10
	f = num % 10

	// fmt.Println(a, b, c, d, e, f)
	if a+b+c == d+e+f {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
