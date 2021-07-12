/*

 */

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	// fmt.Printf("%v", arr)
	for idx, val := range arr {
		if idx%2 == 1 {
			fmt.Printf("%v ", val)
		}

	}
}
