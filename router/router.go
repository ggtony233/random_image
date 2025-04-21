package router

import (
	"net/http"
	"process"
	"random_image/utils"
	"time"
)

func Router() {
	utils.ReadOneFile()
	utils.Log("初始化文件读取完成")
	utils.StartAutoRefresh(10 * time.Minute)
	utils.RefreshFilelist(10 * time.Hour)
	http.HandleFunc("/", process.RandomImage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
