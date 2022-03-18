package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := bufio.NewReader(bytes.NewBufferString("597"))
		var s int
		fmt.Fscan(r, &s)
		intV(s)
	}
}

func BenchmarkStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := bufio.NewReader(bytes.NewBufferString("597"))
		var s string
		fmt.Fscan(r, &s)
		strV(s)
	}
}

func strV(str string) {
	// var str string = "567"
	a := str[2]
	b := str[1]
	c := str[0]
	a, b, c = c, a, b
}

func intV(n int) {
	// var n int = 567
	a := n % 10
	b := (n % 100 / 10)
	c := (n / 100)
	a, b, c = c, a, b
}

func main() {
	/*
		start := time.Now()
		for i := 0; i < 9990000000; i++ {
			strV("567")
		}
		duration := time.Since(start)
		fmt.Println("string ", duration.Nanoseconds())

		start = time.Now()
		for i := 0; i < 9990000000; i++ {
			intV(567)
		}
		duration = time.Since(start)
		fmt.Println("integer", duration.Nanoseconds())
	*/
}
