/*
	Программирование на Golang. Преобразование типов данных
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt8)   // 127
	fmt.Println(math.MaxUint8)  // 255
	fmt.Println(math.MaxInt16)  // 32767
	fmt.Println(math.MaxUint16) // 65535
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	
}
