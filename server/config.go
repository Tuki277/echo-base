package server

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "echo-base"

func NewConfig() error {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		panic("Error loading .env file")
	}
	return nil
}
