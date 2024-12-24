package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определяем флаги
	logical := flag.Bool("L", false, "Выводит логический путь")
	physical := flag.Bool("P", false, "Выводит физический путь")

	// Парсим флаги
	flag.Parse()

	// Получаем текущую директорию
	var currentDir string
	var err error

	if *physical {
		// Получаем физический путь
		currentDir, err = os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения физического пути: %v", err)
			os.Exit(1)
		}
	} else {
		if *logical {
			// Получаем логический путь
			currentDir, err = os.Readlink("/proc/self/cwd")
			if err != nil {
				// Если нет доступа к /proc/self/cwd, используем os.Getwd()
				currentDir, err = os.Getwd()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Ошибка получения логического пути: %v", err)
					os.Exit(1)
				}
			}
		} else{
			// Получаем логический путь
			currentDir, err = os.Readlink("/proc/self/cwd")
			if err != nil {
				// Если нет доступа к /proc/self/cwd, используем os.Getwd()
				currentDir, err = os.Getwd()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Ошибка получения логического пути: %v", err)
					os.Exit(1)
				}
			}
		}
	}


	fmt.Println(currentDir)
}
