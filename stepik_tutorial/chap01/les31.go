/*

 */

package main

import "fmt"

func main() {
	var N, count int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
		if arr[i] > 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
