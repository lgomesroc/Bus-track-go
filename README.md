# BusTrack Go

Backend de uma aplicação de gerenciamento de ônibus, desenvolvido em Go.

O projeto foi criado com foco em aprendizado prático de desenvolvimento backend, mantendo uma arquitetura simples e evoluindo a aplicação de forma incremental.

## Estrutura do projeto

```text
bus-track-go/
├── backend/
│   ├── go.mod
│   └── main.go
├── frontend/
└── README.md
```

O diretório `backend` concentra a API desenvolvida em Go.<br>

O diretório `frontend` será utilizado posteriormente para a aplicação Vue.js.

## Banco de dados

O banco de dados principal planejado para o projeto é o **Oracle**.

Caso ocorram dificuldades técnicas relevantes durante a integração, o **MySQL** poderá ser utilizado como alternativa.

## Estrutura de pastas

```text
bus-track-go/
├── backend/
│   ├── go.mod
│   └── main.go
├── frontend/
└── README.md
```

## Aulas

### Aula 1 — Inicialização do BusTrack Go

**Objetivo**: preparar a estrutura inicial do projeto e criar a primeira aplicação Go.

Foi realizado:

- [x] criação da estrutura inicial do projeto;
- [x] separação entre backend e frontend;
- [x] criação e validação do go.mod;
- [x] definição do módulo github.com/lgomesroc/bus-track-go;
- [x] criação do main.go;
- [x] utilização de package main;
- [x] criação da func main();
- [x] execução da aplicação com go run main.go;
- [x] validação da saída BusTrack Go;
- [x] inicialização do repositório Git;
- [x] definição da branch principal como main;
- [x] criação do repositório público no GitHub;
- [x] configuração do remote;
- [x] primeiro commit e envio para o GitHub.

Resultado:

```text
BusTrack Go
```

## Próximas aulas

### Aula 2 — Primeira API HTTP em Go

Aqui começamos realmente o backend.

- [x] `net/http`;
- [x] servidor HTTP;
- [x] rotas;
- [x] métodos HTTP;
- [x] `GET`;
- [x] status HTTP;
- [x] resposta JSON;
- [x] endpoint `/health`.

Exemplo conceitual:

```text
GET /health

{
    "status": "ok"
}
```

Nada de framework ainda.

### Aula 3 — Modelando o domínio

Vamos definir o que é um ônibus dentro do sistema.

```text
Bus
├── ID
├── Prefix
├── LicensePlate
├── Model
├── Capacity
└── Status
```

Aqui entram:

- [x] structs;
- [x] tipos;
- [x] JSON;
- [x] organização do código;
- [x] separação básica entre domínio e HTTP.

### Aula 4 — CRUD de ônibus em memória

Antes de envolver Oracle, vamos fazer o sistema funcionar.

Endpoints:

```text
GET    /api/buses
GET    /api/buses/{id}
POST   /api/buses
PUT    /api/buses/{id}
DELETE /api/buses/{id}
```

Usaremos memória temporariamente.

Isso permite aprender **API + Go** sem colocar banco, Docker e frontend simultaneamente na mesa.

### Aula 5 — Persistência com Oracle

Agora entra uma das partes novas para o projeto.

- [x] Oracle;
- [x] Docker;
- [x] conexão com banco;
- [x] `database/sql`;
- [x] driver Oracle;
- [x] configuração da conexão;
- [x] primeira tabela;
- [x] `INSERT`;
- [x] `SELECT`;
- [x] `UPDATE`;
- [x] `DELETE`.

Aqui o BusTrack deixa de ser apenas uma API em memória.

 ### Aula 6 — Repository

Agora fazemos uma organização simples para separar:

```text
HTTP
  ↓
Service
  ↓
Repository
  ↓
Oracle
```

> Não vamos criar interfaces e abstrações desnecessárias apenas para aumentar a complexidade da arquitetura.

> O objetivo é entender a responsabilidade de cada parte.

### Aula 7 — Validação e tratamento de erros

Vamos deixar a API minimamente profissional.

- [x] validação de entrada;
- [x] HTTP 400;
- [x] HTTP 404;
- [x] HTTP 500;
- [x] erros do banco;
- [x] respostas JSON padronizadas;
- [x] tratamento de situações inválidas.

Exemplo:

```json
{
  "error": "bus not found"
}
```

### Aula 8 — Testes da API

Entramos nos testes.

- [x] testes unitários;
- [x] testes dos handlers;
- [x] httptest;
- [x] casos de sucesso;
- [x] casos de erro;
- [x] testes do service.

> Sem criar uma infraestrutura desnecessariamente complexa de testes.

### Aula 9 — Docker

Agora vamos organizar o ambiente.

- [x] Docker;
- [x] container Oracle;
- [x] variáveis de ambiente;
- [x] conexão da aplicação com o banco;
- [x] configuração para desenvolvimento.

A ideia é chegar a algo próximo de:

```text
Docker
 └── Oracle

Go API
 └── conecta no Oracle
 ```

### Aula 10 — Frontend Vue.js

Agora entra o frontend.

- [x] criação do projeto Vue;
- [x] JavaScript;
- [x] estrutura básica;
- [x] componentes;
- [x] páginas;
- [x] consumo da API.

Primeira tela:

```text
BusTrack
────────────────────────

Ônibus cadastrados

001   Mercedes-Benz   Ativo
002   Volvo           Ativo
003   Scania          Manutenção
```

### Aula 11 — CRUD no Vue

Integração completa.

- [x] listar ônibus;
- [x] cadastrar;
- [x] editar;
- [x] excluir;
- [x] formulário;
- [x] tratamento de erros;
- [x] loading;
- [x] comunicação com API Go.

Nesse ponto teremos:

```text
Vue.js
   ↓ HTTP
Go API
   ↓
Oracle
```

### Aula 12 — Fechamento e documentação

Aqui fazemos o acabamento.

- [x] revisar estrutura;
- [x] revisar código;
- [x] testar aplicação completa;
- [x] revisar Docker;
- [x] atualizar README;
- [x] explicar arquitetura;
- [x] registrar decisões técnicas;
- [x] criar versão final do projeto;
- [x] Git/GitHub.