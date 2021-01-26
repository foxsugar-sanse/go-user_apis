package controllers

import (
	"beego/Web_Project/网站登录注册案例/utils"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

var UserVerifiedCode = make(map[string]string)

var UserNewUserVerifiedCode = make(map[string]string)

func ReturnVcCodeData(c *gin.Context) {
	vccode,imagebase64 := utils.VerificationCodeStart()
	// 保证id唯一不重复
	oldNum := make([]byte,32)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(oldNum); i++ {
		oldNum[i] = byte(rand.Intn(10))
	}
	// 分配验证码的id
	md5obj := md5.New()
	md5obj.Write(oldNum)
	cid := hex.EncodeToString(md5obj.Sum(nil))
	var exagray [6]int
	year,month,day := time.Now().Date()
	times,_,_ := time.Now().Clock()
	oldTimeMTT := time.Now().Minute()
	oldTimeSCD := time.Now().Second()
	exagray[0] = year
	exagray[1] = int(month)
	exagray[2] = day
	exagray[3] = times
	exagray[4] = oldTimeMTT
	exagray[5] = oldTimeSCD
	CookieSetTimeSigns[cid] = exagray
	LoginUseTimeSlice[cid] = strconv.FormatInt(time.Now().Unix(), 10)
	// 设置超时10分钟的cookie
	c.SetCookie("vc_cid",cid,10,"/","127.0.0.1:8080",false,true)
	c.JSON(200,gin.H{
		"status": "ok",
		"test": "base64/image",
		"body":imagebase64,
		"vc_cid":cid,
		"t_s":LoginUseTimeSlice[cid],
	})
	UserVerifiedCode[cid] = vccode
}

func ReturnNewUserVcCodeData(c * gin.Context)  {
	// 生成用户的注册验证码
	vccode,imagebase64 := utils.VerificationCodeStart()
	// 保证id唯一不重复
	oldNum := make([]byte,32)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(oldNum); i++ {
		oldNum[i] = byte(rand.Intn(10))
	}
	// 分配验证码的id
	md5obj := md5.New()
	md5obj.Write(oldNum)
	cid := hex.EncodeToString(md5obj.Sum(nil))
	// 存储cookie设置的时间
	var exagray [6]int
	year,month,day := time.Now().Date()
	times,_,_ := time.Now().Clock()
	oldTimeMTT := time.Now().Minute()
	oldTimeSCD := time.Now().Second()
	exagray[0] = year
	exagray[1] = int(month)
	exagray[2] = day
	exagray[3] = times
	exagray[4] = oldTimeMTT
	exagray[5] = oldTimeSCD
	CookieSetTimeNewUser[cid] = exagray
	CookieSetTimeRegis[cid] = strconv.FormatInt(time.Now().Unix(), 10)
	// 设置超时10分钟的cookie
	c.SetCookie("new_user_vc_cid",cid,10,"/","127.0.0.1:8080",false,true)
	c.JSON(200,gin.H{
		"status": "ok",
		"test": "base64/image",
		"body":imagebase64,
		"new_user_vc_cid":cid,
		"t_s":CookieSetTimeRegis[cid],
	})
	UserNewUserVerifiedCode[cid] = vccode
}