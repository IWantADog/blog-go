package src

import (
	"os"
	"path/filepath"
	"testing"
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
	group *Group
	web   *Web
}

func TestLoadYaml(t *testing.T) {
	currentFile, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	testYaml := filepath.Dir(currentFile) + "testdata/a.yaml"

	var tf TestConf
	LoadYaml(testYaml, &tf)

	if tf.group.ID != 1 {
		t.Fail()
	}
}
