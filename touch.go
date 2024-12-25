package main

import (
	"fmt"
	"os"
	"time"
)

func printHelp() {
	fmt.Println("Использование: touch [КЛЮЧИ] ФАЙЛ...")
	fmt.Println("Ключи:")
	fmt.Println("  -a          Изменить время последнего доступа.")
	fmt.Println("  -m          Изменить время последней модификации.")
	fmt.Println("  -c          Не создавать файл, если он не существует.")
	fmt.Println("  --help      Показать это сообщение.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: Не указаны файлы.")
		printHelp()
		return
	}

	var accessTime, modifyTime bool
	cOption := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-a":
			accessTime = true
		case "-m":
			modifyTime = true
		case "-c":
			cOption = true
		case "--help":
			printHelp()
			return
		default:
			fileName := arg
			_, err := os.Stat(fileName)
			if os.IsNotExist(err) && cOption {
				// Если файл не существует и указан ключ -c, просто пропускаем
				continue
			}

			if err != nil && !os.IsNotExist(err) {
				fmt.Printf("Ошибка: Не удалось получить информацию о файле %s: %v", fileName, err)
				return
			}

			if err != nil {
				// Файл не существует, создаем его
				file, err := os.Create(fileName)
				if err != nil {
					fmt.Printf("Ошибка: Не удалось создать файл %s: %v", fileName, err)
					return
				}
				file.Close()
			}

			now := time.Now()
			if accessTime {
				err := os.Chtimes(fileName, now, now)
				if err != nil {
					fmt.Printf("Ошибка: Не удалось изменить время доступа файла %s: %v", fileName, err)
					return
				}
			}
			if modifyTime {
				err := os.Chtimes(fileName, now, now)
				if err != nil {
					fmt.Printf("Ошибка: Не удалось изменить время модификации файла %s: %v", fileName, err)
					return
				}
			}
		}
	}
}
