package test

import (
	"fmt"
	"github.com/fushiliang321/validator"
	"testing"
)

func Test_alpha(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"address":       "XXX小区A区",
			"register_time": "2024-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"sex":           1,
		}
		rules = map[string]string{
			"name|sex|address|register_time|birthday": "alpha",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_alpha_num(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"address":       "xxx小区A区12栋",
			"register_time": "2024-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"sex":           1,
		}
		rules = map[string]string{
			"name|sex|address|register_time|birthday": "alpha_num",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_alpha_dash(t *testing.T) {
	var (
		data = map[string]any{
			"name":          "JOHN",
			"address":       "xxx小区A区12栋",
			"register_time": "2024-02-01 12:12:12",
			"birthday":      "2024-12-01",
			"sex":           1,
		}
		rules = map[string]string{
			"name|sex|address|register_time|birthday": "alpha_dash",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_mac_address(t *testing.T) {
	var (
		data = map[string]any{
			"name":        "john",
			"sex":         1,
			"mac_address": "D8-BB-C1-35-CF-3A",
		}
		rules = map[string]string{
			"name|sex|mac_address": "mac_address",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_email(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "john",
			"sex":   1,
			"email": "invalid-email@qq.com",
		}
		rules = map[string]string{
			"name|sex|email": "email",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_phone(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "john",
			"sex":   1,
			"phone": "15559144999",
		}
		rules = map[string]string{
			"name|sex|phone": "phone",
		}
	)
	fmt.Println(validator.Check(data, rules))
}

func Test_regex(t *testing.T) {
	var (
		data = map[string]any{
			"name":  "john",
			"sex":   1,
			"phone": "15559144999",
		}
		rules = map[string]string{
			"name|sex|phone": `regex:/H/i`,
		}
	)
	fmt.Println(validator.Check(data, rules))
}
