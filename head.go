package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func printUsage() {
	fmt.Println("Использование: head [Ключи] [Файл]")
	fmt.Println("Ключи:")
	fmt.Println("  -c N    Вывести первые N байтов файла")
	fmt.Println("  -b N    Вывести первые N байтов файла (аналогично -c)")
	fmt.Println("  -n N    Вывести первые N строк файла")
	fmt.Println("  --help  Показать это сообщение")
}

func main() {
	cFlag := flag.Int("c", -1, "Вывести первые N байтов файла")
	bFlag := flag.Int("b", -1, "Вывести первые N байтов файла (аналогично -c)")
	nFlag := flag.Int("n", -1, "Вывести первые N строк файла")
	helpFlag := flag.Bool("help", false, "Показать это сообщение")

	flag.Parse()

	if *helpFlag {
		printUsage()
		return
	}

	var nBytes, nLines int
	if *cFlag >= 0 {
		nBytes = *cFlag
	} else if *bFlag >= 0 {
		nBytes = *bFlag
	} else {
		nLines = *nFlag
	}

	var filename string
	if flag.NArg() > 0 {
		filename = flag.Arg(0)
	} else {
		fmt.Println("Ошибка: Не указан файл.")
		printUsage()
		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %s", err)
		return
	}

	if nBytes > 0 {
		if nBytes > len(data) {
			nBytes = len(data)
		}
		fmt.Print(string(data[:nBytes]))
		return
	}

	if nLines > 0 {
		lines := strings.Split(string(data), "\n")
		if nLines > len(lines) {
			nLines = len(lines)
		}
		for i := 0; i < nLines; i++ {
			fmt.Println(lines[i])
		}
		return
	}

	fmt.Println("Ошибка: Не указано количество байтов или строк для вывода.")
	printUsage()
}
