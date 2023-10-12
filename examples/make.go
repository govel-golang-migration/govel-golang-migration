package main

import govel_migration "github.com/govel-golang-migration/govel-golang-migration"

func main() {
	path := "./examples"
	govel_migration.Make("create_users_table", path)
}
