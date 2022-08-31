[![ci](https://github.com/ayanovskiy/banner-service/actions/workflows/tests.yaml/badge.svg)](https://github.com/ayanovskiy/banner-service/actions/workflows/tests.yaml)
# Проектная работа "Ротация баннеров"

## Общее описание сервиса
Сервис "Ротация баннеров" предназначен для выбора наиболее эффективных (кликабельных) баннеров,
в условиях меняющихся предпочтений пользователей и набора баннеров.

Ссылка на ТЗ - https://github.com/OtusGolang/final_project/blob/master/02-banners-rotation.md

## Основные команды для работы с проектом
1. make build - основная сборка проекта
2. make dev-build - сборка проекта для разработке
3. make run - запуск проекта
4. make test - запуск всех тестов в проекте
5. make test-uint - запуск unit-тестов
6. make test-integration - запуск интеграционных тестов
7. make lint - запуск линтера
8. make docker-build - сборка образов docker-compose
9. make docker-up - запуск docker-compose
10. make docker-down - остановка docker-compose

## API сервиса
Все запросы и ответы (кроме главной страницы) выполняются в JSON формате
* "/" - главная страница сервиса
* "/slot/add-banner" - добавление баннера в слот
* "/slot/remove-banner" - удаление баннера из слота
* "/banner/select" - выбор баннера для показа
* "/banner/hit" - переход по баннеру
* "/stats/send" - отправка статистики по показам и переходам баннеров в очередь

### "/slot/add-banner"
POST запрос: {"slot_id": int, "banner_id": int}

Ответ: {"response":"ok"} | {"error":<описание ошибки>}
### "/slot/remove-banner"
POST запрос: {"slot_id": int, "banner_id": int}

Ответ: {"response":"ok"} | {"error":<описание ошибки>}
### "/banner/select"
POST запрос: {"slot_id": int, "group_id": int}

Ответ: {"response":"ok"} | {"error":<описание ошибки>}
### "/banner/hit"
POST запрос: {"banner_id": int, "slot_id": int, "group_id": int}

Ответ: {"response":"ok"} | {"error":<описание ошибки>}
### "/stats/send"
GET запрос

Ответ: {"response":"ok"} | {"error":<описание ошибки>}
