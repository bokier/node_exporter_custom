package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"node_exporter_custom/src/bash"
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
	return bash.PublicModule().IntToFloat64(len(lines) - 1)
}

func (u *UserCollector) getMachineIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("Error:", err)
		return "1.1.1.1"
	}

	for _, addr := range addrs {
		// 检查地址类型为IP地址，并且不是回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("服务器IP地址:", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return "2.2.2.2"
}
