package main

import (
"bufio"
"flag"
"fmt"
"os"
)

// Функция для чтения последних N строк из файла
func tailFile(filename string, n int, follow bool) error {
file, err := os.Open(filename)
if err != nil {
return fmt.Errorf("не удалось открыть файл: %v", err)
}
defer file.Close()

var lines []string
scanner := bufio.NewScanner(file)

// Читаем все строки и сохраняем их в слайс
for scanner.Scan() {
	lines = append(lines, scanner.Text())
}

if err := scanner.Err(); err != nil {
	return fmt.Errorf("ошибка при чтении файла: %v", err)
}

// Если строк меньше, чем нужно, то просто выводим все
start := len(lines) - n
if start < 0 {
	start = 0
}

// Выводим последние N строк
for _, line := range lines[start:] {
	fmt.Println(line)
}

// Если ключ follow активирован, продолжаем выводить новые строки
if follow {
	fmt.Println("\nОжидаем новых строк...")
	tailFollow(file, n)
}

return nil
}

// Функция для следования за новыми строками в файле
func tailFollow(file *os.File, n int) {
// Используем bufio.NewReader для получения новых строк
reader := bufio.NewReader(file)
for {
line, err := reader.ReadString('\n')
if err != nil {
if err.Error() != "EOF" {
fmt.Println("Ошибка при чтении новых строк:", err)
}
break
}
fmt.Print(line)
}
}

func main() {
// Параметры командной строки
lines := flag.Int("n", 10, "Количество строк для отображения (по умолчанию 10)")
follow := flag.Bool("f", false, "Следить за изменениями файла (аналогично 'tail -f')")
help := flag.Bool("help", false, "Показать справку")

flag.Parse()

if *help {
	fmt.Println("Программа tail для отображения последних N строк из файла.")
	fmt.Println("Ключи:")
	fmt.Println("  -n <число>    Количество строк для отображения (по умолчанию 10).")
	fmt.Println("  -f            Следить за изменениями файла.")
	fmt.Println("  --help        Показать справку.")
	return
}

// Проверка наличия аргументов
if len(flag.Args()) == 0 {
	fmt.Println("Ошибка: не указан файл.")
	fmt.Println("Используйте --help для получения справки.")
	return
}

filename := flag.Args()[0]

// Выполняем команду tail
err := tailFile(filename, *lines, *follow)
if err != nil {
	fmt.Println("Ошибка:", err)
}
}