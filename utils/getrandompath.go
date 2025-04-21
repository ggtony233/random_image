package utils

import (
	"encoding/json"
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
		GenJsonFile()
		jsonfile, _ = os.ReadFile(jsonlocation)

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
	Ipath := RandomImagePath(GetJsonPath())
	Log("读取图片...")
	data, err := os.ReadFile(Ipath)
	if err != nil {
		return err
	}
	Log("图片" + Ipath + "读取成功")
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
			os.Stdout.WriteString("刷新图片缓存...")
			if err := ReadOneFile(); err != nil {
				os.Stdout.WriteString("缓存刷新失败：" + err.Error())
			} else {
				os.Stdout.WriteString("图片缓存已刷新")
			}
		}
	}()
}
