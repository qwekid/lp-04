package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func ls() {
	// Определение флагов
	rec := flag.Bool("R", false, "рекурсивный вывод")
	reverse := flag.Bool("r", false, "обратный порядок")
	all := flag.Bool("a", false, "включить скрытые файлы")
	long := flag.Bool("l", false, "длинный формат")
	human := flag.Bool("h", false, "читаемый размер")

	flag.Parse()

	// Получаем список аргументов (директорий)
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, ".") // Если аргументы не указаны, используем текущую директорию
	}

	for _, arg := range args {
		if err := listDirectory(arg, *rec, *reverse, *all, *long, *human); err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v \n", err)
		}
	}
}

func listDirectory(path string, recursive, reverse, all, long, human bool) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	if reverse {
		// Обратный порядок вывода
		for i := len(entries) - 1; i >= 0; i-- {
			printEntry(entries[i], path, long, human)

		}
	} else {
		for _, entry := range entries {
			if !all && entry.Name()[0] == '.' {
				continue
			}
			printEntry(entry, path, long, human)
		}
	}

	if recursive {
		for _, entry := range entries {
			if entry.IsDir() {
				newPath := filepath.Join(path, entry.Name())
				if err := listDirectory(newPath, recursive, reverse, all, long, human); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func printEntry(entry os.DirEntry, path string, long, human bool) {
	if long {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при получении информации о файле: %v \n", err)
			return
		}
		size := info.Size()
		if human {
			size = humanReadableSize(size)
		}
		modTime := info.ModTime().Format(time.RFC822)
		fmt.Printf("%s %d %s %s \n", entry.Name(), size, modTime, info.Mode())
	} else {
		fmt.Println(entry.Name())
	}
}

func humanReadableSize(size int64) int64 {
	// Функция для преобразования размера в читаемый формат (например, КБ, МБ)
	return size // Здесь можно добавить логику для форматирования размера
}
