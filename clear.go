package main

import (
"flag"
"fmt"
)

func main() {
// Объявление флагов
showHelp := flag.Bool("help", false, "Показать справку")
softClear := flag.Bool("s", false, "Мягкая очистка экрана (перенос строки без полной очистки)")
clearHistory := flag.Bool("h", false, "Очистить экран и историю терминала")
customMessage := flag.String("m", "", "Очистить экран и вывести сообщение")

// Разбор флагов
flag.Parse()

// Если указан флаг --help, показываем справку и выходим
if *showHelp {
	printHelp()
	return
}

// Логика работы флагов
if *softClear {
	softClearScreen()
	return
}

if *clearHistory {
	clearTerminalHistory()
	return
}

if *customMessage != "" {
	clearScreenWithMessage(*customMessage)
	return
}

// Если флаги не указаны, выполняем стандартное очищение экрана
clearScreen()
}

// Функция для вывода справки
func printHelp() {
fmt.Println("Использование: clear [опции]")
fmt.Println("Опции:")
fmt.Println(" -help Показать справку")
fmt.Println(" -s Мягкая очистка экрана (перенос строки без полной очистки)")
fmt.Println(" -h Очистить экран и историю терминала (не во всех терминалах)")
fmt.Println(" -m <сообщение> Очистить экран и вывести сообщение")
}

// Стандартная очистка экрана
func clearScreen() {
fmt.Print("\033[H\033[2J") // ANSI escape code для очистки экрана
}

// Мягкая очистка экрана (без удаления содержимого)
func softClearScreen() {
fmt.Println("\n\n") // Просто добавляем несколько пустых строк
}

// Очистка экрана и истории терминала
func clearTerminalHistory() {
// ANSI escape code для очистки экрана и истории терминала
fmt.Print("\033c")
}

// Очистка экрана и вывод пользовательского сообщения
func clearScreenWithMessage(message string) {
clearScreen()
fmt.Println(message)
}

