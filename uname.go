package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

// Функция для вывода информации о системе
func printUname(sysname, nodename, release, version, machine, os string) {
	fmt.Println("Системная информация:")
	fmt.Printf("Имя операционной системы: %s\n", sysname)
	fmt.Printf("Имя узла: %s\n", nodename)
	fmt.Printf("Релиз: %s\n", release)
	fmt.Printf("Версия: %s\n", version)
	fmt.Printf("Архитектура: %s\n", machine)
	fmt.Printf("Операционная система: %s\n", os)
}

func main() {
	// Определяем флаги командной строки
	all := flag.Bool("a", false, "Вывести всю информацию (аналогично 'uname -a')")
	kernel := flag.Bool("s", false, "Имя операционной системы (аналогично 'uname -s')")
	nodenameFlag := flag.Bool("n", false, "Имя узла (аналогично 'uname -n')")
	kernelVersion := flag.Bool("v", false, "Версия ядра (аналогично 'uname -v')")
	help := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы командной строки
	flag.Parse()

	// Если пользователь запросил справку
	if *help {
		fmt.Println("Программа uname для вывода информации о системе.")
		fmt.Println("Ключи:")
		fmt.Println("  -a             Вывести всю информацию о системе.")
		fmt.Println("  -s             Имя операционной системы.")
		fmt.Println("  -n             Имя узла.")
		fmt.Println("  -v             Версия ядра.")
		fmt.Println("  --help         Показать справку.")
		return
	}

	// Получаем информацию о системе
	sysname := runtime.GOOS      // имя операционной системы
	nodename, _ := os.Hostname() // имя узла
	release := runtime.Version() // версия ядра
	version := "неизвестно"      // версия ядра
	machine := runtime.GOARCH    // архитектура процессора
	os := runtime.GOOS           // операционная система

	// Если запросили всю информацию
	if *all {
		printUname(sysname, nodename, release, version, machine, os)
	} else {
		// Если запросили конкретную информацию
		if *kernel {
			fmt.Println(sysname)
		}
		if *nodenameFlag {
			fmt.Println(nodename)
		}
		if *kernelVersion {
			fmt.Println(release)
		}
		// Если не был указан ни один ключ, выводим все по умолчанию
		if !(*kernel || *nodenameFlag || *kernelVersion) {
			printUname(sysname, nodename, release, version, machine, os)
		}
	}
}
