package main

import (
	"beego/Web_Project/网站登录注册案例/controllers"
	"beego/Web_Project/网站登录注册案例/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	routers.Router(r)
	// 循环检测验证码id的生命周期并在服务器删除
	go func() {
		for true {
			controllers.GC_VC_ID()
		}

	}()
	//改变配置
	//localpath, _ := os.Getwd()
	//println(localpath)
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("/Users/harder/GoLang_Project/Web_Project/网站登录注册案例/config")
	// 读取配置文件
	err1 := viper.ReadInConfig()
	if err1 != nil {
		panic(err1)
	}
	// 改变其他文件的配置
	controllers.SettingsData["sqlite3_file"] = viper.GetString("database.sqlite3.file")
	ip_port := viper.GetString("server.IPAddress")+viper.GetString("server.port")
	err := r.Run(ip_port)
	if err != nil {
		panic(err)
	}
}
