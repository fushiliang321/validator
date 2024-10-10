package test

import (
	"fmt"
	"github.com/fushiliang321/validator"
	"testing"
)

func Test_integer(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "John",
			"email":       "invalid-email@qq.com",
			"sex":         1,
			"age":         10,
			"height":      180,
			"verifyEmail": false,
		}
		rules = map[string]string{
			"sex":    "integer",
			"age":    "integer:2",
			"height": "integer:2,3",
		}
		rules2 = map[string]string{
			"sex":   "integer:2,3",
			"age":   "integer:1",
			"email": "integer:2,3",
		}
	)
	fmt.Println(validator.Check(data, rules))
	fmt.Println(validator.Check(data, rules2))
}

func Test_url(t *testing.T) {
	var (
		data = map[string]any{
			"name":     "John",
			"email":    "invalid-email@qq.com",
			"sex":      1,
			"age":      10,
			"website":  "https://e.gitee.com",
			"website1": "https://e.gitee.com/",
			"website2": "http://e.gitee.com/",
		}
		rules = map[string]string{
			"website|email":     "url",
			"website1|website2": "url:https",
		}
	)
	fmt.Println(validator.Check(data, rules))
}
func Test_date(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "John",
			"email":         "invalid-email@qq.com",
			"sex":           1,
			"age":           10,
			"register_time": "2024-02-01 12:12:12",
			"birthday":      "2024-12-01",
		}
		rules = map[string]string{
			"register_time|email": "date",
			"birthday":            "date:YYYY-MM-DD",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_string(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "John",
			"email":         "invalid-email@qq.com",
			"sex":           1,
			"age":           10,
			"register_time": "2024-02-01 12:12:12",
			"birthday":      "2024-12-01",
		}
		rules = map[string]string{
			"register_time|email|sex": "string",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_array(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"email": "invalid-email@qq.com",
			"sex":   1,
			"age":   10,
			"hobby": []string{"xxx", "xxx2"},
			"home": map[string]any{
				"city": "beijing",
			},
		}
		rules = map[string]string{
			"hobby|sex|home": "array",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_object(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "John",
			"email": "invalid-email@qq.com",
			"sex":   1,
			"age":   10,
			"hobby": []string{"xxx", "xxx2"},
			"home": map[string]any{
				"city": "beijing",
			},
		}
		rules = map[string]string{
			"hobby|sex|home": "object",
		}
	)
	fmt.Println(validator.Check(data, rules))
}
