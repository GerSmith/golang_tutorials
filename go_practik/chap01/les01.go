/*
	Напишите программу, которая считает количество минут во временном отрезке.
	1h30m = 90 min
	300s = 5 min
	10m = 10 min
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	/*fmt.Scan(&s)

	d, _ := time.ParseDuration(s)

	fmt.Println(s, " = ", d.Minutes(), "min")
	*/
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)
	fmt.Printf("%s = %v min\n", s, d.Minutes())
	fmt.Printf("s is %T, d is %T\n", s, d)
	fmt.Printf("%v\n", d.Hours())
	fmt.Printf("d.Hours() is %T, d.Minutes() is %T\n", d.Hours(), d.Minutes())
}
