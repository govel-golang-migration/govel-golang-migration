package main

import govel_migration "github.com/govel-golang-migration/govel-golang-migration"

func main() {
	mysqlDsn := "test:test@tcp(mysql:3306)/migration?charset=utf8mb4&parseTime=True&loc=Local"
	govel_migration.Install(mysqlDsn)
}
