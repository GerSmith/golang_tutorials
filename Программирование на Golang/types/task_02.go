package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func result(s string) float64 {
	// убираем пробелы
	s = strings.ReplaceAll(s, " ", "")
	// меняем запятые на точки
	s = strings.ReplaceAll(s, ",", ".")
	// разбиваем на две части
	s_arr := strings.Split(s, ";")
	// приводим к вещественному формату
	one, _ := strconv.ParseFloat(s_arr[0], 64)
	two, _ := strconv.ParseFloat(s_arr[1], 64)
	result := one / two
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	result := result(s)
	fmt.Printf("%.4f", result)
}
