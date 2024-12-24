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

## Как использовать св
Требуется запустить исполняемый файл cd
```
eval $(./cd [-P] [-L] [-v] [-h] <путь>)
```
### Поддерживаемые флаги для pwd
	-L - выводит логический путь
    -P - выводит физический путь
    -help - вывод поддерживаемых команд с описанием их действия