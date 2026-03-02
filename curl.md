# Запросы для проверки добавления нового url

# Технические ошибки
# Проверка метода. Должен быть POST
Успешный сценарий:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" -d "https://habr.ru/" http://localhost:8080/

Негативный сценарий:
curl -v --request GET -H "Content-Type: text/plain; charset=utf-8" -d "https://vlgu.ru/" http://localhost:8080/  
curl -v --request DELETE -H "Content-Type: text/plain; charset=utf-8" -d "https://mail.ru/" http://localhost:8080/

# Проверка пути. Должен быть /
Успешный сценарий:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" -d "https://rambler.ru/" http://localhost:8080/

Негативный сценарий:
curl -v --request POST -H "Content-Type: text/plain" -d "https://yandex.ru/" http://localhost:8080/api

# Проверка заголовка. Должен быть text/plain
Успешный сценарий:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" -d "https://auto.ru/" http://localhost:8080/

Негативный сценарий:
curl -v --request POST -H "Content-Type: application/json; charset=utf-8" -d "https://avito.ru/" http://localhost:8080/

# Проверка тела запроса. Должен быть не пустым
Успешный сценарий:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" -d "https://kinopoisk.ru/" http://localhost:8080/

Негативный сценарий:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" -d "" http://localhost:8080/
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/

# Логические ошибки
# Проверка повторного добавления

Успешный сценарий. Ошибки не бывает:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" -d "https://kinopoisk.ru/" http://localhost:8080/

# #####################################################################################################################

# Запросы для проверки извлечения оригинальных url

# Технические ошибки
# Проверка метода. Должен быть GET
Успешный сценарий:
curl -v --request GET -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/3 

Негативный сценарий:
curl -v --request POST -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/4  
curl -v --request DELETE -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/2

# Проверка наличия короткой ссылки (корректного End Point-a)
Успешный сценарий:
curl -v --request GET -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/2

Негативный сценарий:
curl -v --request GET -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/
curl -v --request GET -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080

# Логические ошибки
# Проверяем запрос с несуществующим shortlink
curl -v --request GET -H "Content-Type: text/plain; charset=utf-8" http://localhost:8080/10