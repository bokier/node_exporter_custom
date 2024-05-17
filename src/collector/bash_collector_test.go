package collector

import (
	nes "node_exporter_custom/src/nes"
	"reflect"
	"testing"
)

func TestMetricsLabelsHandle(t *testing.T) {
	strList := []string{"nodeName=192.168.0.100", "serverName=diskMetrics", "from=zbwyy", "disk_free=/dev/sda"}

	k, v := nes.PublicModule().MetricsLabelsHandle(strList)
	key := []string{"nodeName", "serverName", "from", "disk_free"}
	value := []string{"192.168.0.100", "diskMetrics", "zbwyy", "/dev/sda"}

	if !reflect.DeepEqual(k, key) { // 如果不匹配的错误返回
		t.Errorf("程序得到的key结果：%s, 我们想要的key结果: %s\n", key, k)
	}

	if !reflect.DeepEqual(v, value) { // 如果不匹配的错误返回
		t.Errorf("程序得到的value结果：%s, 我们想要的value结果: %s\n", value, v)
	}
}
