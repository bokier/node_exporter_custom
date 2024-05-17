package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	nes "node_exporter_custom/src/nes"
	"os/exec"
)

/*
	这里是处理自定义bash脚本的collector处理方式
*/

type BashCollector struct {
	metrics []TypedDesc
}

type BashResultData struct {
	Name        string   // 在这里唯一用来代替FqName
	ValueType   []string // 数据类型 prometheus.valueType
	Help        string   // 帮助信息
	LabelsKey   []string // labels中的key值
	LabelsValue []string // labels中的value值
	Metrics     float64  // 指标的实际数值
}

type BashResult map[string]*BashResultData

// BashData save all metrics in collect
var BashData BashResult

func InitBashCollector() {

	for _, c := range nes.Conf.Metrics {

		cmd := exec.Command("sh", "-c", c) // 命令接收
		r, err := cmd.Output()
		if err != nil {
			fmt.Println("命令行出错: ", err)
			return
		}
		// 需要判断结果的行数，并排除空行
		lineList := nes.PublicModule().YamlFileProcessNull(r)
		metricsDataGeneration(lineList)

	}
	// Bash Collector register
	registerCollector("bash", NewBashCollector)
}

// metricsDataGeneration 将数据结构化成 type BashResultData{}
func metricsDataGeneration(s []string) {

	b := make(map[string]*BashResultData)

	for i := 0; i < len(s); i++ {
		resSplit := nes.PublicModule().StringToSplit(s[i], "@")
		labelsList := nes.PublicModule().StringToSplit(resSplit[1], ",")

		k, v := nes.PublicModule().MetricsLabelsHandle(labelsList)

		//将bash返回的结果进行结构化
		b[resSplit[4]] = &BashResultData{
			Name: resSplit[2],
			//FqName:      resSplit[4],
			Help:        resSplit[3],
			Metrics:     nes.PublicModule().StrToFloat64(resSplit[5]),
			LabelsKey:   k,
			LabelsValue: v,
		}
	}
	BashData = b
}

func NewBashCollector() (Collector, error) {

	var bashData []TypedDesc // 手动创建类似BashCollector结构
	var BashDataKey []string

	for k, _ := range BashData {
		BashDataKey = append(BashDataKey, k)
	}

	for i := 0; i < len(BashData); i++ {
		bashData = append(bashData, TypedDesc{
			desc: prometheus.NewDesc(
				BashData[BashDataKey[i]].Name,
				BashData[BashDataKey[i]].Help,
				BashData[BashDataKey[i]].LabelsKey,
				nil,
			),
			valueType: prometheus.GaugeValue,
		})
	}
	return &BashCollector{
		metrics: bashData,
	}, nil
}

func (b *BashCollector) Update(ch chan<- prometheus.Metric) error {
	var BashDataKey []string
	for k, _ := range BashData {
		BashDataKey = append(BashDataKey, k)
	}
	for i := 0; i < len(BashData); i++ {
		ch <- b.metrics[i].mustNewConstMetric(
			BashData[BashDataKey[i]].Metrics, BashData[BashDataKey[i]].LabelsValue...,
		)
	}
	return nil
}
