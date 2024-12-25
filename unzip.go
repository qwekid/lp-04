package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Функция вывода справки
func printHelp() {
	fmt.Println("Использование: unzip [OPTIONS] ARCHIVE.zip")
	fmt.Println("Опции:")
	fmt.Println("  -h, --help     Показать это сообщение")
	fmt.Println("  -o, --output   Путь к директории для распаковки")
	fmt.Println("  -q, --quiet    Не выводить сообщения")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: недостаточно аргументов.")
		printHelp()
		return
	}

	var outputDir string
	var quietMode bool

	// Обрабатываем аргументы командной строки
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "-h", "--help":
			printHelp()
			return
		case "-o", "--output":
			if i+1 < len(os.Args) {
				outputDir = os.Args[i+1]
				i++
			} else {
				fmt.Println("Ошибка: отсутствует аргумент для опции -o/--output.")
				return
			}
		case "-q", "--quiet":
			quietMode = true
		default:
			if strings.HasSuffix(arg, ".zip") {
				// Предполагаем, что это файл архива
				if !quietMode {
					fmt.Printf("Распаковка архива: %s", arg)
				}
				cmd := exec.Command("unzip", arg)
				if outputDir != "" {
					cmd.Args = append(cmd.Args, "-d", outputDir)
				}
				if err := cmd.Run(); err != nil {
					fmt.Printf("Ошибка при распаковке: %s", err)
				}
				return
			} else {
				fmt.Printf("Ошибка: неопознанный аргумент: %s", arg)
				printHelp()
				return
			}
		}
	}

	if outputDir == "" {
		fmt.Println("Ошибка: не указан путь для распаковки. Используйте -o/--output.")
		printHelp()
	}
}
