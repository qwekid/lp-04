package main

import (
	"flag"
	"fmt"
	"time"
)

// Структура для хранения информации о выполненной команде
type HistoryEntry struct {
	Command   string
	Timestamp time.Time
}

// История команд
var history []HistoryEntry

// Функция для вывода истории команд
func printHistory(n int) {
	// Если n меньше или равно 0, показываем всю историю
	if n <= 0 || n > len(history) {
		n = len(history)
	}

	// Выводим последние n команд
	for i := len(history) - n; i < len(history); i++ {
		fmt.Printf("%d  %s  %s\n", i+1, history[i].Timestamp.Format("2006-01-02 15:04:05"), history[i].Command)
	}
}

// Функция для очистки истории
func clearHistory() {
	history = []HistoryEntry{}
	fmt.Println("История команд была очищена.")
}

func main() {
	// Флаги командной строки
	showCount := flag.Int("n", 0, "Показать последние n команд (аналогично 'history -n')")
	clear := flag.Bool("c", false, "Очистить историю команд (аналогично 'history -c')")
	help := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы командной строки
	flag.Parse()

	// Если пользователь запросил справку
	if *help {
		fmt.Println("Программа history для отображения и управления историей команд.")
		fmt.Println("Ключи:")
		fmt.Println("  -n n           Показать последние n команд.")
		fmt.Println("  -c             Очистить историю команд.")
		fmt.Println("  --help         Показать справку.")
		return
	}

	// Чтение и запись в историю (эмуляция истории команд)
	// Эмулируем выполнение команд в процессе работы программы.
	history = append(history, HistoryEntry{
		Command:   "ls -l",
		Timestamp: time.Now(),
	})
	history = append(history, HistoryEntry{
		Command:   "cd /home/user",
		Timestamp: time.Now(),
	})
	history = append(history, HistoryEntry{
		Command:   "cat file.txt",
		Timestamp: time.Now(),
	})

	// Обработка флагов
	if *clear {
		clearHistory()
		return
	}

	if *showCount > 0 {
		printHistory(*showCount)
		return
	}

	// Если ни один флаг не был передан, выводим всю историю
	printHistory(len(history))
}
