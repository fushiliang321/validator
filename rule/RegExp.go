package rule

import (
	"fmt"
	"github.com/fushiliang321/validator/rule/regExp"
	"github.com/fushiliang321/validator/value"
	"regexp"
)

func init() {
	Register("alpha", alpha)
	Register("alpha_dash", alphaDash)
	Register("alpha_num", alphaNum)
	Register("ascii", ascii)
	Register("mac_address", macAddress)
	Register("email", email)
	Register("phone", phone)
	Register("regex", regex)
}

func regExpBase(data *value.Data, fieldName Field, regexp *regexp.Regexp) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}
	for _, _value := range values {
		if b := regexp.MatchString(fmt.Sprintf("%v", _value)); !b {
			return Error("", fieldName, _value, "")
		}
	}
	return
}

// 验证字段是否是字母(包含中文)
func alpha(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.Alpha)
}

// 验证字段可以包含字母(包含中文)和数字，以及破折号和下划线
func alphaDash(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.AlphaDash)
}

// 验证字段可以包含字母(包含中文)和数字，以及破折号和下划线
func alphaNum(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.AlphaNum)
}

// 正在验证的字段是否完全是 7 位的 ASCII 字符
func ascii(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.Ascii)
}

// 验证的字段是否是一个 MAC 地址
func macAddress(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.MacAddress)
}

// 验证的字段是否是一个邮箱地址
func email(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.Email)
}

// 验证的字段是否是一个手机号码
func phone(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return regExpBase(data, fieldName, regExp.Phone)
}

// 根据正则表达式验证字段值
func regex(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	re, err := regexp.Compile(argStr)
	if err != nil {
		return
	}
	return regExpBase(data, fieldName, re)

}
