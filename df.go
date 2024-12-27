package main

import (
	"fmt"
	"os"
	"syscall"
)

func printHelp() {
	fmt.Println("Использование: df [ОПЦИИ]")
	fmt.Println("Вывод информации о файловых системах.")
	fmt.Println()
	fmt.Println("ОПЦИИ:")
	fmt.Println("  -h       Показать размеры в удобочитаемом формате.")
	fmt.Println("  -T       Показать тип файловой системы.")
	fmt.Println("  --help   Показать эту справку и выйти.")
}

func getFilesystemType(fsPath string) (string, error) {
	var stat syscall.Stat_t
	err := syscall.Stat(fsPath, &stat)
	if err != nil {
		return "", err
	}

	// Получаем тип файловой системы
	var statfs syscall.Statfs_t
	err = syscall.Statfs(fsPath, &statfs)
	if err != nil {
		return "", err
	}

	// Используем f_type для определения типа файловой системы
	switch statfs.Type {
	case 0xEF53: // ext2/ext3/ext4
		return "ext4", nil
	case 0xFF534D42: // NTFS
		return "ntfs", nil
	case 0x9FA0: // tmpfs
		return "tmpfs", nil
	default:
		return "неизвестно", nil
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printHelp()
		return
	}

	humanReadable := false
	showType := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h":
			humanReadable = true
		case "-T":
			showType = true
		default:
			printHelp()
			return
		}
	}

	// Получаем информацию о файловых системах
	fs := syscall.Statfs_t{}
	filesystems := []string{"/", "/home", "/var"} // Пример файловых систем

	fmt.Printf("%-20s %-10s %-10s %-10s \n", "Файловая система", "Размер", "Использовано", "Доступно")
	for _, fsPath := range filesystems {
		err := syscall.Statfs(fsPath, &fs)
		if err != nil {
			fmt.Printf("Ошибка получения информации о %s: %v \n", fsPath, err)
			continue
		}

		total := fs.Blocks * uint64(fs.Bsize)
		free := fs.Bfree * uint64(fs.Bsize)
		used := total - free

		var totalStr, usedStr, freeStr string
		if humanReadable {
			totalStr = fmt.Sprintf("%.1fG", float64(total)/1024/1024)
			usedStr = fmt.Sprintf("%.1fG", float64(used)/1024/1024)
			freeStr = fmt.Sprintf("%.1fG", float64(free)/1024/1024)
		} else {
			totalStr = fmt.Sprintf("%d", total)
			usedStr = fmt.Sprintf("%d", used)
			freeStr = fmt.Sprintf("%d", free)
		}

		var fsType string
		if showType {
			fsType, err = getFilesystemType(fsPath)
			if err != nil {
				fmt.Printf("Ошибка получения типа файловой системы для %s: %v \n", fsPath, err)
				fsType = "неизвестно"
			}
			fmt.Printf("%-20s %-10s %-10s %-10s %-10s \n", fsPath, totalStr, usedStr, freeStr, fsType)
		} else {
			fmt.Printf("%-20s %-10s %-10s %-10s \n", fsPath, totalStr, usedStr, freeStr)
		}
	}
}
