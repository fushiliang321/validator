package rule

import (
	json2 "encoding/json"
	"fmt"
	"github.com/fushiliang321/validator/utils"
	"github.com/fushiliang321/validator/value"
	"net"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"
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
	Register("after", after)
	Register("after_or_equal", afterOrEqual)
	Register("before", before)
	Register("before_or_equal", beforeOrEqual)
	Register("date_equal", dateEqual)
	Register("size", size)
	Register("gt", gt)
	Register("gte", gte)
	Register("lt", lt)
	Register("lte", lte)
	Register("between", between)
	Register("confirmed", confirmed)
	Register("decimal", decimal)
	Register("in", in)
	Register("not_in", notIn)
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

// 时间比较
func timeCompare(data *value.Data, fieldName Field, argStr string, _func func(valueDate, compareDate *time.Time) bool) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		v          string
		parseDate  time.Time
		timeArr    []time.Time
	)
	if !ok {
		return
	}

	var t, err = utils.StrToTime(argStr)
	if err != nil {
		if values2, ok := data.Get(argStr); ok {
			for _, v2 := range values2 {
				if v2Str, ok := v2.(string); ok {
					t, err = utils.StrToTime(v2Str)
					if err != nil {
						return Error("", fieldName, _value, "")
					}
					timeArr = append(timeArr, t)
				} else {
					return Error("", fieldName, _value, "")
				}
			}
		} else {
			return Error("", fieldName, _value, "")
		}
	} else {
		timeArr = append(timeArr, t)
	}

	for _, _value = range values {
		v, ok = _value.(string)
		if !ok {
			return Error("", fieldName, _value, "")
		}
		parseDate, err = utils.ParseDate(v)
		if err != nil {
			return Error("", fieldName, _value, "")
		}

		for _, t = range timeArr {
			if !_func(&parseDate, &t) {
				return Error("", fieldName, _value, "")
			}
		}
	}
	return
}

// 验证字段必须是给定日期之后的一个值
func after(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return timeCompare(data, fieldName, argStr, func(valueDate, compareDate *time.Time) bool {
		return valueDate.After(*compareDate)
	})
}

// 验证字段必须是大于等于给定日期的值
func afterOrEqual(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return timeCompare(data, fieldName, argStr, func(valueDate, compareDate *time.Time) bool {
		return valueDate.After(*compareDate) || valueDate.Equal(*compareDate)
	})
}

// 验证字段必须是给定日期之前的一个值
func before(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return timeCompare(data, fieldName, argStr, func(valueDate, compareDate *time.Time) bool {
		return valueDate.Before(*compareDate)
	})
}

// 验证字段必须是小于等于给定日期的值
func beforeOrEqual(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return timeCompare(data, fieldName, argStr, func(valueDate, compareDate *time.Time) bool {
		return valueDate.Before(*compareDate)
	})
}

// 验证字段必须等于给定日期
func dateEqual(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return timeCompare(data, fieldName, argStr, func(valueDate, compareDate *time.Time) bool {
		return valueDate.Equal(*compareDate)
	})
}

// 验证字段大小在给定的最小值和最大值之间，字符串、数字、数组和map都可以
func between(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
	)
	if !ok {
		return
	}

	args := strings.Split(argStr, ",")
	if len(args) < 2 {
		return
	}
	var (
		minI, _ = strconv.Atoi(args[0])
		maxI, _ = strconv.Atoi(args[1])
		v       int
	)
	if minI >= maxI {
		return
	}

	for _, _value = range values {
		switch _v := _value.(type) {
		case int:
			v = _v
		case int8:
			v = int(_v)
		case int16:
			v = int(_v)
		case int32:
			v = int(_v)
		case int64:
			v = int(_v)
		case uint:
			v = int(_v)
		case uint8:
			v = int(_v)
		case uint16:
			v = int(_v)
		case uint32:
			v = int(_v)
		case uint64:
			v = int(_v)
		case float64:
			if _v > float64(maxI) || _v < float64(minI) {
				return Error("", fieldName, _value, "")
			}
			continue
		case float32:
			if _v > float32(maxI) || _v < float32(minI) {
				return Error("", fieldName, _value, "")
			}
			continue
		default:
			v, ok = func() (int, bool) {
				defer func() {
					if err := recover(); err != nil {
					}
				}()
				return reflect.ValueOf(_value).Len(), true
			}()
			if !ok {
				return Error("", fieldName, _value, "")
			}
		}
		if v > maxI || v < minI {
			return Error("", fieldName, _value, "")
		}
	}

	return
}

// 比较验证字段，这两个字段类型必须一致，适用于字符串、数字、数组和map
// argStr为数值时，数字比较大小，字符串、数组和map比较长度
func compare(data *value.Data, fieldName Field, argStr string, _func func(v1, v2 float64) bool) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
	)
	if !ok || argStr == "" {
		return
	}
	if size, err := strconv.Atoi(argStr); err == nil {
		//比较数值大小
		var (
			size64   = float64(size)
			valueLen float64
		)
		for _, _value = range values {
			valueLen, err = utils.AnyToFloat64(_value)
			if err != nil {
				//比较长度
				valueLen, ok = func() (float64, bool) {
					defer func() {
						recover()
					}()
					return float64(reflect.ValueOf(_value).Len()), true
				}()
				if !ok {
					return Error("", fieldName, _value, "")
				}
			}
			if !_func(valueLen, size64) {
				return Error("", fieldName, _value, "")
			}
		}
	} else {
		//比较字段
		argValues, ok := data.Get(argStr)
		if !ok {
			return Error("", fieldName, _value, "")
		}
		for _, _value = range values {
			for _, argValue := range argValues {
				v1, err := utils.AnyToFloat64(_value)
				if err == nil {
					//比较数值大小
					v2, err := utils.AnyToFloat64(argValue)
					if err != nil || !_func(v1, v2) {
						return Error("", fieldName, _value, "")
					}
				} else {
					//比较长度
					v1Len, v2Len, b := func() (v1Len, v2Len int, b bool) {
						defer func() {
							recover()
						}()
						return reflect.ValueOf(_value).Len(), reflect.ValueOf(argValue).Len(), true
					}()
					if !b || !_func(float64(v1Len), float64(v2Len)) {
						return Error("", fieldName, _value, "")
					}
				}
			}
		}
	}

	return
}

// 验证字段必须大于给定 field 字段，这两个字段类型必须一致，适用于字符串、数字、数组和map
// argStr为数值时，数字比较大小，字符串、数组和map比较长度
func gt(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return compare(data, fieldName, argStr, func(v1, v2 float64) bool {
		return v1 > v2
	})
}

// 验证字段必须大于等于给定 field 字段，这两个字段类型必须一致，适用于字符串、数字、数组和map
// argStr为数值时，数字比较大小，字符串、数组和map比较长度
func gte(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return compare(data, fieldName, argStr, func(v1, v2 float64) bool {
		return v1 >= v2
	})
}

// 验证字段必须小于给定 field 字段，这两个字段类型必须一致，适用于字符串、数字、数组和map
// argStr为数值时，数字比较大小，字符串、数组和map比较长度
func lt(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return compare(data, fieldName, argStr, func(v1, v2 float64) bool {
		return v1 < v2
	})
}

// 验证字段必须小于等于给定 field 字段，这两个字段类型必须一致，适用于字符串、数字、数组和map
// argStr为数值时，数字比较大小，字符串、数组和map比较长度
func lte(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return compare(data, fieldName, argStr, func(v1, v2 float64) bool {
		return v1 <= v2
	})
}

// 验证字段必须等于给定 field 字段，这两个字段类型必须一致，适用于字符串、数字、数组和map
// argStr为数值时，数字比较大小，字符串、数组和map比较长度
func size(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	return compare(data, fieldName, argStr, func(v1, v2 float64) bool {
		return v1 == v2
	})
}

// 验证字段必须有一个匹配字段 foo_confirmation，例如，如果验证字段是 password，必须输入一个与之匹配的 password_confirmation 字段
func confirmed(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		matchField string
	)
	if !ok {
		return
	}

	if argStr == "" {
		matchField = fieldName + "_confirmation"
	} else {
		matchField = argStr
	}

	matchValues, ok := data.Get(matchField)
	if !ok {
		return Error("", fieldName, _value, "")
	}

	for _, _value = range values {
		for _, matchValue := range matchValues {
			if _value != matchValue {
				return Error("", fieldName, _value, "")
			}
		}
	}
	return
}

// 验证字段必须是数值类型，并且必须包含指定的小数位数
func decimal(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok                 = data.Get(fieldName)
		_value                     any
		minLen, maxLen, decimalLen int
		err                        error
		valueFloat64               float64
		str                        string
		strArr                     []string
	)
	if !ok {
		return
	}

	if argStr != "" {
		args := strings.Split(argStr, ",")
		if len(args) > 1 {
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return
			}
			maxLen, err = strconv.Atoi(args[1])
			if err != nil {
				return
			}
		} else {
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return
			}
			maxLen = minLen
		}
		if minLen > maxLen {
			return Error("", fieldName, _value, "")
		}
	}

	for _, _value = range values {
		switch v := _value.(type) {
		case int:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case uint:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case int8:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case uint8:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case int32:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case uint32:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case int64:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case uint64:
			if minLen == 0 {
				return
			}
			valueFloat64 = float64(v)
		case float32:
			valueFloat64 = float64(v)
		case float64:
			valueFloat64 = v
		default:
			return Error("", fieldName, _value, "")
		}

		str = strconv.FormatFloat(valueFloat64, 'g', -1, 64)
		strArr = strings.Split(str, ".")
		if len(strArr) > 1 {
			decimalLen = len(strArr[1])
			if decimalLen < minLen || decimalLen > maxLen {
				return Error("", fieldName, _value, "")
			}
		} else {
			//没有小数部分
			if minLen > 0 {
				return Error("", fieldName, _value, "")
			}
		}
	}

	return
}

// 验证字段值必须在给定的列表中
func in(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		arg        string
	)
	if !ok {
		return
	}
	if argStr == "" {
		return Error("", fieldName, _value, "")
	}
	var (
		args     = strings.Split(argStr, ",")
		valueStr string
	)
	for _, _value = range values {
		valueStr = fmt.Sprintf("%v", _value)
		for _, arg = range args {
			if valueStr == arg {
				return
			}
		}
	}
	return Error("", fieldName, _value, "")
}

// 验证字段值不能在给定列表中
func notIn(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok = data.Get(fieldName)
		_value     any
		arg        string
	)
	if !ok || argStr == "" {
		return
	}
	var (
		args     = strings.Split(argStr, ",")
		valueStr string
	)
	for _, _value = range values {
		valueStr = fmt.Sprintf("%v", _value)
		for _, arg = range args {
			if valueStr == arg {
				return Error("", fieldName, _value, "")
			}
		}
	}
	return
}
