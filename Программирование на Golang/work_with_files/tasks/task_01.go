/*
	Программирование на Golang. Работа с файлами
*/

package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"log"
)

func main() {
	// Открываем zip-файл для чтения
	z, err := zip.OpenReader("./task_01.zip")
	// Обработка ошибки
	if err != nil {
		log.Fatal(err)
	}
	// Закрываем zip-файл после выполнения всей программы
	defer z.Close()

	// Перебираем файлы в архиве
	for _, f := range z.File {
		r, err := f.Open()
		// Обработка ошибки открытия содержимого файла
		if err != nil {
			log.Fatal(err)
		}
		rows, _ := csv.NewReader(r).ReadAll()
		if len(rows) == 10 {
			fmt.Printf("Contents of %s:\n", f.Name)
			fmt.Println(rows[4][2])
		}
		// Закрываем файл
		r.Close()
	}
}
