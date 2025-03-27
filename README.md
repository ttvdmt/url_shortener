# URL Shortener

Сервис по сокращению длинных URL-адресов, написанный на Golang

## Особенности
- Сокращение длинных URL в короткие шестизначные идентификаторы
- Редирект по полученным коротким ссылкам
- Настраиваемое время жизни коротких ссылок (TTL)
- Проверка исходных URL на доступность, корректность и безопасность
## Установка и запуск
Установка происходит в несколько этапов:
```bash
# Клонирование репозитория
git clone https://github.com/ttvdmt/url_shortener
cd url-shortener

# Установка зависимостей
go mod download
```

Перед запуском остается внести изменения в config.yaml:
```yaml
# config.yaml
app:                         //параметры сервиса
  host: "localhost"          //хост сервера
  port: "8080"               //порт сервера
  ttl_default: "720h"        //значение ttl, выставялемое по умолчанию, формат: XXXh
  cleanup_period: "24h"      //период проверки актуальности ссылок в БД, формат: XXXh

storage:                     //параметры БД
  sqlite:                    //параметры для работы с SQLite
    database: ""

  postgresql:                //параметры для работы с PostegreSQL
    host: "localhost"       
    port: "5432"              
    user: "user"          
    password: "password"  
    dbname: "db"          
    sslmode: "disable"      
```
При одновременной установке и параметров SQLite, и параметров PostegreSQL приоритетней будет работа с SQLite

На этом настройки заканчиваются и можно запускать:
```bash
# Запуск сервера
go run cmd\url_shortener\main.go
```
## Использование
Для сокращения URL-адреса нужно выполнить следующую команду:
```bash
curl -X POST -d "url=gianturl&ttl=48h" http://host:port/create
```
Поле ttl необязательное, без него для запроса будет использоваться ttl_default  

Полученную ссылку можно использовать, либо наппрямую в Интернете, либо в терминале:
```bash
curl -v http://host:port/success
```
