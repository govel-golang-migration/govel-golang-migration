package govel_migration

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func Make(fileName string, migrationPath string) {
	absMigrationPath := getMigrateFolder(migrationPath)

	if err := os.MkdirAll(absMigrationPath, os.ModePerm); err != nil {
		panic(err)
	}

	location, _ := time.LoadLocation("Asia/Taipei")
	prefix := time.Now().In(location).Format("2006_01_02_150405")

	filePath := path.Join(absMigrationPath, prefix+"_"+fileName+".go")
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("migration file generate fail: " + filePath)
		return
	}

	content := loadStub(path.Join(getPackageDir(), "migration.stub"))
	content = strings.Replace(content, "{UpFunctionName}", "Up"+toCamelCase(fileName), -1)
	content = strings.Replace(content, "{DownFunctionName}", "Down"+toCamelCase(fileName), -1)

	file.WriteString(content)

	fmt.Println(successMessage(filePath + " created"))
}

func loadStub(fileName string) string {
	content, _ := os.ReadFile(fileName)

	return string(content)
}

func getPackageDir() string {
	_, file, _, ok := runtime.Caller(0)

	if !ok {
		panic("failed to get file path")
	}

	return path.Dir(file)
}
