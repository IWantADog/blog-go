package src

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"
)

type Group struct {
	Name string
	ID   int32
}

type Web struct {
	Host string
	Port int32
}

type TestConf struct {
	Group *Group
	Web   *Web
}

func TestLoadYaml(t *testing.T) {
	currentFile, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	testYaml := filepath.Dir(currentFile) + "/testdata/a.yaml"

	var tf TestConf
	p := &tf
	loadYaml(testYaml, p)

	fmt.Printf("t ---- %T", p)
	if tf.Group.ID != 1 {
		t.Fail()
	}
}

func TestYaml2(t *testing.T) {
	c, _ := os.Getwd()

	testYaml := filepath.Dir(c) + "/testdata/a.yaml"
	content, _ := os.ReadFile(testYaml)
	fmt.Printf("content --\n%v\n", string(content))

	tf := TestConf{}
	_ = yaml.Unmarshal(content, &tf)

	fmt.Printf("t-------\n%v\n", tf)

	if tf.Group.ID != 1 {
		log.Fatal("aaaaa")
	}
}
