/*
	Программирование на Golang. Работа с файлами
*/

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var sum int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		sInt, _ := strconv.Atoi(s)
		sum += sInt
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
