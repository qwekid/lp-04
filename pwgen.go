package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Функция для генерации случайного пароля
func generatePassword(length int, strong bool, numeric bool, capitalize bool) string {
	var charset string
	if numeric {
		charset = "0123456789"
	} else if strong {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	} else {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	// Генерация пароля
	rand.Seed(time.Now().UnixNano())
	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password.WriteByte(charset[randomIndex])
	}

	// Если флаг capitalize установлен, делаем первую букву заглавной
	if capitalize && len(password.String()) > 0 {
		passwordChars := []rune(password.String())
		passwordChars[0] = rune(strings.ToUpper(string(passwordChars[0]))[0])
		return string(passwordChars)
	}

	return password.String()
}

func main() {
	// Флаги командной строки
	length := flag.Int("length", 8, "Длина пароля (по умолчанию 8 символов)")
	strong := flag.Bool("s", false, "Генерировать сильный пароль (смешанные символы)")
	numeric := flag.Bool("n", false, "Генерировать числовой пароль")
	capitalize := flag.Bool("c", false, "Генерировать пароль с заглавной буквы")
	help := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы командной строки
	flag.Parse()

	// Если пользователь запросил справку
	if *help {
		fmt.Println("Программа pwgen для генерации случайных паролей.")
		fmt.Println("Ключи:")
		fmt.Println("  -length        Длина пароля (по умолчанию 8 символов).")
		fmt.Println("  -s             Генерировать сильный пароль (смешанные символы).")
		fmt.Println("  -n             Генерировать числовой пароль.")
		fmt.Println("  -c             Генерировать пароль с заглавной буквы.")
		fmt.Println("  --help         Показать справку.")
		return
	}

	// Генерируем пароль
	password := generatePassword(*length, *strong, *numeric, *capitalize)

	// Выводим результат
	fmt.Println("Сгенерированный пароль:", password)
}
