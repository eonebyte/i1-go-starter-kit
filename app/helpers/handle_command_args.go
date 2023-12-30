package helpers

import (
	"fmt"
	"os"
)

func HandleCommandArgs() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go [namaFile] [tyOfFile]")
		return
	}

	fileName := os.Args[1]
	filePath := os.Args[2]
	fileType := os.Args[3]

	switch fileType {
	case "controllers":
		err := GenerateController(fileName, filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Fungsi GenerateController dijalankan untuk file:", fileName)
	case "models":
		err := GenerateModel(fileName, filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Fungsi GenerateModel dijalankan untuk file:", fileName)
	default:
		fmt.Println("Usage: go run main.go [namaFile] [tyOfFile] for generate Controllers/Models")
	}
}
