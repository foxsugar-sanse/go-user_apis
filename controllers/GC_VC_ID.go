package controllers

import (
	"beego/Web_Project/网站登录注册案例/utils"
	"time"
)

// 7天秒速常量
const DAY7SECOND int64 = 604800

// 10分钟秒速常量
const MINUTETEN = 600

// debug 1分钟
const DEBUGSECOND = 60

func GC_VC_ID() {
	// 键-vc_id 值-时间
	newYear,newMonth,newDay := time.Now().Date()
	newTimeTM  := time.Now().Hour()
	newTimeMIT := time.Now().Minute() // 分钟
	newTimeSCD := time.Now().Second() // 秒钟

	for k := range CookieSetTimeSigns {
		// 清除过期登录验证码匹配id
		if utils.SwitchForTime(newYear, int(newMonth), newDay, newTimeTM, newTimeMIT, newTimeSCD, CookieSetTimeSigns[k][0],CookieSetTimeSigns[k][1],CookieSetTimeSigns[k][2],CookieSetTimeSigns[k][3],CookieSetTimeSigns[k][4],CookieSetTimeSigns[k][5],">=","minute",10) {
			delete(UserVerifiedCode,k)
			delete(CookieSetTimeSigns,k)
		}
		// TODO old
		//if (newYear == CookieSetTimeSigns[k][0] && int(newMonth) == CookieSetTimeSigns[k][1]) && (newDay == CookieSetTimeSigns[k][2] && newTimeTM == CookieSetTimeSigns[k][3]) {
		//	if newTimeMIT - CookieSetTimeSigns[k][4] > 10 {
		//		if CookieSetTimeSigns[k][5] > newTimeSCD {
		//			// 清除匹配的过期id
		//			delete(UserVerifiedCode,k)
		//			delete(CookieSetTimeSigns,k)
		//		}
		//	}
		//}

	}

	for k := range CookieSetTimeNewUser {
		// 清除过期注册验证码匹配id
		if utils.SwitchForTime(newYear, int(newMonth), newDay, newTimeTM, newTimeMIT, newTimeSCD, CookieSetTimeNewUser[k][0],CookieSetTimeNewUser[k][1],CookieSetTimeNewUser[k][2],CookieSetTimeNewUser[k][3],CookieSetTimeNewUser[k][4],CookieSetTimeNewUser[k][5],">=","hour",10) {
			delete(UserNewUserVerifiedCode,k)
			delete(CookieSetTimeSigns,k)
		}
		// TODO Old
		//if (newYear == CookieSetTimeNewUser[k][0] && int(newMonth) == CookieSetTimeNewUser[k][1]) && (newDay == CookieSetTimeNewUser[k][2] && newTimeTM == CookieSetTimeNewUser[k][3]) {
		//	if newTimeMIT - CookieSetTimeNewUser[k][4] > 10 {
		//		if CookieSetTimeNewUser[k][5] > newTimeSCD {
		//			// 清除匹配的过期id
		//			delete(UserVerifiedCode,k)
		//			delete(CookieSetTimeNewUser,k)
		//		}
		//	}
		//}
	}
	newTimeSlice := time.Now().Unix()
	for k,v := range CookieSessionIdLogins{
		// 清除过期的登录信息
		//time_s, _ := strconv.ParseInt(v,10,64)
		if v + DAY7SECOND < newTimeSlice {
			delete(CookieSessionIdLogins,k)
			delete(UserSessionIdAndUid,k)
		}
	}
	//time.Sleep(time.Second * 1)
}
