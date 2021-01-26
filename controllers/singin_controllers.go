package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// 常量区
// Md5加盐
const _SHAYOUR_          = "0x23FF^55ShaTheMD5TokensReader0x235GfdX"

// SessionID加盐
const _SESSIONIDSHAYOUR_ =  "0x67#$%^&SESSIONSSUUUIIID000ddd534ddsd"

type User struct {
	// gorm映射数据表
	UserName 		string
	UserPassword 	string
	Uid 			string
}

type NewUser struct {
	UserName 		string `json:"username"`
	UserPassword 	string `json:"password"`
	UserVcCode 		string `json:"user_vc_code"`
	TimeSlice 		string `json:"t_s"`
}

// 存储登录用户验证码cookie的时间
var CookieSetTimeSigns = make(map[string][6]int)

// 存储注册用户验证码Cookie的时间
var CookieSetTimeRegis = make(map[string]string)

// 存储登录用的时间
var LoginUseTimeSlice = make(map[string]string)

// 存储注册用户验证码的cookie时间
var CookieSetTimeNewUser = make(map[string][6]int)

// 在本地中管理用户登录的sessionid
var CookieSessionIdLogins = make(map[string]int64)

// sessionID与登录用户uid的绑定，身份验证
var UserSessionIdAndUid = make(map[string]string)

type ResetUserPwdBind struct {
	Uid 				string
	TimeSlice			string `json:"t_s"`
	UserPassword 		string `json:"password"`
	UserRePassword 		string `json:"rewtpassword"`
	UserRePasswordTwo 	string `json:"rewtpasswordtwo"`
}

type SetUserInforMations struct {
	TimeSlice		string  `json:"t_s"`
	Usernickname 	string  `json:"usernickname"`
	Userage 		string	`json:"userage"`
	Usergender 		string	`json:"usergender"`
	Usercontact 	string	`json:"usercontact"`
	Userbriefid 	string	`json:"userbriefid"`
}

type UserData struct {
	//username  		string
	//userpassword 	string
	Uid 			string
	Usernickname 	string
	Userage 		string
	Usergender 		string
	Usercontact 	string
	Userbriefid 	string

}

type UserLoginBind struct {
	TimeSlice		string `json:"t_s"`
	UserName 		string `json:"username"`
	UserPassword 	string `json:"password"`
	UserVcCode		string `json:"user_vc_code"`
}

func Signs (c *gin.Context) {
	//登录
	json := UserLoginBind{}
	err := c.BindJSON(&json)
	if err != nil{
		panic("绑定失败！")
	}
	// 获取用户登录信息
	userLoginInfo, err := c.Cookie("user_login")
	// 匹配登录信息
	oldTimeSlice := CookieSessionIdLogins[userLoginInfo]
	if oldTimeSlice != 0 {
		c.JSON(200,gin.H{
			"status":"ok",
			"info":"登录成功",
		})
		c.Abort()
	} else {

	}
	//验证码功能
	//获得客户端的验证码id
	cookie, err := c.Cookie("vc_cid")
	if err != nil {
		c.JSON(500, gin.H{
			"status":"no",
			"info":"请先获取验证码",
		})
		panic("客户端没有验证码ID!")
	}
	// 匹配服务器中的验证码id
	rangenum := 0
	// 服务器中的验证码id没有的情况
	if len(UserVerifiedCode)==0 {
		c.JSON(500, gin.H{
			"status":"no",
			"info":"验证码未获取或已过期",
		})
	}
	for k := range UserVerifiedCode {
		if cookie != k {
			if rangenum == len(UserVerifiedCode) - 1 {
				c.JSON(500, gin.H{
					"status":"no",
					"info":"验证码已过期",
				})
			}
		} else {
			rangenum++
		}
	}


	if (json.UserVcCode == UserVerifiedCode[cookie]) && (json.UserVcCode != "") {
		// 验证码验证成功后校验时间
		if json.TimeSlice == LoginUseTimeSlice[cookie] {
			goto OK
		} else {
			c.JSON(500,gin.H{
				"status":"no",
				"info":"登录失败",
			})
			panic("前端提交时间错误")
		}
		OK:
		// 验证验证码成功之后执行数据库查询逻辑
		md5obj := md5.New()
		md5obj.Write([]byte(json.UserPassword + _SHAYOUR_))
		json.UserPassword = hex.EncodeToString(md5obj.Sum(nil))
		usermod := UserLoginBind{
			UserName:     json.UserName,
			UserPassword: json.UserPassword,
		}
		isOk, uid := UserLoginDataBaseController(&usermod)
		// Uid + 盐值 + 时间戳生成SessionID
		use_time := time.Now().Unix()
		uuid := func() string{
			md5obj2 := md5.New()
			md5obj2.Write([]byte(uid + _SESSIONIDSHAYOUR_ + strconv.FormatInt(use_time, 10)))
			return hex.EncodeToString(md5obj2.Sum(nil))
		}()
		// 存入登录用户
		CookieSessionIdLogins[uuid] = use_time
		UserSessionIdAndUid[uuid] = uid
		if isOk {
			// 验证码与id删除方便下次操作
			delete(UserVerifiedCode,cookie)
			c.SetCookie("user_login",uuid,7200,"/","127.0.0.1",false,true)
			c.JSON(200, gin.H{
				"status":"ok",
				"info":"登录成功",
			})
		} else {
			c.JSON(500, gin.H{
				"status":"no",
				"info":"密码或用户名不正确",
			})
		}
	} else if (json.UserVcCode != UserVerifiedCode[cookie]) && (len(UserVerifiedCode) > 0) {
		c.JSON(500, gin.H{
			"status":"no",
			"info":"验证码不正确",
			"code":UserVerifiedCode,
		})
	}
}


func New(c *gin.Context) {
	// 新建用户
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	json := NewUser{}
	_ = c.BindJSON(&json)
	// TODO: 数据库存入用户信息
	// 新建用户
	// 密码加盐md5存储
	cookie, err := c.Cookie("new_user_vc_cid")
	if err != nil {
		c.JSON(500,gin.H{
			"status":"no",
			"info":"未获得验证码",
		})
	}
	// 匹配服务器中的验证码id
	rangenum := 0
	// 服务器中的验证码id没有的情况
	if len(UserNewUserVerifiedCode)==0 {
		c.JSON(500, gin.H{
			"status":"no",
			"info":"验证码未获取或已过期",
		})
	}
	for k := range UserNewUserVerifiedCode {
		if cookie != k {
			if rangenum == len(UserNewUserVerifiedCode) - 1 {
				c.JSON(500, gin.H{
					"status":"no",
					"info":"验证码已过期",
				})
			}
		} else {
			rangenum++
		}
	}
	if (json.UserVcCode == UserNewUserVerifiedCode[cookie]) && (json.UserVcCode != "") {
		// 校验时间
		if  json.TimeSlice == ""{
			c.JSON(500,gin.H{
				"status":"ok",
				"info":"参数不能为空",
			})
			c.Abort()
		} else if CookieSetTimeRegis[cookie] == json.TimeSlice {
			goto OK
		}else {
			c.JSON(500,gin.H{
				"status":"ok",
				"info":"时间不正确",
			})
			c.Abort()
		}
		OK:
		md5obj := md5.New()
		md5obj.Write([]byte(json.UserPassword + _SHAYOUR_))
		// 为用户设置一个uid
		md5obj2 := md5.New()
		md5obj2.Write([]byte(json.UserName))
		uid := hex.EncodeToString(md5obj2.Sum(nil))
		usermod := User{
			UserName: json.UserName,
			UserPassword: hex.EncodeToString(md5obj.Sum(nil)),
			Uid: uid,
		}
		if DatabaseController(&usermod) == true {
			// 验证成功删除此验证码
			delete(UserNewUserVerifiedCode,cookie)
			c.JSON(200, gin.H{
				"status":   "ok",
				"info":     "注册成功",
				"username": json.UserName,
				"uid":      uid,
				//TODO:debug "usermod":usermod,
			})
		} else {
			c.JSON(500, gin.H{
				"status": "no",
				"info":   "该用户已存在",
			})
		}
	} else if (json.UserVcCode != UserVerifiedCode[cookie]) && (len(UserNewUserVerifiedCode) > 0) {
		c.JSON(500, gin.H{
			"status":"no",
			"info":"验证码不正确",
			"code":UserNewUserVerifiedCode,
		})
	}
}

func ResetUserPwd(c *gin.Context) {
	//TODO	重新设置用户密码
	jsons := ResetUserPwdBind{}
	err := c.BindJSON(&jsons)
	if err != nil {
		panic("错误！")
	}
	uuid,err := c.Cookie("user_login")
	if err != nil {
		c.JSON(500,gin.H{
			"status":"no",
			"info":"登录信息过期",
		})
		c.Abort()
	}
	// 校验时间
	times,_ := strconv.ParseInt(jsons.TimeSlice,10,64)
	if times + 300 > time.Now().Unix() {
		// 客户端与服务器的时间不对齐，请求驳回
		c.JSON(500,gin.H{
			"status":"no",
			"info":"time no",
		})
		c.Abort()
	} else {

	}
	uid := UserSessionIdAndUid[uuid]
	if jsons.UserRePassword == jsons.UserRePasswordTwo {
		// 判断密码与确认密码是否相同
		md5obj := md5.New()
		md5obj.Write([]byte(jsons.UserRePassword + _SHAYOUR_))
		md5obj2 := md5.New()
		md5obj2.Write([]byte(jsons.UserPassword + _SHAYOUR_))
		usermod := ResetUserPwdBind{
			Uid:          uid,
			UserPassword:      hex.EncodeToString(md5obj2.Sum(nil)),
			UserRePassword:    hex.EncodeToString(md5obj.Sum(nil)),
			//UserRePasswordTwo: json.UserRePasswordTwo, // 确认密码，后端验证
		}
		if ReSetUserPwdController(&usermod) == true {
			c.JSON(http.StatusOK,gin.H{
				"status": "ok",
				"info":"更改密码成功",
			})
		} else {
			c.JSON(http.StatusOK,gin.H{
				"status": "no",
				"info":"请输入正确用户名或者原密码",
			})
		}
	} else {
		c.JSON(400,gin.H{
			"status": "no",
			"info":"更改失败请检查密码与确认密码",
		})
	}

}

func PushUserInforMations(c * gin.Context) {
	// TODO 提交关于用户的信息
	json := SetUserInforMations{}
	_ = c.BindJSON(&json)
	cookie, err := c.Cookie("user_login")
	if err != nil {
		c.JSON(500, gin.H{
			"status":"no",
			"info":"未登录请登录",
		})
		panic("用户未登录尝试更改信息")
	}
	// 获得与令牌对应的用户id
	uid := UserSessionIdAndUid[cookie]
	if uid == "" {
		c.JSON(500,gin.H{
			"status":"no",
			"info":"令牌非法或已过期",
		})
		c.Abort()
	}
	times, _ := strconv.ParseInt(json.TimeSlice,10,64)
	if times + 300 > time.Now().Unix() {
		// 校验客户端和服务端的时间
		goto OK
	} else {
		c.JSON(500,gin.H{
			"status":"no",
			"info":"time no",
		})
		panic("客户端时间错误！")
	}
	OK:
	usermod := UserData{
		Uid:          uid,
		Usernickname: json.Usernickname,
		Userage:      json.Userage,
		Usergender:   json.Usergender,
		Usercontact:  json.Usercontact,
		Userbriefid:  json.Userbriefid,
	}
	if ReSetUserInfoemationController(&usermod) {
		c.JSON(200, gin.H{
			"status":   "ok",
			"info":"用户信息更改成功",
		})
	} else {
		c.JSON(500, gin.H{
			"status":   "no",
			"info":"用户信息更改失败",
		})
	}
}
