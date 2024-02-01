package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	fileName := filepath.Base(filePth)
	fileExt := filepath.Ext(filePth)

	// Используем функцию strings.TrimPrefix для удаления точки из расширения
	fileExt = strings.TrimPrefix(fileExt, ".")

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
