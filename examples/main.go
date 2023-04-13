package main

import govel_migration "github.com/govel-golang-migration/govel-golang-migration"

func main() {
	mysqlDsn := "test:test@tcp(mysql:3306)/migration?charset=utf8mb4&parseTime=True&loc=Local"
	path := "./examples"
	govel_migration.Install(mysqlDsn)
	// govel_migration.Status(mysqlDsn, path)
	// govel_migration.Make("create_users_table", path)
	// govel_migration.Migrate(mysqlDsn, path, true)
	// govel_migration.Rollback(1, mysqlDsn, path, true)
}
