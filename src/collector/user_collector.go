package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	nes "node_exporter_custom/src/nes"
	"os/exec"
	"strings"
)

type UserCollector struct {
	metrics TypedDesc
}

func init() {
	registerCollector("user", NewUserCollector)
}

func NewUserCollector() (Collector, error) {
	return &UserCollector{
		metrics: TypedDesc{
			desc:      prometheus.NewDesc("user_login", "machine login user num", []string{"instance"}, nil),
			valueType: prometheus.GaugeValue,
		},
	}, nil
}

func (u *UserCollector) Update(ch chan<- prometheus.Metric) error {
	user := u.getLoginUser()
	ip := u.getMachineIP()

	ch <- u.metrics.mustNewConstMetric(user, ip)

	return nil
}

func (u *UserCollector) getLoginUser() float64 {
	cmd := exec.Command("who")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error running who command:", err)
	}

	// 将输出按换行符分割成行
	lines := strings.Split(string(output), "\n")
	return nes.PublicModule().IntToFloat64(len(lines) - 1)
}

func (u *UserCollector) getMachineIP() string {
	return nes.Conf.NodeIp
}
