package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Функция для поиска файлов
func findFiles(root string, name string, fileType string, maxDepth int) error {
	// Функция обхода директорий с ограничением глубины
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Если возникла ошибка при доступе к файлу
		if err != nil {
			return err
		}

		// Ограничиваем глубину поиска
		relativePath, _ := filepath.Rel(root, path)
		depth := strings.Count(relativePath, string(filepath.Separator))
		if depth > maxDepth {
			return filepath.SkipDir
		}

		// Фильтрация по имени файла
		if name != "" && !strings.Contains(info.Name(), name) {
			return nil
		}

		// Фильтрация по типу файла
		if fileType != "" && fileType != getFileType(info) {
			return nil
		}

		// Печатаем путь к файлу, если все условия выполнены
		fmt.Println(path)
		return nil
	})
}

// Функция для получения типа файла
func getFileType(info os.FileInfo) string {
	if info.IsDir() {
		return "d" // директория
	}
	return "f" // обычный файл
}

func main() {
	// Флаги командной строки
	root := flag.String("path", ".", "Путь к директории для поиска (по умолчанию текущая директория)")
	name := flag.String("name", "", "Поиск по имени файла")
	fileType := flag.String("type", "", "Тип файла (f — обычный файл, d — директория)")
	maxDepth := flag.Int("maxdepth", 0, "Ограничение глубины поиска")
	help := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы командной строки
	flag.Parse()

	// Если пользователь запросил справку
	if *help {
		fmt.Println("Программа find для поиска файлов по различным критериям.")
		fmt.Println("Ключи:")
		fmt.Println("  -path <путь>    Путь к директории для поиска (по умолчанию текущая).")
		fmt.Println("  -name <имя>     Поиск по имени файла.")
		fmt.Println("  -type <тип>     Поиск по типу файла (f — обычный файл, d — директория).")
		fmt.Println("  -maxdepth <глубина>  Ограничение глубины поиска.")
		fmt.Println("  --help          Показать справку.")
		return
	}

	// Проверяем путь
	if _, err := os.Stat(*root); os.IsNotExist(err) {
		fmt.Printf("Ошибка: указанная директория не существует: %s\n", *root)
		return
	}

	// Выполняем поиск файлов
	err := findFiles(*root, *name, *fileType, *maxDepth)
	if err != nil {
		fmt.Printf("Ошибка при поиске файлов: %v\n", err)
	}
}
