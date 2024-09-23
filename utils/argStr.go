package utils

import (
	"encoding/json"
	"errors"
	"github.com/fushiliang321/validator/value"
	"github.com/savsgio/gotils/strconv"
	"strings"
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
			continue
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
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func AnyToFloat64(i any) (float64, error) {
	switch raw := i.(type) {
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
	case float32:
		return float64(raw), nil
	case float64:
		return raw, nil
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

func ConvertDateLayout(layout string) string {
	goLayout := layout

	// 替换自定义格式为 Go 格式
	for k, v := range replacements {
		goLayout = strings.ReplaceAll(goLayout, k, v)
	}

	return goLayout
}
