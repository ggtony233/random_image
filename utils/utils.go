package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"gitee.com/ggtony/folder-scan/filescan"
	"gitee.com/ggtony/folder-scan/task"
)

type RIConfig struct {
	RootPath string `json:"image_root_path"`
}

// 获取配置文件路径
func configRead() string {
	file, err := os.ReadFile("config/RIConfig.json")

	if err != nil {
		log.Printf("读取配置文件失败,异常:%s\n", err.Error())
	}
	var ImageDir RIConfig
	err = json.Unmarshal(file, &ImageDir)
	if err != nil {
		panic(err)
	}

	return ImageDir.RootPath
}

// 获取json文件名(路径)
func GetJsonPath() string {
	if !IsExists("filelist") {
		err := os.MkdirAll("filelist", 0755)
		if err != nil {
			panic(err)
		}
	}
	return "filelist/" + strings.ReplaceAll(configRead(), "/", "_") + ".json"
	//return strings.ReplaceAll("/app/images", "/", "_") + ".json"
}

// 生成json文件
func GenJsonFile() {
	fpath := configRead()
	//fpath := "/app/images"
	if fpath[len(fpath)-1] == '/' {
		fpath = fpath[0 : len(fpath)-1]
	}
	outputname := GetJsonPath()
	var wg sync.WaitGroup
	log.Printf("开始扫描目录:%s\n", fpath)
	Root := filescan.New(fpath)
	filescan.Scan(&Root)
	MaxLine := make(chan struct{}, runtime.NumCPU())
	var result []byte
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		task.Filefound(&Root, MaxLine, &wg)
	}()
	wg.Wait()
	flist := make([]filescan.FileMap, 100000)
	L := filescan.Field(&Root, flist, 100000, "Image", fpath)
	var outputfile struct {
		Rootpath string             `json:"rootpath"` //根目录路径
		Filelist []filescan.FileMap `json:"filelist"` //文件列表

	}
	outputfile.Filelist = flist[:L]
	outputfile.Rootpath = Root.RootDir
	result, err = json.Marshal(outputfile)
	if err != nil {
		fmt.Printf("Map转化为byte数组失败,异常:%s\n", err)
		panic(err)
	}
	CacheLock.Lock()
	defer CacheLock.Unlock()
	err = os.WriteFile(outputname, result, 0644)
	if err != nil {
		panic(err)
	}
}
func Log(s string) {
	log.Println(s + "\n")
}
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
