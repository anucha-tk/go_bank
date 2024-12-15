# GOBANK APP

<!--toc:start-->

- [GOBANK APP](#gobank-app)
  - [Tools and Technicals](#tools-and-technicals)
  - [Learn](#learn) - [Deadlock problem](#deadlock-problem) - [Register custom validation params](#register-custom-validation-params) - [Unit testing](#unit-testing) - [Handle DB Errors](#handle-db-errors)
  <!--toc:end-->

## Tools and Technicals

- HTTP Framework
  - Gin
- Database
  - PostgreSQL
  - Golang-migrate
  - Sqlc
  - Transection
- Testing
  - Testify
  - Faker
  - Gomock(Unit-test)
- Infrastructer
  - Viper(configuration)
  - Docker
  - Makefile
  - Github Actions (CI)

## Learn

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
