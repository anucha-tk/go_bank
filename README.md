# GOBANK APP

<!--toc:start-->

- [GOBANK APP](#gobank-app)
  - [Tools and Technicals](#tools-and-technicals)
  - [Learn](#learn) - [Deadlock problem](#deadlock-problem)
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

### Unit testing

unit testing with gomock separate testing from real database, see more `api/account_test.go`
