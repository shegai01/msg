## Messenger Backend (Mvp)

Backend - сервим для обмена сообщенями на GOlang

## Tech Stack
- Golang
- net/http
- PostgreSQL
- Redis
- Docker

## Project Structure

## Configuration
    Сервис настраивается через переменные окружения 
    variable: HTTP_PORT -> Description: Http server port

## Run Locally
``` bash
export HTTP_PORT=8080
make run ("go run cmd/api/main.go") 