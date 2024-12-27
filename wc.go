package main

import (
"flag"
"fmt"
"io"
"os"
"strings"
)

// Функция для подсчета количества строк, слов и байтов
func countFileStats(file *os.File) (lines, words, bytes int, err error) {
buf := make([]byte, 1024) // Буфер для чтения данных
var lineCount, wordCount, byteCount int

// Чтение данных из файла
for {
	n, err := file.Read(buf)
	if n == 0 && err == io.EOF {
		break
	}
	if err != nil && err != io.EOF {
		return 0, 0, 0, err
	}

	// Подсчет байтов
	byteCount += n

	// Подсчет строк и слов
	for i := 0; i < n; i++ {
		if buf[i] == '\n' {
			lineCount++
		}
		if buf[i] == ' ' || buf[i] == '\n' || buf[i] == '\t' {
			wordCount++
		}
	}
}

// Убираем пустые строки из подсчета слов
if lineCount > 0 && len(strings.TrimSpace(string(buf))) > 0 {
	wordCount++
}

return lineCount, wordCount, byteCount, nil
}

// Функция для вывода справки
func printHelp() {
fmt.Println("Использование: wc [опции] [файл]")
fmt.Println("Опции:")
fmt.Println(" -l Подсчитать количество строк")
fmt.Println(" -w Подсчитать количество слов")
fmt.Println(" -c Подсчитать количество байтов")
fmt.Println(" -help Показать справку")
}

func main() {
// Объявление флагов
lineFlag := flag.Bool("l", false, "Подсчитать количество строк")
wordFlag := flag.Bool("w", false, "Подсчитать количество слов")
byteFlag := flag.Bool("c", false, "Подсчитать количество байтов")
showHelp := flag.Bool("help", false, "Показать справку")

// Парсинг флагов
flag.Parse()

// Показать справку, если установлен флаг -help
if *showHelp {
	printHelp()
	return
}

// Проверка, передан ли файл
var file *os.File
var err error
if len(flag.Args()) > 0 {
	file, err = os.Open(flag.Args()[0])
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
} else {
	// Если файл не передан, читаем из стандартного ввода
	file = os.Stdin
}

// Подсчет статистики
lines, words, bytes, err := countFileStats(file)
if err != nil {
	fmt.Println("Ошибка при чтении файла:", err)
	return
}

// Вывод результатов в зависимости от флагов
if *lineFlag {
	fmt.Printf("Количество строк: %d\n", lines)
}
if *wordFlag {
	fmt.Printf("Количество слов: %d\n", words)
}
if *byteFlag {
	fmt.Printf("Количество байтов: %d\n", bytes)
}

// Если ни один флаг не установлен, выводим все данные
if !*lineFlag && !*wordFlag && !*byteFlag {
	fmt.Printf("Количество строк: %d\n", lines)
	fmt.Printf("Количество слов: %d\n", words)
	fmt.Printf("Количество байтов: %d\n", bytes)
}
}