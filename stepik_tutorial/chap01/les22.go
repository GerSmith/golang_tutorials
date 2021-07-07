/*
	Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
	Гарантируется, что цифры в числах не повторяются
*/

package main

import "fmt"

func main() {
	var A, B string

	fmt.Scan(&A, &B)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				fmt.Println(string(A[i]))
			}
		}
	}
}
