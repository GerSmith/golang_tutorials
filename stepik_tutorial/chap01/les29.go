/*
На ввод подаются пять чисел, которые записываются в массив.
Однако эта часть программы уже написана.
Вам нужно написать фрагмент кода, с помощью которого можно найти
и вывести максимальное число в этом массиве.
*/

package main

import "fmt"

func main() {
	arr := [5]int{}
	max := 0

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
		if max < arr[i] {
			max = arr[i]
		}
	}
	//fmt.Printf("%#v", arr)
	fmt.Println(max)
}