package config

import (
	"os"
	"path/filepath"
	"time"
)

func CreateLogFile() *os.File {
	folder := "log"
	os.Mkdir(folder, os.ModePerm)

	fileName := time.Now().Format("2006-01-02") + ".log"
	path := filepath.Join(folder, fileName)

	file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)

	return file
}
