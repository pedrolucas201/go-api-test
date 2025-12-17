# Go API Test

API simples em Go com endpoints GET e POST, incluindo testes unitários.

## Requisitos
- Go 1.25+ (conforme go.mod)

## Instalação e execução
```bash
# Instalar dependências
go mod tidy

# Executar a API
go run main.go
```
A API inicia na porta 8080.

## Endpoints
- GET /health
  - Retorna: { "status": "ok" }
- POST /message
  - Body (JSON): { "text": "sua mensagem" }
  - Retorna 201 com o mesmo JSON.

### Exemplos
```bash
curl -s http://localhost:3000/health
curl -s -X POST http://localhost:3000/message \
  -H 'Content-Type: application/json' \
  -d '{"text":"Olá"}'
```

## Testes
```bash
go test ./...
```
