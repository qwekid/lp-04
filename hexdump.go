package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// Функция для печати данных в формате hexdump с ASCII в правой части
func hexdump(filePath string, numBytes int, showAscii bool) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Создаем буфер
	buf := make([]byte, numBytes)
	offset := 0

	for {
		// Читаем данные из файла
		n, err := file.Read(buf)
		if n == 0 && err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return fmt.Errorf("ошибка при чтении файла: %v", err)
		}

		// Печать адреса
		fmt.Printf("%08x  ", offset)

		// Печать данных в шестнадцатеричном формате
		for i := 0; i < n; i++ {
			fmt.Printf("%02x ", buf[i])
		}

		// Печать пробела для выравнивания
		for i := n; i < 16; i++ {
			fmt.Printf("   ")
		}

		// Печать ASCII, если нужно
		if showAscii {
			fmt.Print(" |")
			for i := 0; i < n; i++ {
				if buf[i] >= 32 && buf[i] <= 126 {
					fmt.Printf("%c", buf[i])
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("|")
		}

		fmt.Println()

		// Обновляем смещение
		offset += n
	}

	return nil
}

func main() {
	// Флаги командной строки
	filePath := flag.String("file", "", "Путь к файлу для дампа (обязательный)")
	numBytes := flag.Int("n", 0, "Число байтов для вывода")
	showAscii := flag.Bool("C", false, "Показывать ASCII-эквивалент данных")
	help := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы командной строки
	flag.Parse()

	// Если пользователь запросил справку
	if *help {
		fmt.Println("Программа hexdump для отображения данных файла в шестнадцатеричном формате.")
		fmt.Println("Ключи:")
		fmt.Println("  -file <путь>    Путь к файлу для дампа.")
		fmt.Println("  -n <число>      Число байтов для вывода (по умолчанию все данные).")
		fmt.Println("  -C              Показывать ASCII-эквивалент данных.")
		fmt.Println("  --help          Показать справку.")
		return
	}

	// Проверка на наличие обязательного параметра
	if *filePath == "" {
		fmt.Println("Ошибка: необходимо указать путь к файлу с помощью флага -file.")
		fmt.Println("Используйте --help для получения справки.")
		return
	}

	// Если не указано количество байтов, выводим все
	if *numBytes == 0 {
		*numBytes = 16 // стандартный размер строки в hexdump
	}

	// Генерируем hexdump
	err := hexdump(*filePath, *numBytes, *showAscii)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
