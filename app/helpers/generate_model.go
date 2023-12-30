package helpers

import (
	"fmt"
	"os"

)


func GenerateModel(fileName string, filePath string) error {
	// Gabungkan path ke folder models di dalam folder app
	fullPath := "./app/" + filePath

	// Cek apakah folder models sudah ada
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Folder '%s' tidak ditemukan", fullPath)
	}

	// Cek apakah file sudah ada di dalam folder tersebut
	targetFile := fullPath + "/" + fileName + ".go"
	_, err = os.Stat(targetFile)
	if err == nil {
		return fmt.Errorf("File '%s' sudah ada", fileName+".go")
	}

	modelName := fileName
	fileContent := `package models

type ` + modelName + ` struct {
	// Define model fields here
	Id        int       ` + "`json:\"id\" form:\"id\" gorm:\"primary_key\"`" + `
}`

	err = os.WriteFile(targetFile, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("%s file created successfully!\n", fileName+".go")
	return nil
}