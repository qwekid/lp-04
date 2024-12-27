# Учебная практика по УП04
Разработка аналогов команд линукс с помощьюязыка GO

# Список поддерживаемых команд
    ls с ключами -R, -r, -a, -l, -h,
    pwd


## Как использовать ls
Требуется запустить исполняемый файл ls
```
./cd [-R] [-r] [-a] [-l] [-h] [-help] <путь>
```
### Поддерживаемые флаги для ls
	-R - рекурсивный вывод
	-r - обратный порядок 
	-a - включить скрытые файлы в вывод
    -l - длинный формат
	-h - читаемый размер
    -help - вывод поддерживаемых команд с описанием их действия

## Как использовать pwd
Требуется запустить исполняемый файл pwd
```
./pwd
```
### Поддерживаемые флаги для pwd
	-L - выводит логический путь
    -P - выводит физический путь
    -help - вывод поддерживаемых команд с описанием их действия

## Как использовать cd
Требуется запустить исполняемый файл cd
```
eval $(./cd [-P] [-L] [-v] [-h] <путь>)
```
### Поддерживаемые флаги для cd
	-P - Использовать физический путь (без символических ссылок)
	-L - Использовать логический путь (учитывая символические ссылки)
	-v - Показать версию программы
	-h - Показать это сообщение справки

## Как использовать mkdir
Требуется запустить исполняемый файл mkdir
```
./mkdir [-p] [-v] [-m <права доступа>] <путь>
```
### Поддерживаемые флаги для mkdir
	-p: Создает промежуточные каталоги, если они не существуют
	-v: Включает режим подробного вывода, который сообщает о создании директории
	-m: задать права доступа для создаваемой директории в восьмеричном формате

## Как использовать rmdir
Требуется запустить исполняемый файл rmdir
```
./mkdir [-p] [-v] [-m <права доступа>] <путь>
```
### Поддерживаемые флаги для mkdir
    -h - Показать справку
    -ignore-fail-on-non-empty (-p) - Игнорировать ошибки, когда каталог не пуст
    -v - Включить подробный вывод

## Как использовать cat
Требуется запустить исполняемый файл cat
```
./cat [-E] [-b] [-n] <путь/к/файлу>
```
### Поддерживаемые флаги для cat
    -E - Добавлять символ $ в конце каждой строки
    -b - Нумеровать непустые строки
    -n - Нумеровать все строки

## Как использовать file
Требуется запустить исполняемый файл file
```
./file [-b] [-i] [-h] [файлы]
```
### Поддерживаемые флаги для file
    -b - Выводит тип файла в виде MIME-медиа-типов.
    -i - Выводит дополнительную информацию о файле.
    -h - Выводит справку по использованию.

## Как использовать nl
Требуется запустить исполняемый файл nl
```
./nl [-b, --body] [-f, --footer] [-h, -- help] [файлы]
```
### Поддерживаемые флаги для nl
    -b, --body - Нумеровать только строки тела
    -f, --footer - Нумеровать только нижний колонтитул
	-h, --help - Показать это сообщение

## Как использовать zip
Требуется запустить исполняемый файл zip
```
./zip [-h, ---help] [-r] [-q] [-v] [архив] [файлы...]
```
### Поддерживаемые флаги для zip
    -h, --help - показать это сообщение
    -r - рекурсивно добавлять файлы из подкаталогов
	-q - тихий режим (меньше сообщений)
	-v - показать подробную информацию о процессе

## Как использовать unzip
Требуется запустить исполняемый файл unzip
```
./unzip [-h] [-o] [-q] путь/к/файлу.zip
```
### Поддерживаемые флаги для unzip
    -h, --help     Показать это сообщение
	-o, --output   Путь к директории для распаковки
	-q, --quiet    Не выводить сообщения

    
## Как использовать tar
Требуется запустить исполняемый файл tar
```
./tar [-c] [-x] -f путь/к/файлу.zip
```
### Поддерживаемые флаги для tar
  	-c - Создать архив
    -x - Извлечь архив
	-f <файл> -  Имя файла архива
	--help - Показать помощь
    
    
## Как использовать !!
Требуется запустить исполняемый файл !!
```
./\!\!
```
### Если требуется выполнить предыдущую команду
```
eval $(./\!\!)
```

## Как использовать !n
Требуется запустить исполняемый файл !n
```
./\!\n [-echo <текст>] [-count <строка>] [-reverse <строка>] [-help]
```
### Поддерживаемые флаги для !n
 	-echo - выводит сообщение обратно
	-count - выводит количество символов в строке
	-reverse - выводит строку в обратном порядке
	-help - выводит это сообщение

## Как использовать head
Требуется запустить исполняемый файл head
```
./head [-c <число>] [-b <число>] [-n <число>] [--help]
```
### Поддерживаемые флаги для head
 	-c N - Вывести первые N байтов файла
	-b N - Вывести первые N байтов файла (аналогично -c)
	-n N - Вывести первые N строк файла
	--help - Показать это сообщение

## Как использовать touch
Требуется запустить исполняемый файл touch
```
./head [-a] [-m] [-c] [--help] имя/файла
```
### Поддерживаемые флаги для touch
 	-a - Изменить время последнего доступа.
	-m - Изменить время последней модификации.
	-c - Не создавать файл, если он не существует.
	--help - Показать это сообщение.

## Как использовать free
Требуется запустить исполняемый файл free
```
./free [-m] [-t] [-u]
```
### Поддерживаемые флаги для free
 	-m - Отображение колчества памяти в целом
    -t - Отображение количества используемой памяти
    -u - Отображение количества свободной памяти

## Как использовать exit
Требуется запустить исполняемый файл exit
```
./exit [<число>] [-help]
```
### Поддерживаемые флаги для exit
 	-help - показать справку

## Как использовать rm
Требуется запустить исполняемый файл rm
```
./rm [-help] [-R] [-f] путь/к/файлу 
```
### Поддерживаемые флаги для rm
 	-help - показать справку
	-R - рекурсивное удаление
	-а - игнорировать ошибки

## Как использовать date
Требуется запустить исполняемый файл date
```
./date [-help] [-I] [-d] [-t] 
```
### Поддерживаемые флаги для date
 	-help - Показать справку
  	-I - Показать дату и время в ISO 8601 формате
  	-d - Показать только дату
  	-t - Показать только время


## Как использовать arch
Требуется запустить исполняемый файл arch
```
./arch [--help]
```
### Поддерживаемые флаги для arch
 	-help - Показать справку

## Как использовать clear
Требуется запустить исполняемый файл clear
```
./clear [-help] [-s] [-h] [-m <сообщение>]
```
### Поддерживаемые флаги для clear
 	-help - Показать справку
 	-s - Мягкая очистка экрана (перенос строки без полной очистки)
 	-h - Очистить экран и историю терминала (не во всех терминалах)
 	-m <сообщение> - Очистить экран и вывести сообщение


## Как использовать ps
Требуется запустить исполняемый файл ps
```
./ps [-e] [-f] [-u]
```
### Поддерживаемые флаги для ps
 	-e - Показать все процессы
 	-f - Показать дополнительные данные о процессах
 	-u <пользователь> - Фильтровать процессы по пользователю (UID или имяпользователя)


## Как использовать du
Требуется запустить исполняемый файл du
```
./du [-s] [-a] [-help]
```
### Поддерживаемые флаги для du
 	-s - Показывать только общий размер для каждого аргумента
 	-a - Показывать размер всех файлов, а не только директорий
 	-help - Показать справку

## Как использовать df
Требуется запустить исполняемый файл df
```
./df [-h] [-i] [-help]
```
### Поддерживаемые флаги для df
 	-h - Выводить размеры в удобочитаемом формате (например, 1K, 234M)
 	-i - Показывать информацию о индексных узлах (inodes)
 	-help - Показать справку

 ## Как использовать wc
Требуется запустить исполняемый файл wc
```
./wc [-l] [-w] [-c] [-help] /путь/к/файлу
```
### Поддерживаемые флаги для wc
 	-l Подсчитать количество строк
 	-w Подсчитать количество слов
 	-c Подсчитать количество байтов
 	-help Показать справку

 ## Как использовать tail
Требуется запустить исполняемый файл tail
```
./tail [-n <число>] [-f] [--help] /путь/к/файлу
```
### Поддерживаемые флаги для tail
	-n <число> - Количество строк для отображения (по умолчанию 10).
  	-f - Следить за изменениями файла.
  	--help - Показать справку.

 ## Как использовать uname
Требуется запустить исполняемый файл uname
```
./uname [-a] [-s] [-n] [-v] [--help]
```
### Поддерживаемые флаги для uname
	-a - Вывести всю информацию о системе.
  	-s - Имя операционной системы.
  	-n - Имя узла.
  	-v - Версия ядра.
  	--help - Показать справку.








    