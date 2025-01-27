# GOBANK APP

<!--toc:start-->

- [GOBANK APP](#gobank-app)
  - [Tools and Technical](#tools-and-technical)
  - [Learn](#learn)
  - [Migration](#migration)
    - [Deadlock problem](#deadlock-problem)
    - [Register custom validation params](#register-custom-validation-params)
    - [Unit testing](#unit-testing)
    - [Handle DB Errors](#handle-db-errors)
    - [JWT vs Paseto Token](#jwt-vs-paseto-token)
    - [Dockerfile](#dockerfile)
  - [Resource](#resource)

## Tools and Technical

- HTTP Framework
  - Gin
- Database
  - PostgreSQL
  - [migrate](https://github.com/golang-migrate/migrate)
  - Sqlc
  - Transaction
- Testing
  - [Testify](https://github.com/stretchr/testify)
  - Faker
  - Gomock(Unit-test)
- Infrastructure
  - Viper(configuration)
  - Docker
  - Makefile
  - Github Actions (CI)
- ETC
  - JWT and [Paseto](https://github.com/paragonie/paseto)
    - Authentication(Access Token and Refresh Token)
  - Auth Middleware

## Learn

## Migration

- Create a new db migrate

```bash
migrate create -ext sql -dir db/migration -seq <migration_name>
```

### Deadlock problem

avoid deadlock with order see more `/db/sqlc/store.go`

### Register custom validation params

create constant of currency on my app at `/util/currency.go` then custom at `/api/validCurrency.go`
and register at `/api/server.go`

> using `json:"currency" binding:"required,currency"`

### Unit testing

unit testing with gomock separate testing from real database, see more `api/account_test.go`

### Handle DB Errors

> use err.(\*pq.Error) at `/api/account.go`

### JWT vs Paseto Token

learn build jwt token and paseto (modern) token at `/token`

### Dockerfile

learn build stage for reduce file size see more `Dockerfile` and pre-load [migrate](https://github.com/golang-migrate/migrate) and ensure run `start.sh` after postgres container is ready with `wait-for.sh`

## Resource

[backend-master-class-golang-postgresql-kubernetes](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/)
