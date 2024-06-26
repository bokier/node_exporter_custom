package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

/*
	参考 参考 参考
	这里是参考文件, 正常修改这个文件就行。
	注意：init()需要打开，我这里已经关闭
*/

type loadAvgCollector struct {
	metrics []TypedDesc
}

//func init() {
//	registerCollector("loadavg", NewLoadAvgCollector)
//}

func NewLoadAvgCollector() (Collector, error) {
	return &loadAvgCollector{
		metrics: []TypedDesc{
			{desc: prometheus.NewDesc("node_load1", "1m load avg", []string{"instance", "nodeName"}, nil), valueType: prometheus.GaugeValue},
			{desc: prometheus.NewDesc("node_load3", "3m load avg", []string{"instance", "nodeName"}, nil), valueType: prometheus.GaugeValue},
			{desc: prometheus.NewDesc("node_load5", "5m load avg", []string{"instance", "nodeName"}, nil), valueType: prometheus.GaugeValue},
		},
	}, nil
}

func (l *loadAvgCollector) Update(ch chan<- prometheus.Metric) error {
	loads, err := l.getLoad()
	if err != nil {
		fmt.Println("[error] not find cpu load: ", err)
	}
	for i, load := range loads {
		ch <- l.metrics[i].mustNewConstMetric(load, "instance is i", "nodeName is h")
	}
	return nil
}

func (l *loadAvgCollector) getLoad() ([]float64, error) {
	// 这里是固定数据，可以自定义处理逻辑
	return []float64{1.1, 1.3, 1.5}, nil
}
