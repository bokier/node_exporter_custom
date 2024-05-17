package bash

/*
	这里是用于检测bash脚本中的错误 v0.0.1
	具体规则:
	1. 必须含有六个值，第一个值为prometheus.valueType且最后一个值是metrics并且可以转换成float64
	2. 第三个值唯一
    ### 这里没启用 ###
*/

type RulesBashStruct struct {
}

var RulesBash RulesBashStruct

// IsEmptyRow 判断脚本结果是不是空行
func (r RulesBashStruct) IsEmptyRow() bool {
	return true
}

// IsRowSatisfy 判断其中的行是否满足条件
func (r RulesBashStruct) IsRowSatisfy() bool {
	return true
}

// RowSatisfyNumber 满足的行数量
func (r RulesBashStruct) RowSatisfyNumber() int {
	return 1
}
