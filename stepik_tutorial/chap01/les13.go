package main

import "fmt"

/*
	Требуется вывести слово YES, если год является високосным и NO - в противном случае.
*/

func main() {
	var year int

	fmt.Scan(&year)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
