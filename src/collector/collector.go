package collector

/*
	collector 主要文件
*/
import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "node_custom"

type NodeCollector struct{}

func (n NodeCollector) Describe(ch chan<- *prometheus.Desc) {}

func (n NodeCollector) Collect(ch chan<- prometheus.Metric) {}

type Collector interface {
	Update(ch chan<- prometheus.Metric) error
}

type TypedDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

func (t *TypedDesc) mustNewConstMetric(value float64, labels ...string) prometheus.Metric {
	return prometheus.MustNewConstMetric(t.desc, t.valueType, value, labels...)
}
