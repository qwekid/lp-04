package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"
)

func main() {
	// Определяем ключи
	bFlag := flag.Bool("b", false, "Выводит тип файла в виде MIME-медиа-типов.\n")
	iFlag := flag.Bool("i", false, "Выводит дополнительную информацию о файле.\n")
	hFlag := flag.Bool("h", false, "Выводит справку по использованию.\n")

	// Парсим флаги
	flag.Parse()

	// Проверяем, если запрашивается помощь
	if *hFlag {
		printHelp()
		return
	}

	// Получаем аргументы командной строки (файлы для анализа)
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Ошибка: Не указаны файлы для анализа.\n")
		printHelp()
		return
	}

	// Обрабатываем каждый файл
	for _, arg := range args {
		if err := analyzeFile(arg, *bFlag, *iFlag); err != nil {
			fmt.Printf("Ошибка: %v", err)
		}
	}
}

func analyzeFile(filename string, bFlag bool, iFlag bool) error {
	// Читаем содержимое файла
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл '%s': %v", filename, err)
	}

	// Определяем тип файла (по содержимому)
	contentType := "неизвестный тип"
	if bFlag {
		// Определяем MIME-тип на основе расширения файла
		ext := filepath.Ext(filename)
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			contentType = mimeType
		} else {
			contentType = "application/octet-stream" // Если MIME-тип не найден
		}
	}

	if iFlag {
		fmt.Printf("Файл: %s, Размер: %d байт", filename, len(data))
	}

	if bFlag {
		fmt.Printf("Тип файла '%s': %s", filename, contentType)
	} else {
		fmt.Printf("Файл '%s' проанализирован.\n", filename)
	}
	return nil
}

func printHelp() {
	fmt.Println("Использование: file [опции] [файлы]")
	fmt.Println("Опции:")
	fmt.Println("  -b   Выводит тип файла в виде MIME-медиа-типов.")
	fmt.Println("  -i   Выводит дополнительную информацию о файле.")
	fmt.Println("  -h   Выводит справку по использованию.")
	fmt.Println("Пример:")
	fmt.Println("  file -b file.txt")
}
