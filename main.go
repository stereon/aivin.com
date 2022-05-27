package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stereon/aivin.com/routers"


)


func main() {
	InitConfig()
	r := routers.SetupRouter()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}


func InitConfig() {
	viper.SetConfigName("application") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./config")   // 查找配置文件所在的路径
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
