# 💰 FinTech MVP — REST API для управления личными финансами

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> **Учебный проект** — разработан в рамках производственной практики по направлению «Разработка программного обеспечения для проектов, связанных с финансовыми технологиями».

---

## 📌 О проекте

Этот проект — MVP (Minimum Viable Product) финансового сервиса. Он позволяет:
- ✅ Регистрировать и авторизовать пользователей
- ✅ Добавлять доходы
- ✅ Просматривать текущий баланс

**Технологии:** Go, Gin, JWT, bcrypt

---

## 🚀 Быстрый старт

### Требования
- Go 1.21 или выше

### Установка и запуск


# Клонируем репозиторий
git clone https://github.com/WowkaJones/fintech-mvp.git
cd fintech-mvp

# Запускаем сервер
go run main.go

Сервер запустится на http://localhost:8080

## 📡 API Эндпоинты
Метод	URL	Описание
POST	/register	Регистрация пользователя
POST	/login	Вход и получение JWT токена
POST	/add-income	Добавление дохода
Примеры запросов
# Регистрация:
curl -X POST http://localhost:8080/register \
  -d "email=test@mail.ru" \
  -d "password=123"
# Логин:
  curl -X POST http://localhost:8080/login \
  -d "email=test@mail.ru" \
  -d "password=123"
# Добавить доход (с токеном):
  curl -X POST http://localhost:8080/add-income \
  -H "Authorization: ТВОЙ_ТОКЕН" \
  -H "Content-Type: application/json" \
  -d '{"amount": 500}'
## 🛠 Технологии
Компонент	Технология
Язык	Go 1.21
Веб-фреймворк	Gin
Авторизация	JWT
Хэширование	bcrypt
## 👤 Автор
Студент: [Захаров Владиир Игорьевич]
GitHub: @WowkaJones
Практика: Разработка ПО для FinTech-проектов

## ✅ Проект успешно прошёл этапы: анализ требований → проектирование → разработка → тестирование → документация.
