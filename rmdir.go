package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ignoreFailOnNonEmpty bool
	printProgress         bool
	verbose               bool
	help                  bool
)

func init() {
	flag.BoolVar(&ignoreFailOnNonEmpty, "ignore-fail-on-non-empty", false, "Игнорировать ошибки, когда каталог не пуст")
	flag.BoolVar(&ignoreFailOnNonEmpty, "p", false, "Игнорировать ошибки, когда каталог не пуст (псевдоним для --ignore-fail-on-non-empty)")
	flag.BoolVar(&verbose, "v", false, "Включить подробный вывод")
	flag.BoolVar(&help, "h", false, "Показать справку")
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if flag.NArg() == 0 {
		fmt.Println("Ошибка: Не указан каталог.\n")
		flag.Usage()
		return
	}

	for _, dir := range flag.Args() {
		if err := removeDir(dir); err != nil {
			fmt.Printf("Ошибка при удалении каталога %s: %v\n", dir, err)
		}
	}
}

func removeDir(dir string) error {
	if verbose {
		fmt.Printf("Удаление каталога: %s\n", dir)
	}

	// Получаем список файлов и каталогов внутри указанного каталога
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Рекурсивно удаляем содержимое каталога
	for _, item := range items {
		itemPath := filepath.Join(dir, item.Name())
		if item.IsDir() {
			if err := removeDir(itemPath); err != nil {
				if !ignoreFailOnNonEmpty {
					return err
				}
				if verbose {
					fmt.Printf("Пропуск не пустого каталога: %s\n", itemPath)
				}
			}
		} else {
			if verbose {
				fmt.Printf("Удаление файла: %s\n", itemPath)
			}
			if err := os.Remove(itemPath); err != nil {
				return err
			}
		}
	}

	// Наконец, удаляем сам каталог
	if verbose {
		fmt.Printf("Удаление каталога: %s\n", dir)
	}
	return os.Remove(dir)
}
