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
    - [Decisão técnica](#decisão-técnica)
  - [Aula 4 — CRUD de ônibus em memória](#aula-4--crud-de-ônibus-em-memória)
    - [Decisão técnica](#decisão-técnica)
  - [Aula 5 — Persistência com Oracle](#aula-5--persistência-com-oracle)
    - [Decisão técnica](#decisão-técnica)
  - [Aula 6 — Repository](#aula-6--repository)
    - [Decisão técnica](#decisão-técnica)
  - [Aula 7 — Validação e tratamento de erros](#aula-7--validação-e-tratamento-de-erros)
    - [Decisão técnica](#decisão-técnica)
  - [Aula 8 — Testes da API](#aula-8--testes-da-api)
    - [Decisão técnica](#decisão-técnica)
- [Próximas aulas](#próximas-aulas)
  - [Aula 9 — Docker](#aula-9--docker)
  - [Aula 10 — Frontend Vue.js](#aula-10--frontend-vuejs)
  - [Aula 11 — CRUD no Vue](#aula-11--crud-no-vue)
  - [Aula 12 — Fechamento e documentação](#aula-12--fechamento-e-documentação)
- [Evolução](#evolução)
- [Deploy e publicação](#deploy-e-publicação)

## Tecnologias

### Go

Utilizado no backend da aplicação.

Go foi escolhido para este projeto porque permite construir uma API HTTP com uma base enxuta e próxima dos fundamentos da linguagem, sem exigir um framework para começar.

A evolução do backend será feita de forma incremental, começando pela biblioteca padrão e adicionando outras dependências somente quando houver uma necessidade real.

### Oracle Instant Client

Utilizado para permitir que a aplicação Go se comunique com o Oracle Database por meio do driver `godror`.

O Oracle Instant Client fornece as bibliotecas nativas necessárias para a conexão da aplicação com o banco de dados.

Nesta etapa foi utilizado o Oracle Instant Client 23.26.

### godror

Driver utilizado para integração entre Go e Oracle Database.

O `godror` é utilizado em conjunto com o pacote `database/sql`, permitindo que o backend execute operações SQL no Oracle utilizando a API padrão de acesso a bancos de dados do Go.

Versão utilizada:

```text
github.com/godror/godror v0.51.4
```

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

A aplicação Go utiliza o pacote `database/sql` em conjunto com o driver `godror` e o Oracle Instant Client para realizar a comunicação com o banco.

### Docker

Será utilizado para facilitar a criação e reprodução do ambiente de desenvolvimento.

O Docker também permitirá manter dependências de infraestrutura isoladas da máquina local e preparar o projeto para ambientes de publicação.

A utilização do Docker será abordada na Aula 9.

### Git e GitHub

Git será utilizado para controle de versão e o GitHub será utilizado para hospedagem do código-fonte.

O repositório público também faz parte do objetivo de portfólio do projeto, permitindo que recrutadores e outros desenvolvedores tenham acesso ao código e à evolução da aplicação.

### Deploy

O projeto será desenvolvido considerando sua publicação em ambiente de produção.

A solução de hospedagem será definida durante a etapa de deploy, priorizando alternativas gratuitas ou com camada gratuita suficiente para manter a aplicação disponível como projeto de portfólio.

## Estrutura do projeto

```text
## Estrutura do projeto

```text
bus-track-go/
├── backend/
│   ├── database/
│   │   └── oracle.go
│   ├── domain/
│   │   ├── bus.go
│   │   └── bus_test.go
│   ├── repository/
│   │   └── bus_repository.go
│   ├── go.mod
│   ├── go.sum
│   ├── main_test.go
│   └── main.go
├── frontend/
└── README.md
```

O diretório `backend` concentra a API desenvolvida em Go.<br>

O diretório `domain` contém as entidades e testes relacionados ao domínio da aplicação.

O diretório `repository` concentra a persistência dos dados e a comunicação entre a aplicação e o banco de dados.

O diretório `frontend` será utilizado posteriormente para a aplicação Vue.js.

## Banco de dados

O banco de dados principal do projeto é o **Oracle**.

A aplicação utiliza o pacote padrão `database/sql` do Go em conjunto com o driver `godror` para realizar a comunicação com o Oracle.

O ambiente também utiliza o **Oracle Instant Client**, responsável pelas bibliotecas nativas necessárias para a conexão.

A conexão atualmente utiliza:

```text
Oracle Database
    ↑
Oracle Instant Client
    ↑
godror
    ↑
database/sql
    ↑
Go API
```

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

**Objetivo**: criar a primeira API HTTP do BusTrack Go utilizando apenas a biblioteca padrão do Go.

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

### Aula 4 — CRUD de ônibus em memória

**Objetivo**: implementar o CRUD de ônibus utilizando armazenamento temporário em memória, antes da introdução do Oracle.

Foi realizado:

- [x] criação do armazenamento de ônibus em memória;
- [x] implementação do endpoint `GET /api/buses`;
- [x] implementação do endpoint `GET /api/buses/{id}`;
- [x] implementação do endpoint `POST /api/buses`;
- [x] implementação do endpoint `PUT /api/buses/{id}`;
- [x] implementação do endpoint `DELETE /api/buses/{id}`;
- [x] geração do próximo ID para novos ônibus em memória;
- [x] tratamento de ID inválido com HTTP 400;
- [x] tratamento de ônibus inexistente com HTTP 404;
- [x] tratamento de métodos HTTP não permitidos com HTTP 405;
- [x] retorno HTTP 201 Created para criação;
- [x] retorno HTTP 204 No Content para exclusão;
- [x] validação manual dos endpoints utilizando `curl`;
- [x] formatação do código com `gofmt`;
- [x] validação do projeto com `go test ./...`.

Endpoints implementados:

```text
GET    /api/buses
GET    /api/buses/{id}
POST   /api/buses
PUT    /api/buses/{id}
DELETE /api/buses/{id}
```

O armazenamento utilizado nesta etapa é um slice em memória:

```text
API HTTP
   ↓
[]Bus
```

Os dados são perdidos quando a aplicação é reiniciada. A persistência será implementada posteriormente com Oracle.

Comportamentos validados:

```text
GET    /api/buses        → 200 OK
GET    /api/buses/1      → 200 OK
GET    /api/buses/999    → 404 Not Found
GET    /api/buses/abc    → 400 Bad Request

POST   /api/buses        → 201 Created

PUT    /api/buses/3      → 200 OK
PUT    /api/buses/999    → 404 Not Found

DELETE /api/buses/3      → 204 No Content
DELETE /api/buses/999    → 404 Not Found
```

#### Decisão técnica

O CRUD foi implementado diretamente na camada HTTP e utiliza memória como armazenamento temporário.

Nesta etapa não foram criadas camadas de Service ou Repository, nem interfaces para abstrações futuras.

A decisão mantém a implementação simples e permite compreender o funcionamento do CRUD, dos métodos HTTP, dos status de resposta e da manipulação de dados em Go antes da introdução da persistência com Oracle.

A persistência será implementada posteriormente com Oracle.

### Aula 5 — Persistência com Oracle

**Objetivo**: substituir o armazenamento temporário em memória por persistência real utilizando Oracle Database.

Foi realizado:

- [x] configuração do Oracle Database utilizado pelo projeto;
- [x] instalação do Oracle Instant Client;
- [x] configuração das bibliotecas do Oracle no Linux;
- [x] validação das bibliotecas nativas necessárias para o Oracle;
- [x] instalação do driver `github.com/godror/godror`;
- [x] integração do `godror` com `database/sql`;
- [x] criação da conexão com o Oracle;
- [x] criação do package `database`;
- [x] criação da função `NewOracleConnection`;
- [x] validação da conexão Go → Oracle;
- [x] execução de `INSERT`;
- [x] execução de `SELECT`;
- [x] execução de `UPDATE`;
- [x] execução de `DELETE`;
- [x] substituição do armazenamento em memória pela persistência no Oracle;
- [x] validação da aplicação utilizando `go test ./...`;
- [x] validação da compilação utilizando `go build ./...`.

A conexão foi centralizada em:

```text
backend/database/oracle.go
```

A aplicação passou a utilizar o Oracle como fonte persistente dos dados.

Fluxo:
```text
Go API
   ↓
database/sql
   ↓
godror
   ↓
Oracle Instant Client
   ↓
Oracle Database
```

A partir desta etapa, os dados dos ônibus deixam de ser perdidos quando a aplicação é reiniciada.

#### Decisão técnica

Foi utilizado `database/sql` em conjunto com o driver `godror`, evitando acoplamento da aplicação diretamente a uma API específica de acesso ao banco.

O Oracle Instant Client foi utilizado como camada nativa necessária para a comunicação com o Oracle no ambiente Linux.

### Aula 6 — Repository

**Objetivo**: separar a responsabilidade de persistência da camada HTTP.

Foi realizado:

- [x] criação do package `repository`;
- [x] criação do `BusRepository`;
- [x] criação da função `NewBusRepository`;
- [x] implementação de `FindAll`;
- [x] implementação de `FindByID`;
- [x] implementação de `Create`;
- [x] implementação de `Update`;
- [x] implementação de `Delete`;
- [x] utilização de SQL para persistência no Oracle;
- [x] tratamento de registros inexistentes;
- [x] tratamento de erros de banco de dados;
- [x] integração do Repository com a API;
- [x] validação do CRUD completo utilizando Oracle.

Estrutura atual:

```text
HTTP
 ↓
Repository
 ↓
Oracle
```

O Repository foi definido em:

```text
backend/repository/bus_repository.go
```

#### Decisão técnica

O Repository foi introduzido somente quando surgiu uma necessidade real de separar a persistência da camada HTTP.

Não foram criadas interfaces apenas para aumentar a complexidade da arquitetura.

A responsabilidade do Repository é executar operações de persistência e retornar os dados ou erros para a camada superior.

### Aula 7 — Validação e tratamento de erros

**Objetivo**: tornar a API mais segura e previsível no tratamento de entradas inválidas, registros inexistentes e erros de operação.

Foi realizado:

- [x] validação de campos obrigatórios;
- [x] validação de capacidade do ônibus;
- [x] validação de ID recebido pela URL;
- [x] tratamento de ID inválido com HTTP 400;
- [x] tratamento de ônibus inexistente com HTTP 404;
- [x] tratamento de métodos HTTP não permitidos;
- [x] tratamento de erros de persistência;
- [x] retorno HTTP 201 Created para criação;
- [x] retorno HTTP 204 No Content para exclusão;
- [x] validação manual dos endpoints utilizando `curl`;
- [x] validação de criação de registros no Oracle;
- [x] validação de atualização de registros no Oracle;
- [x] validação de exclusão de registros no Oracle.

Comportamentos validados:

```text
POST /api/buses com dados inválidos → 400 Bad Request

GET /api/buses/abc → 400 Bad Request

GET /api/buses/999999 → 404 Not Found

POST /api/buses → 201 Created

PUT /api/buses/{id} → 200 OK

DELETE /api/buses/{id} → 204 No Content

GET /api/buses/{id} após exclusão → 404 Not Found
```

Exemplo de validação:

```text
DELETE /api/buses/42
→ 204 No Content

GET /api/buses/42
→ 404 Not Found
```

A API agora diferencia erros de entrada, recursos inexistentes e operações realizadas com sucesso.

#### Decisão técnica

O tratamento de erros foi mantido simples e explícito, utilizando os códigos HTTP adequados para cada situação.

O objetivo nesta etapa não é criar um framework próprio de erros, mas estabelecer um comportamento consistente para a API antes da introdução de testes automatizados mais abrangentes.

### Aula 8 — Testes da API

**Objetivo**: automatizar os testes da API para validar os handlers, regras de validação e diferentes cenários de sucesso e erro sem depender do Oracle Database durante a execução dos testes.

Foi realizado:

- [x] criação do arquivo `backend/main_test.go`;
- [x] criação de testes unitários;
- [x] criação de testes para os handlers HTTP;
- [x] utilização do pacote `net/http/httptest`;
- [x] criação de um mock do `BusRepository` por meio de interface;
- [x] testes do endpoint `GET /health`;
- [x] testes de validação de dados;
- [x] testes de parsing do ID;
- [x] testes do endpoint `GET /api/buses`;
- [x] testes do endpoint `POST /api/buses`;
- [x] testes do endpoint `GET /api/buses/{id}`;
- [x] testes do endpoint `PUT /api/buses/{id}`;
- [x] testes do endpoint `DELETE /api/buses/{id}`;
- [x] testes de respostas de sucesso;
- [x] testes de entradas inválidas;
- [x] testes de registros inexistentes;
- [x] testes de erros do repository;
- [x] testes de métodos HTTP não permitidos;
- [x] validação dos testes com `go test ./...`;
- [x] execução detalhada dos testes com `go test -v ./...`;
- [x] formatação dos testes com `gofmt`.

Arquivo criado:

```text
backend/main_test.go
```

Estrutura atual relacionada aos testes:

```text
backend/
├── database/
│   └── oracle.go
├── domain/
│   ├── bus.go
│   └── bus_test.go
├── repository/
│   └── bus_repository.go
├── go.mod
├── go.sum
├── main.go
└── main_test.go
```

Os testes dos handlers utilizam `httptest`, permitindo simular requisições HTTP e verificar os códigos de status e as respostas da API sem iniciar um servidor real.

A interface `BusRepository` foi adicionada à camada HTTP para permitir que os handlers recebam uma implementação real do Repository ou um mock durante os testes.

Resultado:

```text
go test ./...

ok  	github.com/lgomesroc/bus-track-go
ok  	github.com/lgomesroc/bus-track-go/domain
```

Execução detalhada:

```text
go test -v ./...

PASS
```

Foram cobertos cenários de sucesso e erro, incluindo:

```text
GET    /health              → 200 OK
POST   /health              → 405 Method Not Allowed

GET    /api/buses           → 200 OK
GET    /api/buses           → 500 Internal Server Error
POST   /api/buses           → 201 Created
POST   /api/buses           → 400 Bad Request

GET    /api/buses/{id}      → 200 OK
GET    /api/buses/{id}      → 400 Bad Request
GET    /api/buses/{id}      → 404 Not Found
GET    /api/buses/{id}      → 500 Internal Server Error

PUT    /api/buses/{id}      → 200 OK
PUT    /api/buses/{id}      → 400 Bad Request
PUT    /api/buses/{id}      → 404 Not Found

DELETE /api/buses/{id}      → 204 No Content
DELETE /api/buses/{id}      → 404 Not Found
DELETE /api/buses/{id}      → 500 Internal Server Error
```

#### Decisão técnica

Os testes foram introduzidos utilizando recursos simples da biblioteca padrão do Go, principalmente `testing` e `net/http/httptest`.

Para testar os handlers sem depender diretamente do Oracle, foi criada uma interface `BusRepository`. Dessa forma, os testes podem utilizar uma implementação simulada do repository.

Não foi criada uma infraestrutura externa de testes, banco de dados específico para testes ou framework de mocking.

A abordagem mantém os testes simples e adequados ao tamanho e objetivo do BusTrack Go.

## Próximas aulas

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

## Evolução
```text
Aula 1
  ↓
Inicialização Go

Aula 2
  ↓
HTTP

Aula 3
  ↓
Domínio

Aula 4
  ↓
CRUD em memória

Aula 5
  ↓
Oracle + godror + Instant Client

Aula 6
  ↓
Repository

Aula 7
  ↓
Validação + erros

Aula 8
  ↓
Testes da API

Aula 9
  ↓
Docker
```

## Deploy e publicação

O projeto será desenvolvido desde o início considerando a publicação da aplicação em ambiente de produção.

Ao final do desenvolvimento, o BusTrack Go terá:

- [ ] backend publicado;
- [ ] frontend publicado;
- [ ] aplicação acessível por URL pública;
- [ ] configuração de produção;
- [ ] documentação do processo de deploy.

O objetivo é disponibilizar tanto o repositório no GitHub quanto uma versão funcional da aplicação para utilização como projeto de portfólio.