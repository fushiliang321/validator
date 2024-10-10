package test

import (
	"fmt"
	"github.com/fushiliang321/validator"
	"testing"
)

func Test_required(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"name|name1": "required",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}
func Test_required_if(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"sex":         1,
			"verifyEmail": true,
			"email":       "invalid-email@qq.com",
			"verifyAge":   true,
		}
		rules = map[string]string{
			"name|name1": "required_if:sex,1",
			"name2":      "required_if:sex,0",
			"email":      "required_if:sex,1,verifyEmail,true",
			"age":        "required_if:sex,1,verifyAge,true",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_required_unless(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"sex":         1,
			"verifyEmail": true,
			"email":       "invalid-email@qq.com",
			"verifyAge":   true,
		}
		rules = map[string]string{
			"name|name1": "required_if:sex,1",
			"name2":      "required_if:sex,0",
			"email":      "required_unless:sex,1,verifyEmail,true",
			"age":        "required_if:sex,1,verifyAge,true",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_required_with(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"name": "required_with:sex",
			"sex1": "required_with:sex2,sex",
			"sex2": "required_with:sex1",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_required_with_all(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"email|sex1": "required_with_all:name,sex",
			"sex2":       "required_with_all:name,sex1",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_required_without(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"email|sex1": "required_without:name,sex",
			"sex2":       "required_without:name,sex1",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_required_without_all(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"email|sex1": "required_without_all:name,sex",
			"sex2":       "required_without_all:name,sex1",
			"sex3":       "required_without_all:name1,sex1",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_prohibited(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"email|name|name1": "prohibited",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_prohibited_if(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "xxx1",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"email": "prohibited_if:sex,1",
			"name":  "prohibited_if:sex,0",
			"sex":   "prohibited_if:name,xxx",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_missing(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"email|name|name1": "missing",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_missing_if(t *testing.T) {
	var (
		data = map[string]any{
			"name":    "xxx1",
			"sex":     1,
			"email":   "invalid-email@qq.com",
			"address": "",
		}
		rules = map[string]string{
			"email|email1|address": "missing_if:sex,1",
			"name":                 "missing_if:sex,0",
			"sex":                  "missing_if:name,xxx",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_missing_unless(t *testing.T) {
	var (
		data = map[string]any{
			"name":    "xxx1",
			"sex":     1,
			"email":   "invalid-email@qq.com",
			"address": "",
		}
		rules = map[string]string{
			"email":          "missing_unless:sex,1,name,xxx",
			"email1|address": "missing_unless:sex,0",
			"name":           "missing_unless:sex,0,name,xxx1",
			"sex":            "missing_unless:name,xxx",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_missing_with(t *testing.T) {
	var (
		data = map[string]any{
			"name":    "xxx1",
			"sex":     1,
			"email":   "invalid-email@qq.com",
			"address": "",
		}
		rules = map[string]string{
			"email|email1": "missing_with:sex,1",
			"name":         "missing_with:sex1,0",
			"sex":          "missing_with:name,xxx",
			"address":      "missing_with:sex,1",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_filled(t *testing.T) {
	var (
		data = map[string]any{
			"name":    "xxx1",
			"sex":     1,
			"address": "",
		}
		rules = map[string]string{
			"name|email|address|age": "filled",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}
