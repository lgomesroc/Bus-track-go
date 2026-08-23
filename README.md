# BusTrack Go

Backend de uma aplicação de gerenciamento de ônibus, desenvolvido em Go.

O projeto foi criado com foco em aprendizado prático de desenvolvimento backend, mantendo uma arquitetura simples e evoluindo a aplicação de forma incremental.

## Índice

- [Estrutura do projeto](#estrutura-do-projeto)
- [Banco de dados](#banco-de-dados)
- [Estrutura de pastas](#estrutura-de-pastas)
- [Aulas](#aulas)
  - [Aula 1 — Inicialização do BusTrack Go](#aula-1--inicialização-do-bustrack-go)
  - [Aula 2 — Primeira API HTTP em Go](#aula-2--primeira-api-http-em-go)
    - [Decisão técnica](#decisão-técnica)
- [Próximas aulas](#próximas-aulas)
  - [Aula 3 — Modelando o domínio](#aula-3--modelando-o-domínio)
  - [Aula 4 — CRUD de ônibus em memória](#aula-4--crud-de-ônibus-em-memória)
  - [Aula 5 — Persistência com Oracle](#aula-5--persistência-com-oracle)
  - [Aula 6 — Repository](#aula-6--repository)
  - [Aula 7 — Validação e tratamento de erros](#aula-7--validação-e-tratamento-de-erros)
  - [Aula 8 — Testes da API](#aula-8--testes-da-api)
  - [Aula 9 — Docker](#aula-9--docker)
  - [Aula 10 — Frontend Vue.js](#aula-10--frontend-vuejs)
  - [Aula 11 — CRUD no Vue](#aula-11--crud-no-vue)
  - [Aula 12 — Fechamento e documentação](#aula-12--fechamento-e-documentação)

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

### Aula 2 — Primeira API HTTP em Go

Objetivo: criar a primeira API HTTP do BusTrack Go utilizando apenas a biblioteca padrão do Go.

Foi realizado:

- [x] utilização do pacote net/http;
- [x] criação de um servidor HTTP;
- [x] criação da rota /health;
- [x] criação de um handler HTTP;
- [x] utilização do método GET;
- [x] validação do método HTTP recebido;
- [x] retorno do status 200 OK;
- [x] retorno do status 405 Method Not Allowed para métodos não permitidos;
- [x] utilização do status 404 Not Found para rotas inexistentes;
- [x] definição do Content-Type como application/json;
- [x] utilização do pacote encoding/json;
- [x] serialização da resposta para JSON;
- [x] criação do endpoint GET /health.

Endpoint:

```text
GET /health
```

Resposta:

```text
{
  "status": "ok"
}
```

Comportamentos validados:

```text
GET  /health  → 200 OK
POST /health  → 405 Method Not Allowed
GET  /buses   → 404 Not Found
```

#### Decisão técnica

Nesta etapa não foi utilizado nenhum framework HTTP.

A implementação utiliza somente recursos da biblioteca padrão do Go, principalmente net/http e encoding/json.

A decisão foi intencional: antes de adicionar abstrações ou frameworks, o projeto deve estabelecer uma compreensão dos fundamentos de HTTP, incluindo rotas, métodos, status, headers e respostas JSON.

O objetivo do BusTrack Go é evoluir de forma incremental, adicionando complexidade somente quando ela resolver um problema real do projeto.

## Próximas aulas

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

- [ ] structs;
- [ ] tipos;
- [ ] JSON;
- [ ] organização do código;
- [ ] separação básica entre domínio e HTTP.

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

- [ ] Oracle;
- [ ] Docker;
- [ ] conexão com banco;
- [ ] `database/sql`;
- [ ] driver Oracle;
- [ ] configuração da conexão;
- [ ] primeira tabela;
- [ ] `INSERT`;
- [ ] `SELECT`;
- [ ] `UPDATE`;
- [ ] `DELETE`.

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

- [ ] validação de entrada;
- [ ] HTTP 400;
- [ ] HTTP 404;
- [ ] HTTP 500;
- [ ] erros do banco;
- [ ] respostas JSON padronizadas;
- [ ] tratamento de situações inválidas.

Exemplo:

```json
{
  "error": "bus not found"
}
```

### Aula 8 — Testes da API

Entramos nos testes.

- [ ] testes unitários;
- [ ] testes dos handlers;
- [ ] httptest;
- [ ] casos de sucesso;
- [ ] casos de erro;
- [ ] testes do service.

> Sem criar uma infraestrutura desnecessariamente complexa de testes.

### Aula 9 — Docker

Agora vamos organizar o ambiente.

- [ ] Docker;
- [ ] container Oracle;
- [ ] variáveis de ambiente;
- [ ] conexão da aplicação com o banco;
- [ ] configuração para desenvolvimento.

A ideia é chegar a algo próximo de:

```text
Docker
 └── Oracle

Go API
 └── conecta no Oracle
 ```

### Aula 10 — Frontend Vue.js

Agora entra o frontend.

- [ ] criação do projeto Vue;
- [ ] JavaScript;
- [ ] estrutura básica;
- [ ] componentes;
- [ ] páginas;
- [ ] consumo da API.

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

- [ ] listar ônibus;
- [ ] cadastrar;
- [ ] editar;
- [ ] excluir;
- [ ] formulário;
- [ ] tratamento de erros;
- [ ] loading;
- [ ] comunicação com API Go.

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

- [ ] revisar estrutura;
- [ ] revisar código;
- [ ] testar aplicação completa;
- [ ] revisar Docker;
- [ ] atualizar README;
- [ ] explicar arquitetura;
- [ ] registrar decisões técnicas;
- [ ] criar versão final do projeto;
- [ ] Git/GitHub.