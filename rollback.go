package govel_migration

import (
	"fmt"
	"os/exec"
	"path"
	"plugin"
	"regexp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Rollback(stage int, mysqlDsn string) {
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	folder := getMigrateFolder()

	cmd := exec.Command("go", "build", "-buildmode=plugin")
	cmd.Dir = folder

	err = cmd.Run()
	if err != nil {
		panic("build error")
	}

	soPath := path.Join(folder, "migrations.so")

	plug, err := plugin.Open(soPath)

	if err != nil {
		panic(err)
	}

	fileNames := getRollbackFileName(stage, db)

	for _, migration := range fileNames {
		println(infoMessage(fmt.Sprintf("rollback %s on batch %d", migration.Name, migration.Batch)))
		r := regexp.MustCompile(`^(\d{4}_\d{2}_\d{2}_\d{6})_(.+)\.go$`)
		
		match := r.FindStringSubmatch(migration.Name)

		runLib, err := plug.Lookup("Down" + toCamelCase(match[2]))

		if err != nil {
			panic(err)
		}

		runLib.(func())()
	}

	println(successMessage("rollback success"))

	return
}

func getRollbackFileName(stage int, db *gorm.DB) []Migration {
	var migrations []Migration
	err := db.Where("batch > ((select max(batch) from migrations) - ?)", stage).Order("id DESC").Find(&migrations).Error

	if err != nil {
		fmt.Println(errorMessage("Get rollback file name error"))
		panic(err)
	}

	return migrations
}
