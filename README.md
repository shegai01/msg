## Messenger Backend (MVP)

Backend - сервис mvp для обмена сообщенями на Golang
постепенным усложнением (rest,postgresql,redis,acync service)
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
## HealthCHeck
- GET /health
- response {"status":"ok"}

 bash
export HTTP_PORT=8080
make run ("go run cmd/api/main.go") 
