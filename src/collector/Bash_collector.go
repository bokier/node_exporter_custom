package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"node_exporter_custom/src/bash"
	"os/exec"
	"strings"
)

/*
	这里是处理自定义bash脚本的collector处理方式
*/

type BashCollector struct {
	metrics []TypedDesc
}

func InitBashCollector() {
	fmt.Println("[metrics length]: ", len(bash.Conf.Metrics))
	for k, c := range bash.Conf.Metrics {
		fmt.Println("[info] colletor ", k)

		cmd := exec.Command("sh", "-c", c) // 命令接收
		r, err := cmd.Output()
		if err != nil {
			fmt.Println("命令行出错: ", err)
		}
		// 需要判断结果的行数，并排除空行
		lineList := yamlFileProcessNull(r)

		//
		for i := 0; i < len(lineList); i++ {
			res := lineList[i]
			fmt.Printf("[数据]: %v \n", res)
			//resSplit := stringSplit(res, "@")
		}
	}
}

func (b *BashCollector) Update(ch chan<- prometheus.Metric) error {
	return nil
}

// stringSplit string 按照@分割
func stringSplit(str string, s string) []string {
	return strings.Split(str, s)
}

func yamlFileProcessNull(r []byte) []string {
	var lineList []string
	lineRes := stringSplit(string(r), "\n")

	for _, v := range lineRes {
		if len(v) == 0 {
			continue
		}
		lineList = append(lineList, v)
	}
	return lineList
}

// goPrecessNum 启动多少个线程来处理
func goPrecessNum() int {
	return 3
}
