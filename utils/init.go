package utils

import (
	"log"
	"os"
)

func init() {
	_, err := os.Stat("config/RIConfig.json")
	if os.IsNotExist(err) {
		log.Printf("配置文件不存在,创建配置文件")
		if IsExists("config") {
			err1 := os.MkdirAll("config", 0755)
			if err1 != nil {
				log.Printf("创建配置文件目录失败,异常:%s\n", err1.Error())
				os.Exit(1)
			}
		}

		//err1 = os.WriteFile("config/RIConfig.json", []byte(`{"image_root_path":"/app/images"}`), 0644)
		err = os.WriteFile("config/RIConfig.json", []byte(`{"image_root_path":"/WindowsData/Windows Data/Picture/水淼Aqua 213套合集[48.9G]"}`), 0644)
		log.Println("初始化配置文件")
		if err != nil {
			log.Printf("创建配置文件失败,异常:%s\n", err.Error())
			os.Exit(1)
		}
	}
}
