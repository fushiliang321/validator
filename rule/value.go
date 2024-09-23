package rule

import (
	json2 "encoding/json"
	"fmt"
	"github.com/fushiliang321/validator/utils"
	"github.com/fushiliang321/validator/value"
	"net"
	"slices"
	"unicode"
)

func init() {
	Register("accepted", accepted)
	Register("accepted_if", acceptedIf)
	Register("declined", declined)
	Register("declined_if", declinedIf)
	Register("boolean", boolean)
	Register("different", different)
	Register("lowercase", lowercase)
	Register("uppercase", uppercase)
	Register("ip", ip)
	Register("ipv4", ipv4)
	Register("ipv6", ipv6)
	Register("json", json)
}

// 验证字段的值必须是 yes、on、1 或 true
func accepted(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		valStr     string
	)
	if !ok {
		return
	}
	for _, _value = range values {
		valStr = fmt.Sprintf("%v", _value)
		if valStr != "yes" &&
			valStr != "on" &&
			valStr != "1" &&
			valStr != "true" {
			return Error("", fieldName, _value, "")
		}
	}
	return
}

// 如果另一个正在验证的字段等于指定的值，则验证中的字段必须为 yes、on、1 或 true
func acceptedIf(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" || !utils.FieldEqValue(data, argStr) {
		return
	}
	return accepted(data, fieldName, argStr)
}

// 正在验证的字段必须是 no、off、0 或者 false
func declined(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		valStr     string
	)
	if !ok {
		return
	}
	for _, _value = range values {
		valStr = fmt.Sprintf("%v", _value)
		if valStr != "no" &&
			valStr != "off" &&
			valStr != "0" &&
			valStr != "false" {
			return Error("", fieldName, _value, "")
		}
	}

	return
}

// 如果另一个验证字段的值等于指定值，则验证字段的值必须为 no、off、0 或 false
func declinedIf(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" || !utils.FieldEqValue(data, argStr) {
		return
	}
	return declined(data, fieldName, argStr)
}

// 验证字段必须可以被转化为布尔值，接收 true, false, 1, 0, "1" 和 "0" 等输入
// argStr=="strict"时验证字段必须可以被转化为布尔值，仅接收 true 和 false
func boolean(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		valStr      string
		values, ok  = data.Get(fieldName)
		matchValues []string
	)
	if !ok {
		return
	}
	if argStr == "strict" {
		matchValues = []string{"true", "false"}
	} else {
		matchValues = []string{"true", "false", "0", "1"}
	}

	for _, _value := range values {
		valStr = fmt.Sprintf("%v", _value)
		if slices.Index[[]string, string](matchValues, valStr) == -1 {
			return Error("", fieldName, _value, "")
		}
	}
	return
}

// 验证字段必须是一个和指定字段不同的值
func different(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values1, ok1 = data.Get(fieldName)
		values2, ok2 = data.Get(argStr)
	)
	if !ok1 || !ok2 {
		return
	}
	for _, v1 := range values1 {
		for _, v2 := range values2 {
			if v1 == v2 {
				return Error("", fieldName, v1, "")
			}
		}
	}
	return
}

// 验证的字段必须是小写的
func lowercase(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		r          rune
	)
	if !ok {
		return
	}

	for _, _value = range values {
		switch v := _value.(type) {
		case string:
			for _, r = range v {
				if !unicode.IsLower(r) {
					return Error("", fieldName, v, "")
				}
			}
		default:
			return Error("", fieldName, v, "")
		}
	}
	return
}

// 验证字段必须为大写
func uppercase(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		r          rune
	)
	if !ok {
		return
	}

	for _, _value = range values {
		switch v := _value.(type) {
		case string:
			for _, r = range v {
				if !unicode.IsUpper(r) {
					return Error("", fieldName, v, "")
				}
			}
		default:
			return Error("", fieldName, v, "")
		}
	}
	return
}

// 验证字段必须是 IP 地址
func ip(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
	)
	if !ok {
		return
	}
	for _, _value = range values {
		if net.ParseIP(fmt.Sprintf("%v", _value)) == nil {
			return Error("", fieldName, _value, "")
		}
	}
	return
}

// 验证字段必须是 IPv4 地址
func ipv4(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		_ip        net.IP
	)
	if !ok {
		return
	}

	for _, _value = range values {
		_ip = net.ParseIP(fmt.Sprintf("%v", _value))
		if _ip == nil || _ip.To4() == nil {
			return Error("", fieldName, _value, "")
		}

	}
	return
}

// 验证字段必须是 IPv6 地址
func ipv6(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		_ip        net.IP
	)
	if !ok {
		return
	}

	for _, _value = range values {
		_ip = net.ParseIP(fmt.Sprintf("%v", _value))
		if _ip == nil || _ip.To16() == nil || _ip.To4() != nil {
			return Error("", fieldName, _value, "")
		}

	}
	return
}

// 验证字段必须是有效的 JSON 字符串
func json(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		js         json2.RawMessage
	)
	if !ok {
		return
	}

	for _, _value = range values {
		switch v := _value.(type) {
		case string:
			if json2.Unmarshal([]byte(v), &js) != nil {
				return Error("", fieldName, _value, "")
			}
		default:
			return Error("", fieldName, _value, "")
		}
	}
	return
}
