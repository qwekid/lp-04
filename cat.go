package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определение флагов
	numberLines := flag.Bool("n", false, "Нумеровать все строки")
	numberNonEmptyLines := flag.Bool("b", false, "Нумеровать непустые строки")
	addDollar := flag.Bool("E", false, "Добавлять символ $ в конце каждой строки")

	flag.Parse()

	// Получение оставшихся аргументов (файлов)
	files := flag.Args()
	if len(files) == 0 {
		// Если не указаны файлы, читаем из stdin
		process(os.Stdin, *numberLines, *numberNonEmptyLines, *addDollar)
	} else {
		// Обрабатываем каждый указанный файл
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка открытия файла %s: %v", file, err)
				continue
			}
			defer f.Close()
			process(f, *numberLines, *numberNonEmptyLines, *addDollar)
		}
	}
}

func process(r *os.File, numberLines bool, numberNonEmptyLines bool, addDollar bool) {
	scanner := bufio.NewScanner(r)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()

		// Добавляем символ $ в конце строки, если установлен флаг -E
		if addDollar {
			line += "$"
		}

		// Нумеруем строки
		if numberLines {
			fmt.Printf("%d\t%s", lineNumber, line)
			lineNumber++
		} else if numberNonEmptyLines {
			if strings.TrimSpace(line) != "" {
				fmt.Printf("%d\t%s", lineNumber, line)
			} else {
				fmt.Println(line)
			}
			if strings.TrimSpace(line) != "" {
				lineNumber++
			}
		} else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения: %v", err)
	}
}
