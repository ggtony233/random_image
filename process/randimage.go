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
	imageData, contentType, _ := utils.GetFile()

	//http.ServeFile(w, r, ImagePath)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(imageData)))
	w.Header().Set("Content-Type", contentType)

	//w.Header().Set("Content-Disposition", "attachment; filename="+Name)
	_, err := io.Copy(w, bytes.NewReader(imageData))
	if err != nil {
		log.Printf("%s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
func Refresh(w http.ResponseWriter, r *http.Request) {
	// 调用刷新函数
	err := utils.ReadOneFile()
	if err != nil {
		log.Printf("%s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("刷新成功"))

}
