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
			"name|verifyEmail": "accepted",
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
			"name|verifyEmail|verifyEmail1": "accepted_if:sex,1",
			"verifyEmail2":                  "accepted_if:sex,0",
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
			"name|verifyEmail": "declined",
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
			"name|verifyEmail|verifyEmail1": "declined_if:sex,1",
			"verifyEmail2":                  "declined_if:sex,0",
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

func Test_lowercase(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"address":     "beijing",
			"sex":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex|verifyEmail|name|email|address": "lowercase",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_uppercase(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "JOHN",
			"email":       "invalid-email@qq.com",
			"address":     "beijing",
			"sex":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex|verifyEmail|name|email|address": "uppercase",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_ip(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "JOHN",
			"email":       "invalid-email@qq.com",
			"ip":          "192.168.31.1",
			"sex":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"name|ip": "ip|ipv4",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_ipv6(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "JOHN",
			"email":       "invalid-email@qq.com",
			"ip":          "192.168.31.1",
			"ipv6":        "240e:379:ad93:1f00:a66f:801b:4454:c6ba",
			"sex":         1,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"name|ip|ipv6": "ipv6",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_json(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "JOHN",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"verifyEmail": false,
			"data":        "{\"a\":\"xxx\",\"b\":1}",
			"data1":       "{\"a\":\"xxx\",\"b\":1",
		}
		rules = map[string]string{
			"name|data|data1": "json",
		}
		res = validator.Check(data, rules)
	)
	fmt.Println(res)
}

func Test_after(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
		}
		rules = map[string]string{
			"register_time|graduate_time": "after:2024-02-02",
		}
		rules1 = map[string]string{
			"register_time|graduate_time": "after:tomorrow",
		}
		rules2 = map[string]string{
			"register_time|graduate_time": "after:birthday",
		}
	)
	fmt.Println(validator.Check(data, rules))
	fmt.Println(validator.Check(data, rules1))
	fmt.Println(validator.Check(data, rules2))
}

func Test_after_or_equal(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
		}
		rules = map[string]string{
			"register_time|graduate_time|birthday": "after_or_equal:2024-12-01",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_before(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
		}
		rules = map[string]string{
			"register_time|graduate_time": "before:2024-02-02",
		}
		rules1 = map[string]string{
			"register_time|graduate_time": "before:tomorrow",
		}
		rules2 = map[string]string{
			"register_time|graduate_time": "before:birthday",
		}
	)
	fmt.Println(validator.Check(data, rules))
	fmt.Println(validator.Check(data, rules1))
	fmt.Println(validator.Check(data, rules2))
}

func Test_before_or_equal(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
		}
		rules = map[string]string{
			"register_time|graduate_time|birthday": "after_or_equal:2024-12-01",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_size(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"data":          []string{"a", "b", "C", ""},
			"data1":         []string{"a", "b", "C"},
			"data2":         []string{"a", "b", "C", ""},
		}
		rules = map[string]string{
			"name|birthday|data":     "size:4",
			"register_time|birthday": "size:graduate_time",
			"data|data1":             "size:data2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_gt(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"data":          []string{"a", "b", "C", ""},
			"data1":         []string{"a", "b", "c"},
			"data2":         []string{"a", "b", "C"},
		}
		rules = map[string]string{
			"name|birthday|data":     "gt:5",
			"register_time|birthday": "gt:birthday",
			"data|data1":             "gt:data2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}
func Test_gte(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"data":          []string{"a", "b", "C", "d", "e"},
			"data1":         []string{"a", "b", "c"},
			"data2":         []string{"a", "b", "C", "d", "e"},
		}
		rules = map[string]string{
			"name|birthday|data":     "gte:5",
			"register_time|birthday": "gte:graduate_time",
			"data|data1":             "gte:data2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_lt(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"data":          []string{"a", "b", "C", "d"},
			"data1":         []string{"a", "b", "c"},
			"data2":         []string{"a", "b", "C", "d"},
		}
		rules = map[string]string{
			"name|birthday|data":     "lt:5",
			"register_time|birthday": "lt:graduate_time",
			"data|data1":             "lt:data2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}
func Test_lte(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"sex":           1,
			"register_time": "2024-02-01 12:12:12",
			"graduate_time": "2034-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"data":          []string{"a", "b", "C", "d", "e"},
			"data1":         []string{"a", "b", "c"},
			"data2":         []string{"a", "b", "C"},
		}
		rules = map[string]string{
			"name|birthday|data":     "lte:5",
			"register_time|birthday": "lte:graduate_time",
			"data|data1":             "lte:data2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_in(t *testing.T) {
	var (
		data = map[string]any{
			"name": "JOHN",
			"sex":  1,
		}
		rules = map[string]string{
			"name|sex": "in:0,1,2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_not_in(t *testing.T) {
	var (
		data = map[string]any{
			"name": "JOHN",
			"sex":  1,
		}
		rules = map[string]string{
			"name|sex": "not_in:0,1,2",
		}
	)
	fmt.Println(validator.Check(data, rules))
}
