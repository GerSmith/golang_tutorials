/*
	Программирование на Golang. Работа с файлами
*/


package main

import (
	"fmt"
	"os"
)

func main() {
	perm := os.FileMode(0600)
	fmt.Printf("%s | %d | %X\n", perm, perm, int(perm))
	// -rw------- | 384 | 180
}
