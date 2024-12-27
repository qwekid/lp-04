package main

import (
"bufio"
"fmt"
"os"
"os/user"
"path/filepath"
"strconv"
"os/exec"
)

func main() {
// Проверяем, передан ли номер команды
if len(os.Args) < 2 {
fmt.Println("Использование: go run main.go <номер команды>")
return
}

// Получаем номер команды из аргументов
commandNumber, err := strconv.Atoi(os.Args[1])
if err != nil {
	fmt.Println("Ошибка: номер команды должен быть числом.")
	return
}

// Определяем текущего пользователя
usr, err := user.Current()
if err != nil {
	fmt.Println("Ошибка получения текущего пользователя:", err)
	return
}

// Путь к файлу истории Bash
historyFile := filepath.Join(usr.HomeDir, ".bash_history")

// Открываем файл истории
file, err := os.Open(historyFile)
if err != nil {
	fmt.Println("Ошибка открытия файла истории:", err)
	return
}
defer file.Close()

// Считываем файл построчно
scanner := bufio.NewScanner(file)
var history []string
for scanner.Scan() {
	history = append(history, scanner.Text())
}

if err := scanner.Err(); err != nil {
	fmt.Println("Ошибка чтения файла истории:", err)
	return
}

// Проверяем, существует ли команда с указанным номером
if commandNumber < 1 || commandNumber > len(history) {
	fmt.Printf("Команда с номером %d не найдена в истории.\n", commandNumber)
	return
}

// Получаем команду
targetCommand := history[commandNumber-1]
fmt.Printf("Команда с номером %d: %s\n", commandNumber, targetCommand)

execCmd := exec.Command("bash", "-c", targetCommand)
execCmd.Stdout = os.Stdout
execCmd.Stderr = os.Stderr
err = execCmd.Run()

if err != nil {
	fmt.Println("Ошибка выполнения команды:", err)
	return
}
}