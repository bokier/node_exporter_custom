package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type loadAvgCollector struct {
	metrics []TypedDesc
}

func init() {
	registerCollector("loadavg", NewLoadAvgCollector)
}

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
	return []float64{1.1, 1.3, 1.5}, nil
}
