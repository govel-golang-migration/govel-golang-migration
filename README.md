## Develop

`docker-compose build`

`docker-compose up -d`

## Features functions

* Install table
  * create table for migrations record
* Status
  * check status of migrations
* Make
  * create new migration file
* Migrate
  * run migrations
* Rollback
  * rollback migrations 

## Usage

see examples in `examples` folder

```bash
go run examples/install.go
go run examples/status.go
go run examples/migrate.go
go run examples/rollback.go
```
