# Todo API

API REST para gerenciamento de tarefas (To-Do List), desenvolvida com Golang, Gin e MongoDB.

Este projeto foi criado com o objetivo de praticar conceitos de APIs RESTful, organização de camadas (handlers, services, repositories) e integração com banco de dados.

## Funcionalidades

- Criar tarefas
- Listar tarefas
- Buscar tarefa por ID
- Atualizar tarefas
- Deletar tarefas

## Regras de negócio implementadas

- O `title` precisa ter tamanho mínimo válido
- O `status` da tarefa pode ser apenas:
  - `pending`
  - `in_progress`
  - `completed`
  - `cancelled`
- A `priority` pode ser:
  - `low`
  - `medium`
  - `high`
- A `due_date` não pode estar no passado
- Tarefas concluídas (`completed`) não podem ser editadas
- Também é possível filtrar tarefas por status e prioridade.

## Tecnologias utilizadas

- Golang
- Gin Gonic
- MongoDB
- Docker / Docker Compose

## Como rodar o projeto

### Usando Docker (forma mais fácil)

Clone o repositório ou navegue até a pasta do projeto.

Suba os containers:

```bash
docker-compose up -d
```

A API ficará disponível em: `http://localhost:8080`  
MongoDB: `localhost:27017`

Para ver os logs:

```bash
docker-compose logs -f
```

Para parar os containers:

```bash
docker-compose down
```

### Rodando localmente (sem Docker)

Caso prefira rodar direto na máquina, certifique-se que o MongoDB está rodando em `localhost:27017`.

Instale as dependências:

```bash
go mod download
```

Rode a aplicação:

```bash
go run cmd/main.go
```

Ou gere o executável:

```bash
go build -o api.exe cmd/main.go
./api.exe
```

## Endpoints da API

### Criar tarefa

**POST** `/tasks`

**Body:**

```json
{
  "title": "Estudar Golang",
  "description": "Revisar conceitos de goroutines",
  "priority": "high",
  "due_date": "2026-02-10"
}
```

### Listar tarefas

**GET** `/tasks`

**Filtros opcionais:**
- `/tasks?status=pending`
- `/tasks?priority=high`
- `/tasks?status=pending&priority=high`

### Buscar tarefa por ID

**GET** `/tasks/{id}`

### Atualizar tarefa

**PUT** `/tasks/{id}`

**Exemplo de body:**

```json
{
  "title": "Estudar Golang - Atualizado",
  "status": "in_progress"
}
```

> **Obs:** Se a tarefa estiver com status = `completed`, a API retorna erro e não permite edição.

### Deletar tarefa

**DELETE** `/tasks/{id}`

## Estrutura do projeto

A estrutura foi organizada separando responsabilidades em camadas:

``text
cmd/
  main.go          -> ponto de entrada da aplicação

handlers/
  -> recebe as requisições HTTP

services/
  -> regras de negócio e validações

repositories/
  -> comunicação com o banco MongoDB

models/
  -> estruturas de dados

database/
  -> conexão com o MongoDB

routes/
  -> definição das rotas da API
``

## Objetivo do projeto

Esse projeto foi desenvolvido como exercício prático para estudar APIs REST com Go, incluindo:

- Organização de projeto
- Boas práticas de backend
- Separação de camadas
- Validações de negócio
- Integração com banco de dados