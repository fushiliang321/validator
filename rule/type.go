package rule

import (
	"fmt"
	"github.com/fushiliang321/validator/utils"
	"github.com/fushiliang321/validator/value"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func init() {
	Register("digits", digits)
	Register("digits_between", digitsBetween)
	Register("url", _url)
	Register("date", date)
	Register("string", _string)
}

// 验证字段必须是数字
// argStr存在且为数字时验证字段长度必须为argStr指定的长度值
func digits(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		str        string
		values, ok = data.Get(fieldName)
		_value     any
		isFloat    bool
	)
	if !ok {
		return
	}
	for _, _value = range values {
		isFloat = false
		switch _value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32:
		case uint64, float32:
			isFloat = true
		default:
			return Error("", fieldName, _value, "")
		}
		_len, _ := strconv.Atoi(argStr)
		if _len > 0 {
			//需要判断长度
			str = fmt.Sprintf("%v", _value)
			if isFloat {
				if strings.Index(str, ".") != -1 {
					//有小数点
					_len++
				}
			}
			if len(str) != _len {
				return Error("", fieldName, _value, "")
			}
		}
	}
	return
}

// 验证字段数值长度必须介于最小值和最大值之间
func digitsBetween(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	args := strings.Split(argStr, ",")
	if len(args) < 2 {
		return
	}

	var (
		minStr     = args[0]
		maxStr     = args[1]
		minI       float64
		maxI       float64
		toFloat64  float64
		err        error
		_value     any
		values, ok = data.Get(fieldName)
	)
	if !ok {
		return
	}
	minI, err = strconv.ParseFloat(minStr, 64)
	if err != nil {
		return
	}
	maxI, err = strconv.ParseFloat(maxStr, 64)
	if err != nil {
		return
	}
	if minI >= maxI {
		return
	}

	for _, _value = range values {
		toFloat64, err = utils.AnyToFloat64(_value)
		if err != nil {
			return Error("", fieldName, nil, "")
		}
		if toFloat64 < minI || toFloat64 > maxI {
			return Error("", fieldName, nil, "")
		}
	}

	return
}

// 判断是否为url地址
func _url(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	values, ok := data.Get(fieldName)
	if !ok {
		return
	}

	var (
		err    error
		_value any
	)

	for _, _value = range values {
		switch v := _value.(type) {
		case string:
			_, err = url.Parse(v)
			if err != nil {
				_, err = url.Parse("https://" + v)
				if err != nil {
					return Error("", fieldName, v, "")
				}
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
