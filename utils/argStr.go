package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fushiliang321/validator/value"
	"github.com/savsgio/gotils/strconv"
	strconv2 "strconv"
	"strings"
	"time"
)

// FieldAllEqValue 判断字段是否全部等于指定值
func FieldAllEqValue(data *value.Data, argStr string) bool {
	if argStr == "" {
		return true
	}
	var (
		args              = strings.Split(argStr, ",")
		i                 int
		_len              = len(args)
		anotherField      string
		anotherFieldValue string
		values            []any
		fieldValueJson    []byte
		fieldValueStr     string
		err               error
		ok                bool
	)
	for i = 0; i < _len; i++ {
		anotherField = args[i]
		i++
		if i >= _len {
			return false
		}

		anotherFieldValue = args[i]
		values, ok = data.Get(anotherField)
		if !ok {
			return false
		}
		for _, v := range values {
			if v == anotherFieldValue {
				continue
			}
			fieldValueStr, ok = v.(string)
			if !ok {
				fieldValueJson, err = json.Marshal(v)
				if err != nil {
					return false
				}
				fieldValueStr = strconv.B2S(fieldValueJson)
			}
			if fieldValueStr != anotherFieldValue {
				return false
			}
		}
	}
	return true
}

// FieldEqValue 判断是否有字段等于指定值
func FieldEqValue(data *value.Data, argStr string) bool {
	if argStr == "" {
		return true
	}
	var (
		args              = strings.Split(argStr, ",")
		i                 int
		_len              = len(args)
		anotherField      string
		anotherFieldValue string
		values            []any
		fieldValueJson    []byte
		fieldValueStr     string
		err               error
		ok                bool
	)
	for i = 0; i < _len; i++ {
		anotherField = args[i]
		i++
		if i >= _len {
			return false
		}
		anotherFieldValue = args[i]
		values, ok = data.Get(anotherField)
		if !ok {
			continue
		}
		for _, v := range values {
			if v == anotherFieldValue {
				return true
			}

			fieldValueStr, ok = v.(string)
			if !ok {
				fieldValueJson, err = json.Marshal(v)
				if err != nil {
					continue
				}
				fieldValueStr = strconv.B2S(fieldValueJson)
			}
			if fieldValueStr != anotherFieldValue {
				continue
			}
			return true
		}
	}
	return false
}

func IsNumber(v interface{}) bool {
	switch v.(type) {
	case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

func AnyToFloat64(i any) (float64, error) {
	switch raw := i.(type) {
	case float64:
		return raw, nil
	case float32:
		return float64(raw), nil
	case int:
		return float64(raw), nil
	case int8:
		return float64(raw), nil
	case int16:
		return float64(raw), nil
	case int32:
		return float64(raw), nil
	case int64:
		return float64(raw), nil
	case uint:
		return float64(raw), nil
	case uint8:
		return float64(raw), nil
	case uint16:
		return float64(raw), nil
	case uint32:
		return float64(raw), nil
	case uint64:
		return float64(raw), nil
	default:
		return 0, errors.New("not a number")
	}
}

var replacements = map[string]string{
	"YYYY": "2006",
	"MM":   "01",
	"DD":   "02",
	"HH":   "15",
	"mm":   "04",
	"ss":   "05",
}

// 将日期格式转换为 Go 支持的日期格式
func ConvertDateLayout(layout string) string {
	goLayout := layout

	// 替换自定义格式为 Go 格式
	for k, v := range replacements {
		goLayout = strings.ReplaceAll(goLayout, k, v)
	}

	return goLayout
}

// 日期字符串转换为时间对象
func ParseDate(dateStr string) (time.Time, error) {
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"2006/01/02",
		"02 Jan 2006",
		"Jan 2, 2006",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"02 Jan 2006 15:04:05",
	}

	for _, format := range formats {
		date, err := time.Parse(format, dateStr)
		if err == nil {
			return date, nil // 找到匹配的格式，返回解析后的时间
		}
	}
	return time.Time{}, fmt.Errorf("无法解析日期: %s", dateStr) // 所有格式都未能匹配
}

// 字符串转换为时间对象
func StrToTime(str string) (time.Time, error) {
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"2006/01/02",
		"02 Jan 2006",
		"Jan 2, 2006",
		"15:04:05",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"02 Jan 2006 15:04:05",
	}

	for _, format := range formats {
		date, err := time.Parse(format, str)
		if err == nil {
			return date, nil // 找到匹配的格式，返回解析后的时间
		}
	}
	args := strings.Split(str, ",")
	if len(args) == 3 {
		//years,months,days
		var (
			year, _  = strconv2.Atoi(args[0])
			month, _ = strconv2.Atoi(args[1])
			day, _   = strconv2.Atoi(args[2])
		)
		return time.Now().AddDate(year, month, day), nil
	}
	var (
		_time = time.Now()
	)
	switch args[0] {
	case "now": //当前
		return _time, nil
	case "today": //今天
		return time.Date(_time.Year(), _time.Month(), _time.Day(), 0, 0, 0, 0, _time.Location()), nil
	case "tomorrow": //明天
		_time = _time.AddDate(0, 0, 1)
		return time.Date(_time.Year(), _time.Month(), _time.Day(), 0, 0, 0, 0, _time.Location()), nil
	case "yesterday": //昨天
		_time = _time.AddDate(0, 0, -1)
		return time.Date(_time.Year(), _time.Month(), _time.Day(), 0, 0, 0, 0, _time.Location()), nil
	case "week": //本周起始
		var _, week = _time.ISOWeek()
		return time.Date(_time.Year(), time.January, (week-1)*7, 0, 0, 0, 0, _time.Location()), nil
	case "lastweek": //上周起始
		var _, week = _time.ISOWeek()
		return time.Date(_time.Year(), time.January, (week-2)*7, 0, 0, 0, 0, _time.Location()), nil
	case "nextweek": //下周起始
		var _, week = _time.ISOWeek()
		return time.Date(_time.Year(), time.January, week*7, 0, 0, 0, 0, _time.Location()), nil
	case "month": //本月起始
		return time.Date(_time.Year(), _time.Month(), 1, 0, 0, 0, 0, _time.Location()), nil
	case "lastmonth": //上个月起始
		_time = _time.AddDate(0, -1, 0)
		return time.Date(_time.Year(), _time.Month(), 1, 0, 0, 0, 0, _time.Location()), nil
	case "nextmonth": //下个月起始
		_time = _time.AddDate(0, 1, 0)
		return time.Date(_time.Year(), _time.Month(), 1, 0, 0, 0, 0, _time.Location()), nil
	case "year": //本年起始
		return time.Date(_time.Year(), time.January, 1, 0, 0, 0, 0, _time.Location()), nil
	case "lastyear": //去月起始
		return time.Date(_time.Year()-1, time.January, 1, 0, 0, 0, 0, _time.Location()), nil
	case "nextyear": //明月起始
		return time.Date(_time.Year()+1, time.January, 1, 0, 0, 0, 0, _time.Location()), nil
	}

	return time.Time{}, fmt.Errorf("无法解析时间: %s", str) // 所有格式都未能匹配
}
