
package main

import (
"flag"
"fmt"
"syscall"
"strconv"
)

// Маппинг сигналов
var signalNames = map[string]syscall.Signal{
"TERM": syscall.SIGTERM,
"KILL": syscall.SIGKILL,
"INT": syscall.SIGINT,
"HUP": syscall.SIGHUP,
"QUIT": syscall.SIGQUIT,
"USR1": syscall.SIGUSR1,
"USR2": syscall.SIGUSR2,
"STOP": syscall.SIGSTOP,
"CONT": syscall.SIGCONT,
"SEGV": syscall.SIGSEGV,
"PIPE": syscall.SIGPIPE,
"ALRM": syscall.SIGALRM,
}

// Функция для печати доступных сигналов
func printSignalList() {
fmt.Println("Доступные сигналы:")
for name := range signalNames {
fmt.Println(name)
}
}

// Функция для отправки сигнала процессу
func sendSignal(pid int, sig syscall.Signal) error {
err := syscall.Kill(pid, sig)
if err != nil {
return fmt.Errorf("не удалось отправить сигнал процессу с PID %d: %v", pid, err)
}
return nil
}

// Функция для обработки справки
func printHelp() {
fmt.Println("Использование: kill [опции] [PID]")
fmt.Println("Опции:")
fmt.Println(" -s сигнал Отправить указанный сигнал (по умолчанию SIGTERM)")
fmt.Println(" -l Показать список всех доступных сигналов")
fmt.Println(" -help Показать справку")
}

func main() {
// Флаги командной строки
signalFlag := flag.String("s", "TERM", "Укажите сигнал для отправки (по умолчанию SIGTERM)")
listFlag := flag.Bool("l", false, "Показать список всех доступных сигналов")
helpFlag := flag.Bool("help", false, "Показать справку")

// Парсим флаги
flag.Parse()

// Показать справку
if *helpFlag {
	printHelp()
	return
}

// Показать список сигналов, если указан флаг -l
if *listFlag {
	printSignalList()
	return
}

// Проверка, был ли передан PID
if len(flag.Args()) < 1 {
	fmt.Println("Ошибка: не указан PID процесса.")
	printHelp()
	return
}

// Получаем PID из аргументов
pid, err := strconv.Atoi(flag.Args()[0])
if err != nil {
	fmt.Printf("Ошибка: некорректный PID '%s'.\n", flag.Args()[0])
	printHelp()
	return
}

// Определяем сигнал по имени
sig, found := signalNames[*signalFlag]
if !found {
	fmt.Printf("Ошибка: сигнал '%s' не поддерживается.\n", *signalFlag)
	printHelp()
	return
}

// Отправляем сигнал
err = sendSignal(pid, sig)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("Сигнал %s (PID: %d) успешно отправлен процессу %d.\n", *signalFlag, pid, pid)
}

