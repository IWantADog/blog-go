package global

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

func LoadYaml(filename string, structPoint interface{}) {
	root := getConfFolderPath()
	filePath := filepath.Join(root, filename)
	loadYaml(filePath, structPoint)
}

func getConfFolderPath() string {
	// yaml与main.go在同一层级
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("callers error")
	}
	dirName := filepath.Clean(filepath.Join(currentFile, "../../.."))
	return dirName
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
