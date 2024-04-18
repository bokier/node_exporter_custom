package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type memAvgCollector struct {
	metrics []TypedDesc
}

func init() {
	registerCollector("memLoad", NewMemAvgCollector)
}

func NewMemAvgCollector() (Collector, error) {
	return &memAvgCollector{
		metrics: []TypedDesc{
			{desc: prometheus.NewDesc("node_memLoad1", "1m memLoad avg", []string{"instance", "nodeName"}, nil), valueType: prometheus.GaugeValue},
			{desc: prometheus.NewDesc("node_memLoad3", "3m memLoad avg", []string{"instance", "nodeName"}, nil), valueType: prometheus.GaugeValue},
			{desc: prometheus.NewDesc("node_memLoad5", "5m memLoad avg", []string{"instance", "nodeName"}, nil), valueType: prometheus.GaugeValue},
		},
	}, nil
}

func (m *memAvgCollector) Update(ch chan<- prometheus.Metric) error {
	loads, err := m.getLoad()
	if err != nil {
		fmt.Println("[error] not find mem load: ", err)
	}
	for i, load := range loads {
		ch <- m.metrics[i].mustNewConstMetric(load, "instance is i", "nodeName is h")
	}
	return nil
}

func (m *memAvgCollector) getLoad() ([]float64, error) {
	return []float64{2.1, 2.3, 2.5}, nil
}
