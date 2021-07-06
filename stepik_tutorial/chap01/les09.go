package main

import "fmt"

func main() {
	var (
		num string
		n int
	)
	
	fmt.Scan(&num)
	fmt.Println(string(num[0]))

	fmt.Scan(&n)
	switch {
	case n < 10:
		fmt.Println(n)
	case n < 100:
		fmt.Println(n / 10)
	case n < 1000:
		fmt.Println(n / 100)
	case n < 10000:
		fmt.Println(n / 1000)
	case n < 100000:
		fmt.Println(n / 10000)
	}

}
