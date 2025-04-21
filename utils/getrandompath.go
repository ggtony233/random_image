package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var ImageData []byte
var Type string
var CacheLock sync.RWMutex

type Myfile struct {
	Path  string `json:"path"`
	Ftype string `json:"type"`
}
type TrueFilelist struct {
	RootPath string   `json:"rootpath"`
	Files    []Myfile `json:"filelist"`
}

func RandomImagePath(jsonlocation string) string {
	//读取json文件
	jsonfile, _ := os.ReadFile(jsonlocation)
	if jsonfile == nil {
		return ""
		//generate json file
		//		utils.GenJsonFile()
		//		jsonfile, _ = os.ReadFile(jsonlocation)
	}

	FileList := TrueFilelist{
		RootPath: "",
		Files:    []Myfile{},
	}
	err := json.Unmarshal(jsonfile, &FileList)
	if err != nil {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return FileList.RootPath + FileList.Files[r.Intn(len(FileList.Files))].Path
	//如果每天只展示一张图片，则直接读取json文件，否则运行json生成函数，并写入json文件
	//读取json文件，返回图片路径

}
func ReadOneFile() error {
	data, err := os.ReadFile(RandomImagePath(GetJsonPath()))
	if err != nil {
		return err
	}
	CacheLock.Lock()
	defer CacheLock.Unlock()
	Type = http.DetectContentType(data)
	ImageData = data
	return nil
}
func GetFile() ([]byte, string) {
	CacheLock.RLock()
	defer CacheLock.RUnlock()
	return ImageData, Type
}

// StartAutoRefresh 启动定时缓存刷新
func StartAutoRefresh(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			fmt.Println("刷新图片缓存...")
			if err := ReadOneFile(); err != nil {
				fmt.Println("缓存刷新失败：", err)
			} else {
				fmt.Println("图片缓存已刷新")
			}
		}
	}()
}
