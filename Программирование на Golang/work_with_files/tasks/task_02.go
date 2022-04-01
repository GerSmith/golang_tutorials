/*
	Программирование на Golang. Работа с файлами
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Открываем файл с данными, проверяем на ошибку открытия
	file, err := os.Open("./task_02.data")
	if err != nil {
		log.Fatal(err)
	}
	// Закрываем файл
	defer file.Close()

	// Создаём bufio.Reader
	reader := bufio.NewReader(file)

	var pos int = 1 // счётчик позиций
	// Перебираем данный из файла используя в качестве разделителя ';'
	for {
		data, err := reader.ReadString(';')
		if err != nil {
			// Если файл закончился
			if err == io.EOF {
				break
			} else {
				log.Println(err)
				return
			}
		}
		// Наши данные имеют на конце "хвостик", учитываем это при сравнениии
		if data == "0;" {
			fmt.Printf("%v is %T, pos is %d, err is %v\n", data, data, pos, err)
		}
		pos++
	}
}
