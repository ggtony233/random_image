package router

import (
	"net/http"
	"process"
	"random_image/utils"
	"time"
)

func Router() {
	utils.ReadOneFile()
	utils.StartAutoRefresh(10 * time.Minute)
	http.HandleFunc("/", process.RandomImage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
