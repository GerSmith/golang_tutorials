/*
срезы
*/

package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Printf("empty: %#v\n", s) //Шаблон %#v возвращает «внутреннее представление» значения
	// empty: []string{"", "", ""}

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	// set: [a b c]

	fmt.Println("get:", s[2])
	// get: c

	fmt.Println("len:", len(s))
	// len: 3

	fmt.Println("src:", s)
	// src: [a b c]

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("upd:", s)
	// upd: [a b c d e f]

	//Срез можно скопировать через встроенную функцию copy()
	dst := make([]string, len(s))
	copy(dst, s)
	fmt.Println("copy:", dst)
	// copy: [a b c d e f]

	// Выражение slice[from:to] вернет срез от элемента с индексом from включительно до элемента с индексом to не включительно
	sl1 := s[2:5]
	fmt.Println("sl1:", sl1)
	// sl1: [c d e]
	sl2 := s[:5] //Этот срез включает все элементы, кроме s[5]
	fmt.Println("sl2:", sl2)
	// sl2: [a b c d e]
	sl3 := s[2:] //А этот срез включает элементы от s[2] и до конца
	fmt.Println("sl3:", sl3)
	// sl3: [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("init:", t)
	// init: [g h i]
}
