/*
	карты
*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key"] = 7
	m["other"] = 13

	fmt.Println("map:", m)
	// map: map[key:7 other:13]

	val := m["key"]
	fmt.Println("val:", val)
	// val: 7

	fmt.Println("len:", len(m))
	// len: 2

	delete(m, "other")
	fmt.Println("map:", m)
	// map: map[key:7]

	_, ok := m["other"]
	fmt.Println("has other:", ok)
	// has other: false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	// map: map[bar:2 foo:1]
}
