/*
Укорот байтовой строки
Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:

text = Eyjafjallajokull, width = 6 → Eyjafj...
Если строка не превышает указанной длины, укорачивать ее не следует:

text = hello, width = 6 → hello
Гарантируется, что в исходной строке text используются только однобайтовые символы без пробелов, а длина width строго больше 0.

Sample Input:
Eyjafjallajokull 6

Sample Output:
Eyjafj...
*/

package main

import "fmt"

func main() {
	var text, res string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	// fmt.Printf("%#v, %#v", text, width)

	/*
		cut := []byte(text)
		fmt.Printf("slice of byte: %v\n", cut)
		fmt.Printf("cut slice of byte: %v\n", cut[:width])
		fmt.Printf("try string: %v\n", string(cut[:width]))
		fmt.Println("try to answer:")
		fmt.Printf("%v...\n", string(cut[:width]))
	*/

	if len(text) > width {
		cut := make([]string, width)
		for idx, val := range text {
			if idx < width {
				//fmt.Println(idx, val)
				cut[idx] = string(val)
			} else {
				break
			}
		}
		//fmt.Println(cut)

		for i := 0; i < 3; i++ {
			cut = append(cut, ".")
		}
		//fmt.Println(cut)

		for _, val := range cut {
			res = res + val
		}
	} else {
		res = text
	}
	fmt.Println(res)
}
