package src

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func LoadYaml(filename string, structPoint interface{}) {
	root := getConfFolderPath()
	filePath := filepath.Join(root, filename)
	loadYaml(filePath, structPoint)
}

func getConfFolderPath() string {
	// yaml与main.go在同一层级
	currentFile, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentFile
}

func loadYaml(filePath string, point interface{}) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, point)
	if err != nil {
		log.Fatal(err)
	}
}
