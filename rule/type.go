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
	Register("integer", integer)
	Register("url", _url)
	Register("date", date)
	Register("string", _string)
	Register("array", array)
	Register("object", object)
}

// 验证字段必须是整型
// argStr存在且为数字时验证字段长度必须为argStr指定的长度值
func integer(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		values, ok               = data.Get(fieldName)
		_value                   any
		minLen, maxLen, valueLen int
		err                      error
	)
	if !ok {
		return
	}

	if argStr != "" {
		args := strings.Split(argStr, ",")

		if len(args) > 1 {
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return Error("", fieldName, _value, "")
			}
			maxLen, err = strconv.Atoi(args[1])
			if err != nil {
				return Error("", fieldName, _value, "")
			}
			if minLen > maxLen {
				return Error("", fieldName, _value, "")
			}
		} else {
			minLen, err = strconv.Atoi(args[0])
			if err != nil {
				return Error("", fieldName, _value, "")
			}
			maxLen = minLen
		}
	}

	for _, _value = range values {
		switch _value.(type) {
		case float64, float32:
			if strings.IndexAny(fmt.Sprint(_value), ".") != -1 {
				return Error("", fieldName, _value, "")
			}
		case uint64, uint32, int, int8, int16, int32, int64, uint, uint8, uint16:
		default:
			return Error("", fieldName, _value, "")
		}
		if argStr == "" {
			continue
		}
		valueLen = len(fmt.Sprintf("%v", _value))
		if valueLen < minLen || valueLen > maxLen {
			return Error("", fieldName, _value, "")
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
