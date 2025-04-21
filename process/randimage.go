package process

import (
	"bytes"
	"fmt"
	"io"
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
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(imageData)))
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=86400")
	_, err := io.Copy(w, bytes.NewReader(imageData))
	if err != nil {
		log.Printf("%s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
