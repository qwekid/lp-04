package main

import (
"fmt"
"os"
"strings"
)

func main() {
args := os.Args[1:]

if len(args) == 0 || args[0] == "help" {
	printHelp()
	return
}

exitCode := 0
exitMessage := ""

switch args[0] {
case "-f":
	exitCode = 1
	exitMessage = "Фатальная ошибка"
case "-s":
	exitCode = 0
	exitMessage = "Успешное завершение программы"
case "-r":
	exitCode = 2
	exitMessage = "Ошибка ввода данных"
default:
	fmt.Println("Неизвестный ключ. Используйте ключ help для справки.")
	return
}

fmt.Println(exitMessage)
os.Exit(exitCode)
}

func printHelp() {
helpText := `
Программа для завершения работы с заданным кодом.

Использование:
- ./exit -f   : завершение с фатальной ошибкой
- ./exit -s   : успешное завершение программы
- ./exit -r   : завершение с ошибкой ввода данных
- ./exit help : показать справку
`
fmt.Println(strings.TrimSpace(helpText))
}