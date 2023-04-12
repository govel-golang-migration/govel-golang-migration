package govel_migration

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"path"
	"path/filepath"
)

func Status(mysqlDsn string) {
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if (!db.Migrator().HasTable(&Migration{})) {
		panic("Migration table does not exist, please using install command")
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	migratePath := path.Join(cwd, "migrations")
	if _, err := os.Stat(migratePath); os.IsNotExist(err) {
		fmt.Println("migration folder does not exist: " + migratePath)
		return
	}

	migrateNameHash, _ := buildMigrateNameHash(db)
	err = filepath.Walk(migratePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".go" {
			_, ok := migrateNameHash[info.Name()]
			if ok {
				fmt.Println(successMessage("Y ") + info.Name())
			} else {
				fmt.Println(errorMessage("N ") + info.Name())
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func infoMessage(message string) string {
	return "\033[34m" + message + "\033[0m"
}

func errorMessage(message string) string {
	return "\033[31m" + message + "\033[0m"
}

func successMessage(message string) string {
	return "\033[32m" + message + "\033[0m"
}

func buildMigrateNameHash(db *gorm.DB) (map[string]int, error) {
	var results []Migration
	var migrateNameHash = make(map[string]int)
	err := db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		migrateNameHash[result.Name] = result.Batch
	}
	return migrateNameHash, nil
}

func getMigrateFolder(migrationPath string) string {
	return path.Join(migrationPath, "migrations")
}

