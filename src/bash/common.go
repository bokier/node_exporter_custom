package bash

import (
	"strconv"
	"strings"
)

/*
	这里是公共的转换/处理通用数据的模块
	(string/[]string切割，格式转换等)
*/

type PublicProcessModule struct{}

func PublicModule() *PublicProcessModule {
	return &PublicProcessModule{}
}

// YamlFileProcessNull 去除文件中的空行， 并将文件内容转成 []string列表的形式
func (p PublicProcessModule) YamlFileProcessNull(r []byte) []string {
	var lineList []string
	lineRes := p.StringToSplit(string(r), "\n")

	for _, v := range lineRes {
		if len(v) == 0 {
			continue
		}
		lineList = append(lineList, v)
	}
	return lineList
}

// StrToFloat64 将string类型转换成float64
func (p PublicProcessModule) StrToFloat64(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)
	return res
}

// IntToFloat64 将int类型转换成float64
func (p PublicProcessModule) IntToFloat64(i int) float64 {
	return float64(i)
}

// StringToSplit string 切割，传入string和分割符号
func (p PublicProcessModule) StringToSplit(str string, s string) []string {
	return strings.Split(str, s)
}

// StringListToStringByIndex 将
func (p PublicProcessModule) StringListToStringByIndex(strList []string, i int) string {
	return ""
}

// MetricsLabelsHandle metrics指标中labels中key value处理方式
func (p PublicProcessModule) MetricsLabelsHandle(str []string) ([]string, []string) {
	var key []string
	var value []string
	for _, v := range str {
		n, m := p.stringLabelsKv(v)
		key = append(key, n)
		value = append(value, m)
	}
	return key, value
}

// stringLabelsKv 将收集上来的数据进行格式化处理，按照"="分割，返回key和value
func (p PublicProcessModule) stringLabelsKv(str string) (k string, v string) {
	splitStr := "="
	res := p.StringToSplit(str, splitStr)
	return res[0], res[1]
}
