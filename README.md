## URL Shortener на Go + Gin + MySQL + Sqids
Простой сервис для сокращения ссылок написанный на 
Gin, БД MySql, а также библиотекой для создания ID для короткой ссылки - Sqids

##  Возможности
- Принимает длинный URL и генерирует короткий код.
- Хранит соответствие в базе данных.
- Все ответы в JSON формате
- Редиректа с короткой ссылки на оригинальную.

## Принцип работы
1.На вход принимается оригинальная ссылка в формате JSON 
    Пример:{"url":"www.youtube.com"}
2.Подсчитывается количество уже созданных строк + 1 и передается в GenerateUniqueID 
для генерации кода(короткой ссылки) для short_url
3.В БД вставляются значения короткой ссылки,оригинальной ссылки и даты добавления

## Запуск
1.Клонирование репозитория
```
git clone https://github.com/SH1roV12/urlshortener.git

```
Go сам подтянет все либы

2.Создание БД (требуется предустановленный MySQL!)
```
CREATE DATABASE urlshortener;
USE urlshortener;

CREATE TABLE urls (
    id INT AUTO_INCREMENT PRIMARY KEY,
    short_code VARCHAR(10) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
3.Создание .env файла по примеру(.env.example)
Имя и пароль который был указан при создании БД(имя по умолчанию root)
4.Запуск проект
```
go run main.go
```
Сервер запустится на http://localhost:8080 

