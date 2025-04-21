package process

import (
	"log"
	"net/http"
	"random_image/utils"
)

// 随机展示图片

func RandomImage(w http.ResponseWriter, r *http.Request) {
	// 调用打开文件函数
	imageData, contentType := utils.GetFile()
	log.Printf("一次请求，图片类型：%s", contentType)
	//http.ServeFile(w, r, ImagePath)

	w.Header().Set("Content-Type", contentType)
	w.Write(imageData)
}
