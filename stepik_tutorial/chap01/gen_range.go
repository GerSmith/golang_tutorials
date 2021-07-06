package main

import "fmt"

func gen(start, end, step int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i < end; i += step {
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {
	for i := range gen(8, 20, 2) {
		fmt.Printf("%d ", i) // 8 10 12 14 16 18
	}
}
