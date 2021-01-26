package controllers

import (
	"beego/Web_Project/网站登录注册案例/config"
	//"beego/Web_Project/网站登录注册案例/config"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var SettingsData = map[string]string{
	"sqlite3_file" : "",
}

func DatabaseController(user *User) bool {
	db, err := gorm.Open("sqlite3",SettingsData["sqlite3_file"])
	if err != nil {panic("数据库连接失败！")}
	defer func() {
		db.Close()
	}()
	// 避免重复，不允许出现重复的用户名
	var users []*User
	db.Where("user_name = ?",user.UserName).Find(&users)
	if len(users) > 0 {
		return false
	} else {
		user_data := UserData{
			Uid: user.Uid,
		}
		db.Create(&user)
		db.Create(user_data)
		return true
	}
}

func ReSetUserPwdController(user *ResetUserPwdBind) bool{
	db, err := gorm.Open("sqlite3","sqlite3",SettingsData["sqlite3_file"])
	if err != nil {panic("数据库连接失败")}
	defer db.Close()
	// 验证密码
	var users []*User
	db.Where("uid = ? AND user_password = ?", user.Uid, user.UserPassword).Find(&users)
	if len(users) > 0 {
		// 账户检测存在，密码匹配进入
		// 先匹配用户名之后再更改密码
		err := db.Model(&User{}).Where("uid = ?",user.Uid).Update("user_password",user.UserRePassword)
		if err != nil {panic("查询失败")}
		return true
	} else {
		return false
	}
}

func UserLoginDataBaseController(user *UserLoginBind)  (bool,string) {
	// 登录的数据库逻辑
	db, err := gorm.Open("sqlite3","sqlite3",SettingsData["sqlite3_file"])
	if err != nil {panic("数据库连接失败")}
	defer db.Close()
	// 验证密码和用户名
	var users []*User
	db.Where("user_name = ? AND user_password = ?", user.UserName, user.UserPassword).Find(&users)
	if len(users) > 0 {
		return true,users[0].Uid
	} else {
		return false,""
	}

}

func ReSetUserInfoemationController(user * UserData) bool{
	db, err := gorm.Open("sqlite3",config.DatabaseDataPath())
	if err != nil {panic("数据库连接失败")}
	defer db.Close()
	// 验证uid
	var users []*UserData
	db.Where("uid = ?", user.Uid).Find(&users)
	if len(users) > 0 {
		//users[0].Userage = user.Userage
		//users[0].Usernickname = user.Usernickname
		//users[0].Usergender = user.Usergender
		//users[0].Userbriefid = user.Userbriefid
		//users[0].Usercontact = user.Usercontact
		//db.Save(&users)

		// 使用Updates更新多个字段
		db.Model(&UserData{}).Where("uid = ?", user.Uid).Updates(map[string]interface{}{
			"Userage" 	   : user.Userage,
			"Usernickname" : user.Usernickname,
			"Usergender"   : user.Usergender,
			"Userbriefid"  : user.Userbriefid,
			"Usercontact"  : user.Usercontact,
		})

		// 逐条更新
		//if user.Usercontact != "" {
		//	err2 := db.Model(&UserData{}).Where("uid = ?",user.Uid).Update("usercontact",user.Usercontact)
		//	if err2 != nil {
		//		panic(err2)
		//	}
		//}
		//if user.Userage != "" {
		//	db.Model(&UserData{}).Where("uid = ?",user.Uid).Update("userage",user.Userage)
		//}
		//if user.Usernickname != "" {
		//	db.Model(&UserData{}).Where("uid = ?",user.Uid).Update("usernickname",user.Usernickname)
		//}
		//if user.Usergender != "" {
		//	db.Model(&UserData{}).Where("uid = ?",user.Uid).Update("usergender",user.Usergender)
		//}
		//if user.Userbriefid != "" {
		//	db.Model(&UserData{}).Where("uid = ?",user.Uid).Update("userbriefid",user.Userbriefid)
		//}

		return true
	} else {
		return false
	}

}
