/*
   Укорот байтовой строки, версия 2
*/

package main

import (
	"fmt"
)

func main() {
	var text, res string
	var width int
	fmt.Scan(&text, &width)

	if len(text) > width {
		res = text[:width] + "..."
	}

	fmt.Println(res)
}
