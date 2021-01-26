package main

import (
	"beego/Web_Project/网站登录注册案例/config"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserName string
	UserPassword string
	Uid string
}

func main() {
	db, err := gorm.Open("sqlite3",config.DatabaseDataPath())
	if err != nil {panic("数据库连接失败")}
	defer db.Close()
	// 验证密码和用户名
	var users []*User
	db.Where("user_name = ? AND user_password = ?", "xiaomi", "fb7c5776f2ec72f0e919403a1bd3f935").Find(&users)
	println(users[0].Uid)
}
