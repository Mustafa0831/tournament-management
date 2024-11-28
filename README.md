MIT License

Copyright (c) 2024 Mustafa

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

# Tournament Management

## Описание проекта

"Tournament Management" — это микросервисное приложение на языке Go для управления турнирами. Проект состоит из нескольких сервисов, каждый из которых отвечает за определенную функциональность:

- **team-service**: управление командами
- **division-service**: управление дивизионами
- **match-service**: управление матчами
- **playoff-service**: управление этапом плей-офф
- **result-generation-service**: генерация и обработка результатов
- **api-gateway**: API шлюз, объединяющий все сервисы в единую точку входа

Все сервисы взаимодействуют с базой данных PostgreSQL и общаются между собой через API Gateway.

## Требования

- Go версии 1.17 или выше
- Docker и Docker Compose
- Git для клонирования репозитория

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/ваш-логин/tournament-management.git
cd tournament-management
```

## Использование

### 1\. API Gateway

API Gateway доступен по адресу `http://localhost:8080`. Все запросы к сервисам проходят через него.

### 2\. Работа с team-service

*   **Получение списка команд:**
    
    ```bash
    curl http://localhost:8080/teams
    ```
    
*   **Добавление новой команды:**
    
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Team A"}' http://localhost:8080/teams
    ```

### 3\. Работа с другими сервисами

#### division-service

*   **Получение списка дивизионов:**
    
    ```bash
    curl http://localhost:8080/divisions
    ```
    
*   **Создание нового дивизиона:**
    
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Division A"}' http://localhost:8080/divisions
    ```

#### match-service

*   **Получение списка матчей:**
    
    ```bash
    curl http://localhost:8080/matches
    ```
    
*   **Создание нового матча:**
    
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"team1_id": 1, "team2_id": 2, "date": "2024-11-08T10:00:00Z"}' http://localhost:8080/matches
    ```

#### playoff-service

*   **Получение информации о плей-офф:**
    
    ```bash
    curl http://localhost:8080/playoffs
    ```
    
*   **Создание плей-офф:**
    
    ```bash
    curl -X POST http://localhost:8080/playoffs
    ```

#### result-generation-service

*   **Генерация результатов для Division A:**
    
    ```bash
    curl -X POST http://localhost:8080/generate-results/divisionA
    ```
    
*   **Генерация результатов для плей-офф:**
    
    ```bash
    curl -X POST http://localhost:8080/generate-results/playoff
    ```

## Makefile и Docker Compose

### Makefile

Для удобного управления сервисами используется Makefile с основными командами:

```makefile
.PHONY: build up down restart logs

build:
        docker-compose build

up:
        docker-compose up -d

down:
        docker-compose down

restart: down up

logs:
        docker-compose logs -f
