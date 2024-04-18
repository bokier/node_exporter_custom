package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"node_exporter_custom/src/bash"
	"os/exec"
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
		fmt.Println(string(r))
	}
}

func (b *BashCollector) Update(ch chan<- prometheus.Metric) error {
	return nil
}
