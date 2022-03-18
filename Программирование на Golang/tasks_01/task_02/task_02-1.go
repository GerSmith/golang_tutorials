package main

import (
	"fmt"
	"time"
)

func strV() {
	var str string = "567"
	a := str[2]
	b := str[1]
	c := str[0]
	a, b, c = c, a, b
}

func intV() {
	var n int = 567
	a := n % 10
	b := (n % 100 / 10)
	c := (n / 100)
	a, b, c = c, a, b
}

func main() {
	start := time.Now()
	for i := 0; i < 9990000000; i++ {
		intV()
	}
	duration := time.Since(start)
	fmt.Println("integer", duration.Nanoseconds())

	start = time.Now()
	for i := 0; i < 9990000000; i++ {
		strV()
	}
	duration = time.Since(start)
	fmt.Println("string ", duration.Nanoseconds())

}
