/*
	срезы
*/

package main

import "fmt"

func main() {
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"} // базовый массив

	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4]  // с 1-го по 4-й
	users3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users1) // [Kate Sam Tom Paul]
	fmt.Println(users2) // [Bob Alice Kate Sam]
	fmt.Println(users3) // [Sam Tom Paul Mike Robert]

	a := []int{1, 2, 3}
	a = append(a, 4, 5)

	fmt.Println(a) // [1 2 3 4 5]
}
