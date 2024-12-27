package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Функция для копирования файла
func copyFile(src, dest string) error {
	// Открываем исходный файл
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("не удалось открыть исходный файл: %v", err)
	}
	defer sourceFile.Close()

	// Создаем целевой файл
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("не удалось создать целевой файл: %v", err)
	}
	defer destFile.Close()

	// Копируем содержимое исходного файла в целевой файл
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("ошибка при копировании: %v", err)
	}

	// Устанавливаем режим доступа для целевого файла (оставляем как у исходного)
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}
	err = os.Chmod(dest, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("не удалось установить права доступа для целевого файла: %v", err)
	}

	return nil
}

// Функция для копирования директории рекурсивно
func copyDir(src, dest string) error {
	// Открываем директорию
	srcDir, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("не удалось открыть исходную директорию: %v", err)
	}
	defer srcDir.Close()

	// Получаем информацию о директории
	srcInfo, err := srcDir.Stat()
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о директории: %v", err)
	}

	// Создаем целевую директорию
	err = os.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("не удалось создать целевую директорию: %v", err)
	}

	// Получаем список файлов и директорий
	entries, err := srcDir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("не удалось прочитать содержимое директории: %v", err)
	}

	// Рекурсивно копируем каждый файл/подкаталог
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// Если это директория, вызываем функцию для копирования директории
			err := copyDir(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// Если это файл, вызываем функцию для копирования файла
			err := copyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Функция для запроса на перезапись файла
func promptOverwrite(dest string) bool {
	var answer string
	fmt.Printf("Файл %s уже существует. Перезаписать? (y/n): ", dest)
	_, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return false
	}

	answer = string(answer[0]) // Берем только первый символ
	return answer == "y" || answer == "Y"
}

func main() {
	// Флаги командной строки
	recursive := flag.Bool("r", false, "Рекурсивное копирование директорий (аналогично 'cp -r')")
	interactive := flag.Bool("i", false, "Запрос на перезапись файла (аналогично 'cp -i')")
	help := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы командной строки
	flag.Parse()

	// Если пользователь запросил справку
	if *help {
		fmt.Println("Программа cp для копирования файлов и директорий.")
		fmt.Println("Ключи:")
		fmt.Println("  -r             Рекурсивное копирование директорий.")
		fmt.Println("  -i             Запрос на перезапись файла, если файл существует.")
		fmt.Println("  --help         Показать справку.")
		return
	}

	// Проверяем, что есть два аргумента (исходный и целевой путь)
	if len(flag.Args()) < 2 {
		fmt.Println("Ошибка: необходимо указать исходный и целевой путь.")
		fmt.Println("Используйте --help для получения справки.")
		return
	}

	src := flag.Args()[0]  // Исходный файл/директория
	dest := flag.Args()[1] // Целевой файл/директория

	// Получаем информацию о исходном файле/директории
	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	// Проверяем, существует ли целевой файл/директория
	destInfo, err := os.Stat(dest)
	if err == nil {
		// Если целевой путь существует
		if destInfo.IsDir() && !srcInfo.IsDir() {
			fmt.Println("Ошибка: нельзя копировать файл в директорию.")
			return
		}
		if *interactive && !promptOverwrite(dest) {
			fmt.Println("Копирование отменено.")
			return
		}
	}

	// Если исходный путь это файл
	if !srcInfo.IsDir() {
		err := copyFile(src, dest)
		if err != nil {
			fmt.Printf("Ошибка при копировании файла: %v\n", err)
			return
		}
	} else {
		// Если исходный путь это директория
		if !*recursive {
			fmt.Println("Ошибка: для копирования директорий используйте ключ -r.")
			return
		}
		err := copyDir(src, dest)
		if err != nil {
			fmt.Printf("Ошибка при копировании директории: %v\n", err)
			return
		}
	}

	fmt.Println("Копирование завершено успешно.")
}
