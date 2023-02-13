# keeper - система для хранения приватных данных

- [Общая информация](#info)
- [Архитектура](#arch)
  - [Блок-схемы](#arch-scheme)
    - [Регистрация/Логин](#arch-scheme-auth)
    - [Добавление приватных данных](#arch-scheme-privatedata)
    - [Загрузка файлов](#arch-scheme-files)
  - [Схема БД](#arch-db)

Работа приложения(#operation)
- [Конфигурация сервера](#operation-config-server)
- [Клиент](#operation-client)

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

Подробнее [migrations/keeper.sql](migrations/keeper.sql)


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
  
Использование: `$keeper-client --<command-flag> [<arg-flags>]`  
  
Сценарии:  
  
1. Регистрация пользователя:  
```
флаг-команда:
--register

   флаги-аргументы:
    --user  
    --password
    --metadata
```
Пример: `$ keeper-client --register --user=john --password=1a23456b --metadata="some data for more info"`
    
    
