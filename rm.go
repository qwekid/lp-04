package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Обработка ключей
	help := flag.Bool("help", false, "Показать помощь и выйти")
	recursive := flag.Bool("R", false, "Рекурсивное удаление каталогов")
	force := flag.Bool("f", false, "Игнорировать ошибки")
	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	// Получение аргументов командной строки (без имени программы)
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Ошибка: требуется указать хотя бы один файл или каталог для удаления.")
		fmt.Println("Используйте -help для получения информации.")
		os.Exit(1)
	}

	// Удаление файлов и каталогов
	for _, target := range args {
		err := removeTarget(target, *recursive, *force)
		if err != nil {
			if !*force {
				fmt.Printf("Ошибка при удалении '%s': %s\n", target, err)
				os.Exit(1)
			}
		}
	}
}

func printHelp() {
	fmt.Println("Использование: rm [ключи] [файлы/каталоги]")
	fmt.Println("Ключи:")
	fmt.Println("  -help       Показать это сообщение и выйти")
	fmt.Println("  -R          Рекурсивное удаление каталогов")
	fmt.Println("  -f          Игнорировать ошибки при удалении")
}

func removeTarget(target string, recursive, force bool) error {
	info, err := os.Stat(target)
	if err != nil {
		if force {
			return nil
		}
		return fmt.Errorf("не удалось получить информацию о файле: %s", err)
	}

	if info.IsDir() {
		if !recursive {
			return errors.New("'" + target + "' является каталогом. Используйте -R для рекурсивного удаления.")
		}
		return os.RemoveAll(target)
	}

	return os.Remove(target)
}
