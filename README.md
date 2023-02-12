# keeper - система для хранения приватных данных

- [Общая информация](#info)
- [Архитектура](#arch)
  - [Блок-схемы](#arch-scheme)
    - [Регистрация/Логин](#arch-scheme-auth)
    - [Добавление приватных данных](#arch-scheme-privatedata)
    - [Загрузка файлов](#arch-scheme-files)
  - [Схема БД](#arch-db)

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
