package main

import (
"flag"
"fmt"
"io/ioutil"
"os"
"path/filepath"
"strconv"
"strings"
)

// Структура для хранения информации о процессе
type ProcessInfo struct {
PID int
UID int
Command string
State string
User string
Args string
}

func main() {
// Флаги
showAll := flag.Bool("e", false, "Показать все процессы")
showFull := flag.Bool("f", false, "Показать дополнительные данные процессов")
userFilter := flag.String("u", "", "Фильтровать по пользователю (UID или имя пользователя)")
showHelp := flag.Bool("help", false, "Показать справку")

flag.Parse()

// Показать справку, если указан флаг --help
if *showHelp {
	printHelp()
	return
}

// Получаем список процессов
processes, err := getProcesses()
if err != nil {
	fmt.Println("Ошибка при получении списка процессов:", err)
	os.Exit(1)
}

// Фильтруем процессы по пользователю, если задан флаг -u
if *userFilter != "" {
	processes = filterByUser(processes, *userFilter)
}

// Показываем информацию о процессах в зависимости от флагов
for _, process := range processes {
	if *showFull {
		fmt.Printf("PID: %d, UID: %d, Пользователь: %s, Команда: %s, Состояние: %s, Аргументы: %s\n",
			process.PID, process.UID, process.User, process.Command, process.State, process.Args)
	} else {
		fmt.Printf("PID: %d, Команда: %s\n", process.PID, process.Command)
	}
}

if !*showAll && len(processes) == 0 {
	fmt.Println("Нет процессов для отображения.")
}
}

// Получение списка всех процессов
func getProcesses() ([]ProcessInfo, error) {
var processes []ProcessInfo

// Чтение всех директорий в /proc
procDirs, err := ioutil.ReadDir("/proc")
if err != nil {
	return nil, fmt.Errorf("не удалось прочитать /proc: %v", err)
}

for _, procDir := range procDirs {
	// Проверяем, является ли это директорией с номером процесса (цифры в названии)
	if procDir.IsDir() && isNumeric(procDir.Name()) {
		pid, err := strconv.Atoi(procDir.Name())
		if err != nil {
			continue
		}

		// Чтение данных из /proc/[pid]/stat
		statPath := filepath.Join("/proc", procDir.Name(), "stat")
		statData, err := ioutil.ReadFile(statPath)
		if err != nil {
			continue
		}

		// Разбор данных из stat
		statFields := strings.Fields(string(statData))
		if len(statFields) < 24 {
			continue
		}

		// Получаем PID, состояние и команду
		state := statFields[2] // Состояние процесса
		command := statFields[1] // Имя команды
		args := strings.Join(statFields[3:], " ") // Аргументы процесса

		// Получаем информацию о пользователе из /proc/[pid]/status
		statusPath := filepath.Join("/proc", procDir.Name(), "status")
		statusData, err := ioutil.ReadFile(statusPath)
		if err != nil {
			continue
		}

		// Разбор данных из status
		var uid int
		var username string
		for _, line := range strings.Split(string(statusData), "\n") {
			if strings.HasPrefix(line, "Uid:") {
				fields := strings.Fields(line)
				uid, _ = strconv.Atoi(fields[1])
				break
			}
			if strings.HasPrefix(line, "Name:") {
				fields := strings.Fields(line)
				username = fields[1]
			}
		}

		processes = append(processes, ProcessInfo{
			PID:     pid,
			UID:     uid,
			Command: command,
			State:   state,
			User:    username,
			Args:    args,
		})
	}
}
return processes, nil
}

// Фильтровать процессы по имени пользователя
func filterByUser(processes []ProcessInfo, user string) []ProcessInfo {
var filtered []ProcessInfo
for _, process := range processes {
if process.User == user {
filtered = append(filtered, process)
}
}
return filtered
}

// Проверка, является ли строка числом (для проверки PID)
func isNumeric(s string) bool {
_, err := strconv.Atoi(s)
return err == nil
}

// Функция для вывода справки
func printHelp() {
fmt.Println("Использование: ps [опции]")
fmt.Println("Опции:")
fmt.Println(" -e Показать все процессы")
fmt.Println(" -f Показать дополнительные данные о процессах")
fmt.Println(" -u <пользователь> Фильтровать процессы по пользователю (UID или имя пользователя)")
fmt.Println(" -help Показать справку")
}

