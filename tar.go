package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Определение флагов
	create := flag.Bool("c", false, "Создать архив")
	extract := flag.Bool("x", false, "Извлечь архив")
	filename := flag.String("f", "", "Имя файла архива")
	help := flag.Bool("help", false, "Показать помощь")

	// Парсинг аргументов
	flag.Parse()

	if *help {
		fmt.Println("Использование:")
		fmt.Println("  -c        Создать архив")
		fmt.Println("  -x        Извлечь архив")
		fmt.Println("  -f <файл> Имя файла архива")
		fmt.Println("  --help    Показать помощь")
		return
	}

	if *create && *extract {
		fmt.Println("Ошибка: нельзя одновременно использовать -c и -x")
		return
	}

	if *create {
		if *filename == "" {
			fmt.Println("Ошибка: необходимо указать имя файла с помощью -f")
			return
		}
		err := createTar(*filename, flag.Args())
		if err != nil {
			fmt.Printf("Ошибка при создании архива: %v", err)
		} else {
			fmt.Println("Архив успешно создан:", *filename)
		}
	} else if *extract {
		if *filename == "" {
			fmt.Println("Ошибка: необходимо указать имя файла с помощью -f")
			return
		}
		err := extractTar(*filename)
		if err != nil {
			fmt.Printf("Ошибка при извлечении архива: %v", err)
		} else {
			fmt.Println("Архив успешно извлечен:", *filename)
		}
	} else {
		fmt.Println("Ошибка: необходимо использовать -c для создания архива или -x для извлечения")
	}
}

func createTar(tarFile string, files []string) error {
	out, err := os.Create(tarFile)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := tar.NewWriter(out)
	defer writer.Close()

	for _, file := range files {
		err := addFileToTar(writer, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func addFileToTar(writer *tar.Writer, file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = file

	if err := writer.WriteHeader(header); err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(writer, f)
	return err
}

func extractTar(tarFile string) error {
	file, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := tar.NewReader(file)

	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = extractFileFromTar(header, reader)
		if err != nil {
			return err
		}
	}
	return nil
}

func extractFileFromTar(header *tar.Header, reader io.Reader) error {
	outFile, err := os.Create(header.Name)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, reader)
	return err
}
