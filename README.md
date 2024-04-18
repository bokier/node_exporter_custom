<font color="red">当前时间: 2024/04/18 </font>

## 一、说明
    用于自定义的node_exporter监控扩展.
    实现自定义key和value的指标监控. 初步想法是脚本获取值的方式

### 编译方式


## 二、启动参数
### 默认启动
    ./node_exporter_custom &>/dev/null &&

### 环境变量
    # 这里是正常启动方式
    TZ=Asia/Shanghai  
    FILE=/etc/node_exporter/custom.yml  
    PORT=12024
