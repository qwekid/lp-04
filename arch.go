package main

import (
"flag"
"fmt"
"runtime"
)

func main() {
// Объявляем флаги
flag.Usage = func() {
fmt.Println("Использование: arch [опции]")
fmt.Println("Опции:")
fmt.Println(" --help Показать справку")
}
showHelp := flag.Bool("help", false, "Показать справку")

// Разбираем флаги
flag.Parse()

// Если указан флаг --help, показываем справку и выходим
if *showHelp {
	flag.Usage()
	return
}

// Получаем архитектуру системы
arch := runtime.GOARCH

fmt.Println(arch)
}

