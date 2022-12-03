package src

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYaml(filePath string, aimStruct interface{}) {
	rawContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("load conf error: %s read failed", filePath)
	}
	err = yaml.Unmarshal(rawContent, aimStruct)
	if err != nil {
		log.Fatalf("load conf error: %s yaml load failed", filePath)
	}
}
