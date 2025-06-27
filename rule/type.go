package rule

import (
	json2 "encoding/json"
	"fmt"
	"github.com/fushiliang321/validator/utils"
	"github.com/fushiliang321/validator/value"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func init() {
	Register("integer", integer)
	Register("url", _url)
	Register("date", date)
	Register("string", _string)
	Register("array", array)
	Register("object", object)
}

// 验证字段是否为整型（值可以是整型/整型字符串，strict模式值必须为整型）
// argStr存在且为数字时验证字段长度必须为argStr指定的长度值
func integer(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok               = data.Get(fieldName)
		_value                   any
		valueStr                 string
		minLen, maxLen, valueLen int
		err                      error
		isStrict                 = false
		isVerifyLen              = false //是否验证长度
	)
	if !ok {
		return
	}

	if argStr != "" {
		args := strings.Split(argStr, ",")
		if args[0] == "strict" {
			//严格模式
			isStrict = true
			args = args[1:]
		}
		switch len(args) {
		case 0:
		case 1:
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return Error("", fieldName, values[0], "")
			}
			maxLen = minLen
			isVerifyLen = true
		default:
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return Error("", fieldName, values[0], "")
			}
			maxLen, err = strconv.Atoi(args[1])
			if err != nil {
				return Error("", fieldName, values[0], "")
			}
			if minLen > maxLen {
				return Error("", fieldName, values[0], "")
			}
			isVerifyLen = true
		}
	}

	for _, _value = range values {
		switch _value.(type) {
		case float64, float32:
			if isStrict {
				return Error("", fieldName, _value, "")
			}
			valueStr = fmt.Sprint(_value)
			if strings.IndexAny(valueStr, ".") != -1 {
				return Error("", fieldName, _value, "")
			}
		case uint64, uint32, int, int8, int16, int32, int64, uint, uint8, uint16:
			valueStr = fmt.Sprint(_value)
		case json2.Number:
			var v any
			if isStrict {
				v, err = _value.(json2.Number).Int64()
				if err != nil {
					return Error("", fieldName, _value, "")
				}
				valueStr = fmt.Sprint(v)
			} else {
				v, err = _value.(json2.Number).Float64()
				if err != nil {
					return Error("", fieldName, _value, "")
				}
				valueStr = fmt.Sprint(v)
				if strings.IndexAny(valueStr, ".") != -1 {
					return Error("", fieldName, _value, "")
				}
			}
		case string:
			if isStrict {
				return Error("", fieldName, _value, "")
			}
			i, err := strconv.Atoi(_value.(string))
			if err == nil {
				return Error("", fieldName, _value, "")
			}
			valueStr = fmt.Sprint(i)
		default:
			if isStrict {
				return Error("", fieldName, _value, "")
			}
			i, err := strconv.Atoi(fmt.Sprintf("%v", _value))
			if err == nil {
				return Error("", fieldName, _value, "")
			}
			valueStr = fmt.Sprint(i)
		}
		if !isVerifyLen {
			continue
		}
		valueLen = len(valueStr)
		if valueLen < minLen || valueLen > maxLen {
			return Error("", fieldName, _value, "")
		}
	}
	return
}

// 验证字段是否是数值类型（可以为数值/数值字符串，strict模式必须为数值类型），并且必须包含指定的小数位数；
func decimal(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok                 = data.Get(fieldName)
		_value                     any
		minLen, maxLen, decimalLen int
		err                        error
		valueFloat64               float64
		str                        string
		strArr                     []string
		isStrict                   = false
		isVerifyLen                = false //是否验证小数位数
	)
	if !ok {
		return
	}

	if argStr != "" {
		args := strings.Split(argStr, ",")
		if args[0] == "strict" {
			//严格模式
			isStrict = true
			args = args[1:]
		}
		switch len(args) {
		case 0:
		case 1:
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return
			}
			maxLen = minLen
			isVerifyLen = true
		default:
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return
			}
			maxLen, err = strconv.Atoi(args[1])
			if err != nil {
				return
			}
			if minLen > maxLen {
				return Error("", fieldName, values[0], "")
			}
			isVerifyLen = true
		}
	}

	for _, _value = range values {
		if minLen > 0 {
			switch _value.(type) {
			case float64, float32, json2.Number, string:
				//minLen > 0的时候只有浮点型才能通过验证，整型转换后小数位只会为0
			default:
				return Error("", fieldName, _value, "")
			}
		}
		valueFloat64, err = utils.AnyToFloat64(_value, isStrict)
		if err != nil {
			return Error("", fieldName, _value, "")
		}

		if !isVerifyLen {
			continue
		}

		str = strconv.FormatFloat(valueFloat64, 'f', -1, 64)
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

// 验证字段值是否为url地址
func _url(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}

	var (
		_value any
	)

	for _, _value = range values {
		switch v := _value.(type) {
		case string:
			u, err := url.Parse(v)
			if err != nil {
				return Error("", fieldName, v, "")
			}
			if argStr == "" {
				if u.Scheme == "" {
					return Error("", fieldName, v, "")
				}
			} else if u.Scheme != argStr {
				return Error("", fieldName, v, "")
			}
		default:
			return Error("", fieldName, v, "")
		}
	}
	return
}

// 验证字段值是否为指定日期格式
// 默认格式：YYYY-MM-DD HH:mm:ss
func date(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}
	if argStr == "" {
		argStr = time.DateTime
	}
	var (
		err    error
		layout string
		_value any
	)

	for _, _value = range values {
		switch v := _value.(type) {
		case string:
			layout = utils.ConvertDateLayout(argStr)
			_, err = time.Parse(layout, v)
			if err != nil {
				return Error("", fieldName, v, "")
			}
		default:
			return Error("", fieldName, v, "")
		}
	}

	return
}

// 验证字段必须是字符串
func _string(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}
	for _, _value := range values {
		switch _value.(type) {
		case string:
		default:
			return Error("", fieldName, _value, "")
		}
	}
	return
}

// 验证字段必须可以转换成数组
func array(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}

	for _, _value := range values {
		if !utils.IsArrayOrSlice(_value) {
			return Error("", fieldName, _value, "")
		}
	}

	return
}

// 验证字段必须可以转换成对象
func object(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}

	for _, _value := range values {
		if !utils.IsObject(_value) {
			return Error("", fieldName, _value, "")
		}
	}
	return
}
