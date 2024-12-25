package main

import (
"flag"
"fmt"
"syscall"
)

func main() {
var memFlag, totalFlag, usedFlag bool

flag.BoolVar(&memFlag, "m", false, "Отображение колчества памяти в целом")
flag.BoolVar(&totalFlag, "t", false, "Отображение количества используемой памяти")
flag.BoolVar(&usedFlag, "u", false, "Отображение количества свободной памяти")

flag.Parse()

if flag.NFlag() == 0 || flag.Arg(0) == "help" {
	fmt.Println("Программа для отображения информации о памяти.")
	fmt.Println("Использование:")
	flag.PrintDefaults()
	return
}

var info syscall.Sysinfo_t
err := syscall.Sysinfo(&info)
if err != nil {
	fmt.Println("Ошибка при получении информации о системе")
	return
}

if memFlag {
	fmt.Printf("Всего памяти: %v KB\n", info.Totalram/1024)
}

if totalFlag {
	fmt.Printf("Используется: %v KB\n", (info.Totalram-info.Freeram)/1024)
}

if usedFlag {
	fmt.Printf("Свободно: %v KB\n", info.Freeram/1024)
}
}