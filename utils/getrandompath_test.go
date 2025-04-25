package utils

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGenFileName(t *testing.T) {
	filename := GenFileName("/a/b/c/d/f/e/g.xxx")
	t.Log(filename)
	if filename != "eg.xxx" {
		t.Error("GenFileName error")
	}
}
func TestReadOneFile(t *testing.T) {
	jsondata, _ := os.ReadFile(GetJsonPath())
	FileList := TrueFilelist{
		RootPath: "",
		Files:    []Myfile{},
	}
	json.Unmarshal(jsondata, &FileList)
	FileList.Files = []Myfile{}
	data, _ := json.Marshal(FileList)
	os.WriteFile(GetJsonPath(), data, 0644)
	err := ReadOneFile()
	if err != nil {
		t.Error(err)
	}

}
