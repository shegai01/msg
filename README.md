## Messenger Backend (Mvp)

Backend - сервим для обмена сообщенями на GOlang

## Tech Stack
- Golang
- net/http
- PostgreSQL
- Redis
- Docker

## Project Structure
 cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   └── handlers
│       └── handlers.go
├── Makefile
├── migrations
└── README.md



## Configuration
    Сервис настраивается через переменные окружения 
    variable: HTTP_PORT -> Description: Http server port

## Run Locally
## HealthCHeck
GET /health

 bash
export HTTP_PORT=8080
make run ("go run cmd/api/main.go") 
