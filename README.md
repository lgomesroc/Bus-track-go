# BusTrack Go

Backend de uma aplicação de gerenciamento de ônibus, desenvolvido em Go.

O projeto foi criado com foco em aprendizado prático de desenvolvimento backend, mantendo uma arquitetura simples e evoluindo a aplicação de forma incremental.

## Índice

- [Tecnologias](#tecnologias)
- [Estrutura do projeto](#estrutura-do-projeto)
- [Banco de dados](#banco-de-dados)
- [Aulas](#aulas)
  - [Aula 1 — Inicialização do BusTrack Go](#aula-1--inicialização-do-bustrack-go)
  - [Aula 2 — Primeira API HTTP em Go](#aula-2--primeira-api-http-em-go)
    - [Decisão técnica](#decisão-técnica)
  - [Aula 3 — Modelando o domínio](#aula-3--modelando-o-domínio)
- [Próximas aulas](#próximas-aulas)
  - [Aula 4 — CRUD de ônibus em memória](#aula-4--crud-de-ônibus-em-memória)
  - [Aula 5 — Persistência com Oracle](#aula-5--persistência-com-oracle)
  - [Aula 6 — Repository](#aula-6--repository)
  - [Aula 7 — Validação e tratamento de erros](#aula-7--validação-e-tratamento-de-erros)
  - [Aula 8 — Testes da API](#aula-8--testes-da-api)
  - [Aula 9 — Docker](#aula-9--docker)
  - [Aula 10 — Frontend Vue.js](#aula-10--frontend-vuejs)
  - [Aula 11 — CRUD no Vue](#aula-11--crud-no-vue)
  - [Aula 12 — Fechamento e documentação](#aula-12--fechamento-e-documentação)

## Tecnologias

### Go

Utilizado no backend da aplicação.

Go foi escolhido para este projeto porque permite construir uma API HTTP com uma base enxuta e próxima dos fundamentos da linguagem, sem exigir um framework para começar.

A evolução do backend será feita de forma incremental, começando pela biblioteca padrão e adicionando outras dependências somente quando houver uma necessidade real.

### Vue.js

Será utilizado no frontend da aplicação.

Vue.js foi escolhido por permitir construir uma interface web moderna mantendo uma curva de aprendizado adequada ao objetivo do projeto.

A aplicação frontend será desenvolvida utilizando JavaScript.

### JavaScript

Será utilizado como linguagem do frontend.

A escolha mantém o frontend simples e evita adicionar TypeScript neste projeto, permitindo concentrar o aprendizado do BusTrack Go no backend em Go e na integração entre frontend e API.

### Oracle

Será utilizado como banco de dados principal.

Oracle foi escolhido para permitir trabalhar com um banco relacional utilizado em ambientes corporativos, incluindo conexão, SQL, transações e persistência de dados.

### Docker

Será utilizado para facilitar a criação e reprodução do ambiente de desenvolvimento.

O Docker também permitirá manter dependências de infraestrutura isoladas da máquina local e preparar o projeto para ambientes de publicação.

### Git e GitHub

Git será utilizado para controle de versão e o GitHub será utilizado para hospedagem do código-fonte.

O repositório público também faz parte do objetivo de portfólio do projeto, permitindo que recrutadores e outros desenvolvedores tenham acesso ao código e à evolução da aplicação.

### Deploy

O projeto será desenvolvido considerando sua publicação em ambiente de produção.

A solução de hospedagem será definida durante a etapa de deploy, priorizando alternativas gratuitas ou com camada gratuita suficiente para manter a aplicação disponível como projeto de portfólio.

## Estrutura do projeto

```text
bus-track-go/
├── backend/
│   ├── domain/
│   │   ├── bus.go
│   │   └── bus_test.go
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

### Aula 3 — Modelando o domínio

**Objetivo**: definir o modelo de domínio `Bus` e estabelecer a primeira separação entre o domínio da aplicação e a camada HTTP.

Foi realizado:

- [x] criação do package `domain`;
- [x] criação da struct `Bus`;
- [x] definição dos campos `ID`, `Prefix`, `LicensePlate`, `Model`, `Capacity` e `Status`;
- [x] utilização dos tipos `int` e `string`;
- [x] utilização de struct tags para serialização JSON;
- [x] utilização do `encoding/json`;
- [x] teste de conversão de `Bus` para JSON com `json.Marshal`;
- [x] teste de conversão de JSON para `Bus` com `json.Unmarshal`;
- [x] criação de testes unitários para o domínio;
- [x] validação dos testes com `go test ./...`;
- [x] manutenção da separação básica entre domínio e HTTP.

Modelo:

```text
Bus
├── ID
├── Prefix
├── LicensePlate
├── Model
├── Capacity
└── Status
```

A struct foi definida em:

```text
backend/domain/bus.go
```

Os testes foram definidos em:

```text
backend/domain/bus_test.go
```

O domínio não possui dependência direta de `net/http`. A responsabilidade do Bus é representar uma entidade do negócio, enquanto a camada HTTP continuará responsável pela comunicação com a API.

Também foi validado que a criação do domínio não alterou o funcionamento da API existente:

```text
GET /health → 200 OK
```

Resposta:

```json
{
  "status": "ok"
}
```

#### Decisão técnica

O domínio foi separado da camada HTTP antes da implementação do CRUD.

Neste momento não foram criadas camadas de Service, Repository ou outras abstrações. A arquitetura continuará evoluindo conforme surgirem necessidades reais no projeto.

Essa abordagem mantém o BusTrack Go simples e permite compreender cada responsabilidade antes de adicionar novas camadas.

## Próximas aulas

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

## Deploy e publicação

O projeto será desenvolvido desde o início considerando a publicação da aplicação em ambiente de produção.

Ao final do desenvolvimento, o BusTrack Go terá:

- [ ] backend publicado;
- [ ] frontend publicado;
- [ ] aplicação acessível por URL pública;
- [ ] configuração de produção;
- [ ] documentação do processo de deploy.

O objetivo é disponibilizar tanto o repositório no GitHub quanto uma versão funcional da aplicação para utilização como projeto de portfólio.