package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func tt(data any) {
	switch expr := data.(type) {
	case map[string]any:
		fmt.Println(expr)
		if expr["name"] != nil {
			tt(expr["name"])
		}
	}
}

func get(_map map[string]any, keyStr string) ([]any, bool) {
	//if v, ok := _map[keyStr]; ok {
	//	return v, ok
	//}

	var (
		keys         = strings.Split(keyStr, ".")
		keysMaxIndex = len(keys) - 1
	)
	if keysMaxIndex < 0 {
		return nil, false
	}
	var (
		values    = []any{_map}
		newValues []any
		key       string
	)
	for _, key = range keys {
		newValues = []any{}
		for _, value := range values {
			switch v := value.(type) {
			case map[string]any:
				if key == "*" {
					for _, v1 := range v {
						newValues = append(newValues, v1)
					}
					break
				}
				if v1, ok := v[key]; ok {
					newValues = append(newValues, v1)
				} else {
					break
				}
			case []any:
				if key == "*" {
					for _, v1 := range v {
						newValues = append(newValues, v1)
					}
					break
				}
				index, err := strconv.Atoi(key)
				if err != nil || len(v) <= index {
					break
				}
				newValues = append(newValues, v[index])
			}
		}
		values = newValues
	}
	if len(values) < 1 {
		return nil, false
	}

	return values, true
}

func Test(t *testing.T) {
	var (
		data = map[string]any{
			"name": "test",
			"age":  18,
			"list": []any{
				map[string]any{
					"name": "test1",
					"age":  1,
				},
				map[string]any{
					"name": "test2",
					"age":  1,
				},
			},
			"info": map[string]any{
				"name": "test3",
				"age":  2,
				"list": []any{
					map[string]any{
						"name": "test4",
						"age":  3,
					},
					map[string]any{
						"name": "test5",
						"age":  1,
					},
				},
			},
		}

		keys = []string{
			//"name",
			"*.*.*.name",
			//"info.name",
		}
	)

	for _, key := range keys {
		fmt.Println(get(data, key))
	}
}
