package test

import (
	"fmt"
	"github.com/fushiliang321/validator"
	"testing"
)

func Test_accepted(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"verifyEmail": true,
		}
		rules = map[string]string{
			"name":        "accepted",
			"verifyEmail": "accepted",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_accepted_if(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"sex":         1,
			"email":       "invalid-email@qq.com",
			"verifyEmail": true,
		}
		rules = map[string]string{
			"name":         "accepted_if:sex,1",
			"verifyEmail":  "accepted_if:sex,1",
			"verifyEmail1": "accepted_if:sex,1",
			"verifyEmail2": "accepted_if:sex,0",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_declined(t *testing.T) {
	var (
		data = map[string]any{
			"sex":         1,
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"verifyEmail": false,
		}
		rules = map[string]string{
			"name":        "declined",
			"verifyEmail": "declined",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_decline_if(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"name":         "declined_if:sex,1",
			"verifyEmail":  "declined_if:sex,1",
			"verifyEmail1": "declined_if:sex,1",
			"verifyEmail2": "declined_if:sex,0",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_boolean(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex":         "boolean",
			"verifyEmail": "boolean:strict",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_different(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"age":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex": "different:age",
		}
		res    = validator.Check(data, rules)
		rules1 = map[string]string{
			"sex": "different:name1",
		}
		res1 = validator.Check(data, rules1)
	)
	fmt.Println(res)
	fmt.Println(res1)
}

func Test_digits(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"age":         10,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex": "digits",
			"age": "digits:2",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_digits_between(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"age":         10,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex": "digits_between:0,2",
			"age": "digits_between:16,60",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}
