package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
			fmt.Printf("Ошибка: %v\n", err)
		}
	}
}

func analyzeFile(filename string, bFlag bool, iFlag bool) error {
	// Читаем содержимое файла
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл '%s': %v\n", filename, err)
	}

	// Определяем тип файла (простой анализ)
	contentType := "неизвестный тип"
	if bFlag {
		contentType = "text/plain" // Пример, здесь можно добавить более сложную логику
	}

	if iFlag {
		fmt.Printf("Файл: %s, Размер: %d байт", filename, len(data))
	}

	fmt.Printf("Тип файла '%s': %s\n", filename, contentType)
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
