package validator

import (
	"github.com/fushiliang321/validator/check"
	"github.com/fushiliang321/validator/rule"
)

func Check(data map[rule.Field]any, rules map[rule.Field]rule.RuleStr) []*rule.CheckError {
	return check.Execute(data, rules, true)
}

func CheckOne(data map[rule.Field]any, rules map[rule.Field]rule.RuleStr) *rule.CheckError {
	res := check.Execute(data, rules, false)
	if len(res) == 0 {
		return nil
	}
	return res[0]
}
