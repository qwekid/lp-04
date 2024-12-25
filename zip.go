package main

import (
	"fmt"
	"os"
	"strings"
)

func displayHelp() {
	fmt.Println("Утилита zip - создание zip-архивов.")
	fmt.Println("Использование: zip [ключи] [архив] [файлы...]")
	fmt.Println("Ключи:")
	fmt.Println("  -h, --help         показать это сообщение")
	fmt.Println("  -r                 рекурсивно добавлять файлы из подкаталогов")
	fmt.Println("  -q                 тихий режим (меньше сообщений)")
	fmt.Println("  -v                 показать подробную информацию о процессе")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: недостаточно аргументов.")
		displayHelp()
		return
	}

	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			displayHelp()
			return
		case "-r":
			fmt.Println("Ключ -r: рекурсивное добавление файлов из подкаталогов.")
		case "-q":
			fmt.Println("Ключ -q: тихий режим активирован.")
		case "-v":
			fmt.Println("Ключ -v: показывается подробная информация о процессе.")
		default:
			if strings.HasSuffix(arg, ".zip") {
				fmt.Printf("Создание архива: %s", arg)
			} else {
				fmt.Printf("Добавление файла: %s", arg)
			}
		}
	}
}
