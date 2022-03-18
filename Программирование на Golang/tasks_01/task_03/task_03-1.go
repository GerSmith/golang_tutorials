package main

import (
	"fmt"
	"time"
)

func main() {
	var k int
	fmt.Scan(&k)
	d := time.Date(0, 0, 0, 0, 0, k, 0, time.UTC)

	fmt.Printf("It is %d hours %d minutes.", d.Hour(), d.Minute())
}
