package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Обработка ключа -help
	help := flag.Bool("help", false, "Показать помощь и выйти")
	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	// Получение аргументов командной строки (без имени программы)
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Ошибка: требуется один аргумент для завершения с указанным кодом.")
		fmt.Println("Используйте -help для получения информации.")
		os.Exit(1)
	}

	// Попытка преобразования аргумента в код завершения
	exitCode, err := parseExitCode(args[0])
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Завершение программы с кодом %d\n", exitCode)
	shellPid :=os.Getppid() // получение PID терминала
	syscall.Kill(shellPid, syscall.SIGHUP) 
}

func printHelp() {
	fmt.Println("Использование: exit [код завершения]")
	fmt.Println("Ключи:")
	fmt.Println("  -help       Показать это сообщение и выйти")
	fmt.Println("Аргументы:")
	fmt.Println("  [код завершения] Укажите целое число в диапазоне 0-255 для завершения программы с этим кодом.")
}

func parseExitCode(arg string) (int, error) {
	var code int
	_, err := fmt.Sscanf(arg, "%d", &code)
	if err != nil {
		return 0, fmt.Errorf("некорректный код завершения: %s", arg)
	}

	if code < 0 || code > 255 {
		return 0, fmt.Errorf("код завершения должен быть в диапазоне 0-255, получено: %d", code)
	}

	return code, nil
}
