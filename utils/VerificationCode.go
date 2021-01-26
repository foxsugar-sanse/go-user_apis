package utils

import (
	"github.com/mojocn/base64Captcha"
)

func VerificationCodeStart() (string,string){
	var configC = base64Captcha.ConfigCharacter{
		Height: 88,
		Width:  240,
		// const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexMedium,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexMedium,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         6,
	}
	// GenerateCaptcha 第一个参数为空字符串, 包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	// 以base64编码保存为文件
	//fmt.Println(base64Captcha.CaptchaWriteToBase64Encoding(capC))
	//err := base64Captcha.CaptchaWriteToFile(capC, "./", idKeyC, "png")
	//if err != nil {
	//	panic("验证码图片保存失败")
	//}
	// 获取正确答案
	VerifyValue := capC.(*base64Captcha.CaptchaImageChar).VerifyValue
	// 返回答案和验证码的base64编码
	return VerifyValue, base64Captcha.CaptchaWriteToBase64Encoding(capC)
	// 展示正确答案
	//println(verifyValue)
	// 删除内存键防止内存泄漏
	defer base64Captcha.VerifyCaptchaAndIsClear(idKeyC, VerifyValue, true)
	// os.Remove("./" + idKeyC + "png")
	return "", ""
}
