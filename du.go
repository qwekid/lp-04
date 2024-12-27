package main

import (
"flag"
"fmt"
"os"
"path/filepath"
)

// Структура для хранения информации о размере
type DiskUsage struct {
Path string
Size int64
IsDir bool
}

// Константы для единиц измерения
const (
KB = 1024
MB = KB * 1024
GB = MB * 1024
)

// Функция для конвертации размера в читаемый формат
func humanReadableSize(size int64) string {
switch {
case size >= GB:
return fmt.Sprintf("%.1fG", float64(size)/float64(GB))
case size >= MB:
return fmt.Sprintf("%.1fM", float64(size)/float64(MB))
case size >= KB:
return fmt.Sprintf("%.1fK", float64(size)/float64(KB))
default:
return fmt.Sprintf("%dB", size)
}
}

// Функция для подсчета размера каталога или файла
func getDiskUsage(path string) (int64, error) {
var totalSize int64
err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
if err != nil {
return err
}

	// Если это директория, то не учитываем её размер в подсчете
	if !info.IsDir() {
		totalSize += info.Size()
	}
	return nil
})

if err != nil {
	return 0, err
}
return totalSize, nil
}

// Основная логика программы
func main() {
// Флаги
showHelp := flag.Bool("help", false, "Показать справку")
summarize := flag.Bool("s", false, "Показывать только общий размер для каждого аргумента")
showAll := flag.Bool("a", false, "Показывать размер всех файлов, а не только директорий")

// Разбор флагов
flag.Parse()

// Показать справку
if *showHelp {
	printHelp()
	return
}

// Если флаги не указаны, обрабатываем текущую директорию
paths := flag.Args()
if len(paths) == 0 {
	paths = append(paths, ".")
}

// Обрабатываем каждый путь
for _, path := range paths {
	// Получаем размер для каждого каталога
	size, err := getDiskUsage(path)
	if err != nil {
		fmt.Printf("Ошибка при подсчете размера для %s: %v\n", path, err)
		continue
	}

	// Если флаг -s, выводим только общий размер
	if *summarize {
		fmt.Printf("%s: %s\n", path, humanReadableSize(size))
		continue
	}

	// Если флаг -a, показываем размер всех файлов и директорий
	if *showAll {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Для файлов показываем их размер
			if !info.IsDir() {
				fmt.Printf("%s: %s\n", path, humanReadableSize(info.Size()))
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Ошибка при обходе каталога %s: %v\n", path, err)
		}
	}

	// Показываем общий размер
	if !*showAll {
		fmt.Printf("%s: %s\n", path, humanReadableSize(size))
	}
}
}

// Функция для вывода справки
func printHelp() {
fmt.Println("Использование: du [опции] [путь]")
fmt.Println("Опции:")
fmt.Println(" -s Показывать только общий размер для каждого аргумента")
fmt.Println(" -a Показывать размер всех файлов, а не только директорий")
fmt.Println(" -help Показать справку")
}