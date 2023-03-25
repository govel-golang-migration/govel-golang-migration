package govel_migration

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Migration struct {
	gorm.Model
	Name  string
	Batch int
}

func Install(mysqlDsn string) {
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if (db.Migrator().HasTable(&Migration{})) {
		panic("Migration table already exists")
	}

	db.Migrator().CreateTable(&Migration{})
	fmt.Print("Migration table created")
}
