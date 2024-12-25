package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-mkdir [options] <directory>")
		os.Exit(1)
	}

	var mode os.FileMode = 0755 // По умолчанию права для новой директории
	var parents bool
	var verbose bool

	// Обработка аргументов
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch {
		case arg == "-p":
			parents = true
		case arg == "-v":
			verbose = true
		case strings.HasPrefix(arg, "-m"):
			if len(arg) > 2 {
				perm, err := strconv.ParseUint(arg[2:], 8, 32)
				if err == nil {
					mode = os.FileMode(perm)
				} else {
					fmt.Printf("Invalid mode: %s", arg[2:])
					os.Exit(1)
				}
			}
		default:
			// Это аргумент для директории
			if err := createDir(arg, mode, parents, verbose); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

func createDir(path string, mode os.FileMode, parents bool, verbose bool) error {
	if parents {
		// Создаем директорию с промежуточными каталогами
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Printf("Created directory: %s", path)
		}
	} else {
		// Создаем только конечную директорию
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
		if verbose {
			fmt.Printf("Created directory: %s", path)
		}
	}
	return nil
}
