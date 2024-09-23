package value

import (
	"strconv"
	"strings"
)

type Data struct {
	Raw map[string]any
}

func get(_map map[string]any, keyStr string) ([]any, bool) {
	if v, ok := _map[keyStr]; ok {
		return []any{v}, ok
	}

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

// 获取指定字段的值
// 支持获取子字段的值，字段之间用.分隔，*表示任意字段，例：key.subKey ; key.*.subKey ; key.1.subKey
func (d *Data) Get(key string) ([]any, bool) {
	return get(d.Raw, key)
}
