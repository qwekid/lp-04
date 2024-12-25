package main

import (
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("Использование:")
	fmt.Println("  echo <сообщение>   - выводит сообщение обратно")
	fmt.Println("  count <строка>    - выводит количество символов в строке")
	fmt.Println("  reverse <строка>  - выводит строку в обратном порядке")
	fmt.Println("  help              - выводит это сообщение")
}

func echo(args []string) {
	if len(args) < 2 {
		fmt.Println("Ошибка: недостаточно аргументов для команды echo.")
		return
	}
	message := strings.Join(args[1:], " ")
	fmt.Println(message)
}

func count(args []string) {
	if len(args) < 2 {
		fmt.Println("Ошибка: недостаточно аргументов для команды count.")
		return
	}
	str := strings.Join(args[1:], " ")
	fmt.Printf("Количество символов: %d", len(str))
}

func reverse(args []string) {
	if len(args) < 2 {
		fmt.Println("Ошибка: недостаточно аргументов для команды reverse.")
		return
	}
	str := strings.Join(args[1:], " ")
	reversed := reverseString(str)
	fmt.Println(reversed)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "-echo":
		echo(os.Args)
	case "-count":
		count(os.Args)
	case "-reverse":
		reverse(os.Args)
	case "-help":
		printHelp()
	default:
		fmt.Println("Ошибка: неизвестная команда. Используйте 'help' для получения справки.")
	}
}
