package main

import (
"flag"
"fmt"
"syscall"
)

// Структура для хранения информации о файловой системе
type FileSystemInfo struct {
Filesystem string
Type string
Size uint64
Used uint64
Free uint64
UsagePerc float64
InodesTotal uint64
InodesUsed uint64
InodesFree uint64
}

// Константы для единиц измерения
const (
KB = 1024
MB = KB * 1024
GB = MB * 1024
)

// Функция для конвертации размера в читаемый формат
func humanReadableSize(size uint64) string {
switch {
case size >= GB:
return fmt.Sprintf("%.1fG", float64(size)/float64(GB))
case size >= MB:
return fmt.Sprintf("%.1fM", float64(size)/float64(MB))
case size >= KB:
return fmt.Sprintf("%.1fK", float64(size)/float64(KB))
default:
return fmt.Sprintf("%dB", size)
}
}

// Получаем информацию о файловых системах с помощью syscall.Statfs
func getFileSystemInfo(path string) (FileSystemInfo, error) {
var stat syscall.Statfs_t
err := syscall.Statfs(path, &stat)
if err != nil {
return FileSystemInfo{}, fmt.Errorf("не удалось получить информацию о файловой системе для %s: %v", path, err)
}

fs := FileSystemInfo{}
fs.Filesystem = path
fs.Size = uint64(stat.Blocks) * uint64(stat.Bsize)
fs.Free = uint64(stat.Bfree) * uint64(stat.Bsize)
fs.Used = fs.Size - fs.Free
fs.UsagePerc = float64(fs.Used) / float64(fs.Size) * 100
fs.InodesTotal = uint64(stat.Files)
fs.InodesFree = uint64(stat.Ffree)
fs.InodesUsed = fs.InodesTotal - fs.InodesFree

return fs, nil
}

// Выводим информацию о файловых системах
func printFileSystems(fs FileSystemInfo, humanReadable bool, showInodes bool) {
if humanReadable {
fmt.Printf("%-20s %-20s %-20s %-20s %-20s\n", fs.Filesystem, humanReadableSize(fs.Size), humanReadableSize(fs.Free), humanReadableSize(fs.Used), fmt.Sprintf("%.2f%%", fs.UsagePerc))
} else {
fmt.Printf("%-20s %-20d %-20d %-20d %-20.2f%%\n", fs.Filesystem, fs.Size, fs.Free, fs.Used, fs.UsagePerc)
}

if showInodes {
	fmt.Printf("  Inodes: %d (занято %d, свободно %d)\n", fs.InodesTotal, fs.InodesUsed, fs.InodesFree)
}
fmt.Println()
}

func printHelp() {
fmt.Println("Использование: df [опции] [путь]")
fmt.Println("Опции:")
fmt.Println(" -h Выводить размеры в удобочитаемом формате (например, 1K, 234M)")
fmt.Println(" -i Показывать информацию о индексных узлах (inodes)")
fmt.Println(" -help Показать справку")
}

func main() {
// Флаги командной строки
showHelp := flag.Bool("help", false, "Показать справку")
humanReadable := flag.Bool("h", false, "Выводить размеры в удобочитаемом формате (например, 1K, 234M)")
showInodes := flag.Bool("i", false, "Показывать информацию о индексных узлах (inodes)")

// Парсим флаги
flag.Parse()

// Показать справку
if *showHelp {
	printHelp()
	return
}

// Если путь не передан, используем текущую директорию
path := "."
if len(flag.Args()) > 0 {
	path = flag.Args()[0]
}

// Получаем информацию о файловой системе
fs, err := getFileSystemInfo(path)
if err != nil {
	fmt.Println("Ошибка:", err)
	return
}

// Выводим информацию о файловой системе
printFileSystems(fs, *humanReadable, *showInodes)
}