package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Определяем флаги
	homeFlag := flag.Bool("home", false, "Перейти в домашнюю директорию")
	backFlag := flag.Bool("back", false, "Перейти в предыдущую директорию")
	rootFlag := flag.Bool("root", false, "Перейти в корневую директорию")
	flag.Parse()

	var targetDir string

	// Обработка флагов
	if *homeFlag {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Ошибка: невозможно определить домашнюю директорию")
			os.Exit(1)
		}
		targetDir = homeDir
	} else if *backFlag {
		prevDir := os.Getenv("OLDPWD")
		if prevDir == "" {
			fmt.Println("Ошибка: предыдущая директория не установлена")
			os.Exit(1)
		}
		targetDir = prevDir
	} else if *rootFlag {
		targetDir = "/"
	} else {
		// Если флаги не переданы, берём первый аргумент (путь)
		if len(flag.Args()) > 0 {
			targetDir = flag.Args()[0]
		} else {
			fmt.Println("Ошибка: не указан путь")
			os.Exit(1)
		}
	}

	// Проверяем существование директории
	absPath, err := filepath.Abs(targetDir)
	if err != nil {
		fmt.Printf("Ошибка: невозможность преобразовать путь '%s': %v\n", targetDir, err)
		os.Exit(1)
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) || !info.IsDir() {
		fmt.Printf("Ошибка: директория '%s' недоступна или не существует\n", absPath)
		os.Exit(1)
	}

	// Выводим новую директорию для оболочки
	fmt.Println(absPath)

	// Устанавливаем переменные окружения для текущего процесса
	oldPwd, _ := os.Getwd()
	os.Setenv("OLDPWD", oldPwd)
	os.Setenv("PWD", absPath)

	// Меняем текущую директорию
	err = os.Chdir(absPath)
	if err != nil {
		fmt.Printf("Ошибка: невозможность сменить директорию на '%s': %v\n", absPath, err)
		os.Exit(1)
	}
}