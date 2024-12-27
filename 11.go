package main

import (
	"fmt"
	"os"
	"os/user"
	"bufio"
	"path/filepath"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Ошибка получения пользователя: ", err)
		return
	}
	
	// путь к истории 
	historyFile := filepath.Join(usr.HomeDir, ".bash_history")

	//открываем файл
	file, err := os.Open(historyFile)
	if err != nil {
		fmt.Println("Ошибка открытия файла истории: ", err)
		return
	}
	defer file.Close()

	var lastCommand string
	scanner :=bufio.NewScanner(file)
	for scanner.Scan(){
		lastCommand = scanner.Text()
	}
	if err != nil {
		fmt.Println("Ошибка чтения файла истории: ", err)
		return
	}
	
	fmt.Println(lastCommand)
}
