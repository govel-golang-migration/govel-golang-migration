package govel_migration

import (
	"fmt"
	"strings"
	"os"
	"path"
	"time"
)

func Make(fileName string) {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	migrationPath := path.Join(cwd, "migrations")

	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		fmt.Println("migration folder does not exist: " + migrationPath)
		return
	}

	location, _ := time.LoadLocation("Asia/Taipei")
	prefix := time.Now().In(location).Format("2006_01_02_150405")

	filePath := path.Join(migrationPath, prefix + "_" + fileName + ".go")
	file, err := os.Create(filePath)
	
	if err != nil {
		fmt.Println("migration file generate fail: " + filePath)
		return
	}
	
	content := loadStub(path.Join(cwd, "migration.stub"))
	content = strings.Replace(content, "{UpFunctionName}", "Up" + toCamelCase(fileName), -1)
	content = strings.Replace(content, "{DownFunctionName}", "Down" + toCamelCase(fileName), -1)

	file.WriteString(content)

	fmt.Println(successMessage(filePath + " created"))
}

func loadStub(fileName string) string {
	content, _ := os.ReadFile(fileName)

	return string(content)
}
