# keeper - система для хранения приватных данных

- [Общая информация](#info)
- [Архитектура](#arch)
  - [Блок-схемы](#arch-scheme)
    - [Регистрация/Логин](#arch-scheme-auth)
    - [Добавление приватных данных](#arch-scheme-privatedata)
    - [Загрузка файлов](#arch-scheme-files)
  - [Схема БД](#arch-db)

- [Работа приложения](#operation)
  - [Конфигурация сервера](#operation-config-server)
  - [Клиент](#operation-client)
    - [Регистрация пользователя](#operation-client-register)
    - [Логин](#operation-client-login)
    - [Произвольные текстовые данные](#operation-client-text)
    - [Пары логин/пароль](#operation-client-pair)
    - [Банковские карты](#operation-client-bankcard)
    - [Файлы](#operation-client-file)
  - [Запуск приложения](#operation-run)
    - [Последовательность команд для запуска сервера](#operation-run-sequence)
    - [Запуск клиента](#operation-run-client)
- [TODO](#todo)

# Общая информация <a name="info"/>
keeper - система для хранения приватных данных: тексты, пароли, данные банковских карт, файлы

# Архитектура <a name="arch"/>

## Блок-схемы <a name="arch-scheme"/>

### Регистрация/Логин <a name="arch-scheme-auth"/>
![Регистрация/Логин](docs/arch-scheme-auth.png)


### Добавление приватных данных <a name="arch-scheme-privatedata"/>
![Регистрация/Логин](docs/arch-scheme-privatedata.png)


### Загрузка файлов <a name="arch-scheme-files"/>
![Регистрация/Логин](docs/arch-scheme-files.png)

## Схема БД <a name="arch-db"/>
![Схема БД](docs/arch-db.png)

Подробнее [migrations/01_keeper.sql](migrations/01_keeper.sql)


# Работа приложения <a name="operation"/>
## Конфигурация сервера <a name="operation-config-server"/>

Параметры конфигурация сервиса keeper определяются либо файлом конфигурации, либо флагами командной строки, либо переменными окружения.

| Переменная окружения           | Флаг командной строки | Описание                                      |
|--------------------------------|-----------------------|-----------------------------------------------|
| `CONFIG_KEEPER`                |`-c;--config <file>`   | путь к файлу конфигурации                     |
| `ADDRESS`                      | `-a <host:port>`      | адрес и порт сервиса                          |
| `DATABASE_DSN`                 | `-d <dsn>`            | адрес для подключения к базе данных           |
| `SESSION_KEY`                  | _нет_                 | ключ для аутентификаиции                      |

## Клиент <a name="operation-client"/>

Клиент представляет собой CLI-приложение.  
Сценарий работы клиета конфигурируется флагом-командой и соответствующим флагами-аргументами.  
  
Шаблон вызова: `$keeper-client --<command-flag> [<arg-flags>]`  
  
### Регистрация пользователя <a name="operation-client-register"/>  
```
флаг-команда:
--register

    флаги-аргументы:
    --user  
    --password  

Пример:  
$ keeper-client --register --user=john --password=1a23456b
```
  
### Создание сессии пользователя <a name="operation-client-login"/>  
```
флаг-команда:
--login

    флаги-аргументы:
    --user  
    --password

Пример:  
$ keeper-client --login --user=john --password=1a23456b
```  
  
### Произвольные текстовые данные <a name="operation-client-text"/>  
```
флаг-команда:
--text

    флаги-аргументы:
    --content  
    --metadata

Примеры:
// Добавить произвольный текст
$ keeper-client --text --content="some text" --metatdata="some metadata"

// Получить все тексты авторизованного пользователя
$ keeper-client --text
```  
  
### Пары логин/пароль <a name="operation-client-pair"/>  
```
флаг-команда:
--pair

    флаги-аргументы:
    --user
    --password
    --metadata

Примеры:
// Добавить пару логин/пароль
$ keeper-client --pair --user="newuser" --password="newpassword" --metatdata="some metadata for new user"

// Получить все пары логин/пароль
$ keeper-client --pair
```  
  
### Банковские карты <a name="operation-client-bankcard"/>  
```
флаг-команда:
--card

    флаги-аргументы:
    --pan           PAN - Номер платежной карты
    --till          Valid Thru Date - дата истечения срока действия карты, указанная на поверхности карты 
    --cvv           CVC/CVV — коды верификации пластикового носителя, которые подтверждают его подлинность
    --name          Имя владельца карты, напечатанное на лицевой стороне карты

    --metadata

Примеры:
// Добавить данные банковский карты
$ keeper-client --card --pan="4321 6543 3214 8766" --till="11/23" --cvv="123" --name="JOHN SMITH" --metatdata="some metadata for new user"

// Получить все данные банковских карт
$ keeper-client --card
```  

### Файлы <a name="operation-client-file"/>  
```
флаг-команда:
--file

    флаги-аргументы:
    --upload          загрузить файл на сервер
    --download        скачать файл

    --metadata
Примеры:
// Загрузить файл
$ keeper-client --file --upload=/path/to/file --metatdata="some file for upload"

// Скачать файл
$ keeper-client --file --download="file"

// Получить список загруженных на сервер файлов
$ keeper-client --file
```  
  
## Запуск приложения <a name="operation-run">  
Подготовка частей приложения, тестирование, и запуск выполняется с помощью команды make.  
В Makefile определены основные команды для управления приложением. 
  
Подробнее [Makefile](Makefile)  
  
### Последовательность команд для запуска сервера <a name="operation-run-sequence">  
- Создание и запуск контейнера с базой данных:  
`$ make docker-db-up`  
- Генерация бинарного файла сервера:  
`$ make build-server`  
- Запуск сервера с конфигурацией из файла `configs/keeper.json`:  
`$ make run-server`  
### Запуск клиента <a name="operation-run-client">  
- Генерация бинарного файла клиента:  
`$ make build-client`  
  
После генерации, бинарный файл клиента находится в директории `bin/`.  
Например, команда клиента для регистрации пользователя выглядит так:  
`$ ./bin/keeper-client --register --user=newuser --password=secret`  
  
# TODO <a name="todo"/>
1. Добавить в функционал системы возможность редактирования/удаления объектов.
2. Для клиента добавить поддержку терминального интерфейса
