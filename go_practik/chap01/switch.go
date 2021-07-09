/*
switch и приёмы работы с ним
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		В отличие от многих других языков, не нужно указывать break.
		Go выполняет только подходящую ветку и не проваливается в следующую
	*/

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// Write 2 as two
	/*
		В одной ветке можно указать несколько выражений.
		Ветка default сработает, если остальные не подошли
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// It's the weekend
	/*
		Выражения в ветках не обязательно должны быть константами.
		switch может работать как if
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// It's before noon
}
