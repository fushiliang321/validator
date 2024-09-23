package value

import "reflect"

type Value struct {
	Raw          any
	reflectType  *reflect.Type
	reflectValue *reflect.Value
}

func (v *Value) TypeOf() reflect.Type {
	if v.reflectType == nil {
		var typeOf reflect.Type
		if v.reflectValue == nil {
			typeOf = reflect.TypeOf(v.Raw)
		} else {
			typeOf = v.reflectValue.Type()
		}
		v.reflectType = &typeOf
	}
	return *v.reflectType
}

func (v *Value) ValueOf() reflect.Value {
	if v.reflectValue == nil {
		valueOf := reflect.ValueOf(v.Raw)
		v.reflectValue = &valueOf
	}
	return *v.reflectValue
}

func Transition(data map[string]any) *Data {
	return &Data{
		Raw: data,
	}
}
