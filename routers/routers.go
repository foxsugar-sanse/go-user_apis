package routers

import (
	"beego/Web_Project/网站登录注册案例/config"
	"beego/Web_Project/网站登录注册案例/controllers"
	"github.com/gin-gonic/gin"
)

func Router(r * gin.Engine) {
	// 静态文件路径
	r.Static("/web",config.Path())
	v1 := r.Group("v1")
	{
		v1.POST("/user/new", controllers.New)
		v1.POST("/user/signs", controllers.Signs)
		v1.POST("/user/reset", controllers.ResetUserPwd)
		v1.POST("/user/pushuserinfo", controllers.PushUserInforMations)

	}
	resources := r.Group("resources")
	{
		resources.GET("/user_vc_code",controllers.ReturnVcCodeData)
		resources.GET("/new_user_vc_code", controllers.ReturnNewUserVcCodeData)
		//resources.GET("/cookie", func(c *gin.Context) {
		//	// 获取cookie
		//	cookie, err := c.Cookie("gin_cookie")
		//	if err != nil {
		//		cookie = "NotSet"
		//		c.SetCookie("gin_cookie","test",3600,"/","127.0.0.1:8080",false,true)
		//	}
		//	println(cookie)
		//})
	}

}

