package main
import "fmt"

func main() {
	var n int
	for fmt.Scan(&n); n < 100 || n >= 1000; fmt.Scan(&n) {}
	fmt.Println( n / 100 + n / 10 % 10 + n % 10)
}
