package main

import (
	"fmt"
	"os"
)

func main() {
	// Получаем текущий рабочий каталог
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Выводим текущий рабочий каталог
	fmt.Println(dir)
}
