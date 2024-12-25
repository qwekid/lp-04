package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func printHelp() {
	fmt.Println("Использование: nl [файл] [опции]")
	fmt.Println("Опции:")
	fmt.Println("  -b, --body   Нумеровать только строки тела")
	fmt.Println("  -f, --footer Нумеровать только нижний колонтитул")
	fmt.Println("  -h, --help   Показать это сообщение")
}

func main() {
	body := flag.Bool("b", false, "Нумеровать только строки тела")
	footer := flag.Bool("f", false, "Нумеровать только нижний колонтитул")
	help := flag.Bool("h", false, "Показать справку")

	flag.Parse()

	if *help {
		printHelp()
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Ошибка: Не указан файл.")
		printHelp()
		return
	}

	fileName := args[0]

	// Чтение содержимого файла
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %s\n", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var result []string

	for i, line := range lines {
		if *body && i < len(lines)-1 {
			result = append(result, fmt.Sprintf("%d %s", i+1, line))
		} else if *footer && i == len(lines)-1 {
			result = append(result, fmt.Sprintf("%d %s", i+1, line))
		}
	}

	// Вывод результата
	for _, line := range result {
		fmt.Println(line)
	}
}
