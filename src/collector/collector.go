package collector

/*
	collector 主要文件
*/
import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

var (
	factories              = make(map[string]func() (Collector, error))
	initiatedCollectorsMtx = sync.Mutex{}
	initiatedCollectors    = make(map[string]Collector)
	collectorState         = make(map[string]bool)
)

const namespace = "node_custom"

func registerCollector(collector string, factory func() (Collector, error)) {
	collectorState[collector] = true
	factories[collector] = factory
}

type NodeCollector struct {
	Collectors map[string]Collector
}

func NewNodeCollector(filters ...string) (*NodeCollector, error) {
	f := make(map[string]bool)
	for _, filter := range filters {
		enabled, exist := collectorState[filter]
		if !exist {
			return nil, fmt.Errorf("missing collector: %s", filter)
		}
		if !enabled {
			return nil, fmt.Errorf("disabled collector: %s", filter)
		}
		f[filter] = true
	}

	collectors := make(map[string]Collector)
	initiatedCollectorsMtx.Lock()
	defer initiatedCollectorsMtx.Unlock()

	//collectorState["loadavg"] = true
	// 注意这个地方，如果设置了没有的名称，会有错误
	collectorState["user"] = true
	collectorState["bash"] = true

	for key, enabled := range collectorState {
		if !enabled || (len(f) > 0 && !f[key]) {
			continue
		}
		if collector, ok := initiatedCollectors[key]; ok {
			collectors[key] = collector
		} else {
			collector, err := factories[key]()
			if err != nil {
				return nil, err
			}
			collectors[key] = collector
			initiatedCollectors[key] = collector
		}
	}
	return &NodeCollector{
		Collectors: collectors,
	}, nil
}

func (n NodeCollector) Describe(ch chan<- *prometheus.Desc) {}

func (n NodeCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(n.Collectors))
	for name, c := range n.Collectors {
		go func(name string, c Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
}

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

func execute(name string, c Collector, ch chan<- prometheus.Metric) {
	err := c.Update(ch)
	if err != nil {
		fmt.Println("[error] find a error in update ", name)
	}
}
