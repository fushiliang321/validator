package rule

import (
	"github.com/fushiliang321/validator/utils"
	"github.com/fushiliang321/validator/value"
	"reflect"
	"strings"
)

func init() {
	Register("required", required)
	Register("required_if", requiredIf)
	Register("required_unless", requiredUnless)
	Register("required_with", requiredWith)
	Register("required_with_all", requiredWithAll)
	Register("required_without", requiredWithout)
	Register("required_without_all", requiredWithoutAll)
	Register("prohibited", prohibited)
	Register("prohibited_if", prohibitedIf)
	Register("missing", missing)
	Register("missing_if", missingIf)
	Register("missing_unless", missingUnless)
	Register("missing_with", missingWith)
	Register("missing_with_all", missingWithAll)
	Register("filled", filled)
}

// 验证字段值不能为空，以下情况字段值都为空： 值为null、值是空字符串、值是空数组或者空的对象
func required(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	var (
		v          any
		values, ok = data.Get(fieldName)
	)
	if !ok {
		return Error("", fieldName, nil, "")
	}

	for _, v = range values {
		if v == nil || v == "" {
			//空值或者空字符串
			return Error("", fieldName, v, "")
		}

		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
			if reflect.ValueOf(v).IsZero() {
				//空数组或者空的对象
				return Error("", fieldName, v, "")
			}
		default:
		}
	}
	return
}

// 验证字段在指定字段等于指定值时必须存在且不能为空
func requiredIf(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if !utils.FieldAllEqValue(data, argStr) {
		return
	}
	return required(data, fieldName, "")
}

// 除非指定字段等于指定值，否则验证字段必须存在且不能为空
func requiredUnless(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if !utils.FieldEqValue(data, argStr) {
		return
	}
	return required(data, fieldName, "")
}

// 验证字段只有在任一其它指定字段存在的情况下才是必须的
func requiredWith(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	var (
		args = strings.Split(argStr, ",")
		arg  string
		ok   bool
	)
	for _, arg = range args {
		if _, ok = data.Get(arg); ok {
			return required(data, fieldName, "")
		}
	}
	return
}

// 验证字段只有在所有指定字段都存在的情况下才是必须的
func requiredWithAll(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	var (
		args = strings.Split(argStr, ",")
		arg  string
		ok   bool
	)
	for _, arg = range args {
		if _, ok = data.Get(arg); !ok {
			return
		}
	}
	return required(data, fieldName, "")
}

// 验证字段只有当任一指定字段不存在的情况下才是必须的
func requiredWithout(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	var (
		args = strings.Split(argStr, ",")
		arg  string
		ok   bool
	)
	for _, arg = range args {
		if _, ok = data.Get(arg); !ok {
			return required(data, fieldName, "")
		}
	}
	return
}

// 验证字段只有当所有指定字段都不存在的情况下才是必须的。
func requiredWithoutAll(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	var (
		args = strings.Split(argStr, ",")
		arg  string
		ok   bool
	)
	for _, arg = range args {
		if _, ok = data.Get(arg); ok {
			return
		}
	}
	return required(data, fieldName, "")
}

// 需要验证的字段必须不存在或为空。以下情况字段值都为空： 值为null、值是空字符串、值是空数组或者空的对象
func prohibited(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if required(data, fieldName, "") == nil {
		values, _ := data.Get(fieldName)
		return Error("", fieldName, values, "")
	}
	return
}

// 验证字段在指定字段等于指定值时必须不存在或为空
func prohibitedIf(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" || !utils.FieldEqValue(data, argStr) {
		return
	}
	return prohibited(data, fieldName, "")
}

// 验证的字段在输入数据中必须不存在
func missing(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if vals, ok := data.Get(fieldName); ok {
		return Error("", fieldName, vals, "")
	}
	return
}

// 如果有指定字段等于指定值，则验证的字段必须不存在
func missingIf(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" || !utils.FieldEqValue(data, argStr) {
		return
	}
	return missing(data, fieldName, "")
}

// 验证字段必须不存在，除非指定字段等于指定值
func missingUnless(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if utils.FieldEqValue(data, argStr) {
		return
	}
	return missing(data, fieldName, "")
}

// 验证字段只有在任一其它指定字段存在，则验证的字段必须不存在
func missingWith(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	var (
		args = strings.Split(argStr, ",")
		arg  string
		ok   bool
	)
	for _, arg = range args {
		if _, ok = data.Get(arg); ok {
			return missing(data, fieldName, "")
		}
	}
	return
}

// 如果所有其他指定的字段都存在，则验证的字段必须不存在
func missingWithAll(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	if argStr == "" {
		return
	}
	var (
		args = strings.Split(argStr, ",")
		arg  string
		ok   bool
	)
	for _, arg = range args {
		if _, ok = data.Get(arg); !ok {
			return
		}
	}
	return missing(data, fieldName, "")
}

// 验证字段如果存在则不能为空
func filled(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
	_, ok := data.Get(fieldName)
	if !ok {
		return
	}
	return required(data, fieldName, "")
}
