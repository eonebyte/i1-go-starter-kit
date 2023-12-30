package helpers

import (
	"fmt"
	"os"

)

func GenerateController(fileName string, filePath string) error {
	// Gabungkan path ke folder controllers di dalam folder app
	fullPath := "./app/" + filePath

	// Cek apakah folder controllers sudah ada
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
	

	controllerName := fileName
	fileContent := `package controllers

	import "github.com/gofiber/fiber/v2/middleware/session"


type ` + controllerName + ` struct {
	baseUrl string
	session *session.Store
}

func New` + controllerName + `(baseUrl string, session *session.Store) *` + controllerName + ` {
	return &` + controllerName + `{
		baseUrl: baseUrl,
		session: session,
	}
}`

   err = os.WriteFile(fullPath+"/"+fileName+".go", []byte(fileContent), 0644)
   if err != nil {
	return err
   }
	

	fmt.Printf("%s file created successfully!\n", fileName+".go")
	return nil
}
