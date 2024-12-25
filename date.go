package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Обработка ключей
	help := flag.Bool("help", false, "Показать помощь и выйти")
	iso := flag.Bool("I", false, "Показать дату и время в ISO 8601 формате")
	onlyDate := flag.Bool("d", false, "Показать только дату")
	onlyTime := flag.Bool("t", false, "Показать только время")
	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	// Получение текущего времени
	now := time.Now()

	// Обработка ключей вывода
	if *iso {
		fmt.Println(now.Format(time.RFC3339))
		os.Exit(0)
	}

	if *onlyDate {
		fmt.Println(now.Format("2006-01-02"))
		os.Exit(0)
	}

	if *onlyTime {
		fmt.Println(now.Format("15:04:05"))
		os.Exit(0)
	}

	// Стандартный вывод даты и времени
	fmt.Println(now.Format("02.01.2006 15:04:05"))
}

func printHelp() {
	fmt.Println("Использование: date [ключи]")
	fmt.Println("Ключи:")
	fmt.Println("  -help       Показать это сообщение и выйти")
	fmt.Println("  -I          Показать дату и время в ISO 8601 формате")
	fmt.Println("  -d          Показать только дату")
	fmt.Println("  -t          Показать только время")
}
