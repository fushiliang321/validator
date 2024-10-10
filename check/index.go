package check

import (
	"github.com/fushiliang321/validator/rule"
	"github.com/fushiliang321/validator/value"
	"strings"
)

func Execute(data map[rule.Field]any, rules map[rule.Field]rule.RuleStr, allCheck bool) (res []*rule.CheckError) {
	var (
		sep                                                                int
		_field, _fieldStrItem, _ruleStr, _ruleName, _ruleArgs, ruleStrItem string
		fun                                                                rule.CheckFunc
		err                                                                *rule.CheckError
		ruleStrs, _fieldStrs                                               []string
		newData                                                            = value.Transition(data)
	)
	for _field, _ruleStr = range rules {
		ruleStrs = strings.Split(_ruleStr, "|")
		_fieldStrs = strings.Split(_field, "|")
		for _, ruleStrItem = range ruleStrs {
			sep = strings.Index(ruleStrItem, ":")
			if sep == -1 {
				_ruleName = ruleStrItem
				_ruleArgs = ""
			} else {
				_ruleName = ruleStrItem[:sep]
				_ruleArgs = ruleStrItem[sep+1:]
			}

			if fun = rule.Get(_ruleName); fun != nil {
				for _, _fieldStrItem = range _fieldStrs {
					err = fun(newData, _fieldStrItem, _ruleArgs)
					if err != nil {
						res = append(res, err)
						if !allCheck {
							return
						}
					}
				}
			}
		}
	}
	return
}
