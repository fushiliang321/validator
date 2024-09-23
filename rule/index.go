package rule

import (
	"github.com/fushiliang321/validator/value"
)

type (
	Name    = string //校验规则名称
	Field   = string //校验字段名称
	RuleStr = string //校验规则字符串

	CheckFunc func(data *value.Data, fieldName Field, str RuleStr) *CheckError

	CheckError struct {
		Field    Field
		FieldVal any
		RuleName Name
		errMsg   string
	}

	RuleMap struct {
		_map map[Name]CheckFunc
	}
)

var GlobalRules = New()

func (e *CheckError) Error() string {
	if e.errMsg == "" {
		e.errMsg = e.RuleName + ":" + e.Field
	}
	return e.errMsg
}

func Error(ruleName Name, fieldName Field, val any, msg string) *CheckError {
	return &CheckError{
		Field:    fieldName,
		FieldVal: val,
		RuleName: ruleName,
		errMsg:   msg,
	}
}

func New() *RuleMap {
	instance := RuleMap{
		_map: make(map[Name]CheckFunc),
	}
	return &instance
}

func (m *RuleMap) Register(name Name, _func CheckFunc) {
	m._map[name] = _func
}

func (m *RuleMap) Get(name Name) CheckFunc {
	if _func, ok := m._map[name]; ok {
		return _func
	}
	return nil
}

func (m *RuleMap) GetAll() map[Name]CheckFunc {
	return m._map
}

func (m *RuleMap) Remove(name Name) {
	delete(m._map, name)
}

func Register(name Name, _func CheckFunc) {
	GlobalRules.Register(name, func(data *value.Data, fieldName Field, argStr string) (res *CheckError) {
		res = _func(data, fieldName, argStr)
		if res != nil && res.RuleName == "" {
			res.RuleName = name
		}
		return
	})
}

func Get(name Name) CheckFunc {
	return GlobalRules.Get(name)
}

func GetAll() map[Name]CheckFunc {
	return GlobalRules.GetAll()
}

func Remove(name Name) {
	GlobalRules.Remove(name)
}
