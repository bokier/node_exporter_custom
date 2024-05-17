#!/bin/bash

# 指标名称
METRIC="disk"
# 指标lables,可写多个(最小一个)
LABELS="nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy"
# fqName真实展示的名称
QNAME="disk_free"
AFQNAME="disk_free_a"
BFQNAME="disk_free_b"
# help帮助简介
HELP="disk free MB"

AVALUE_TYPE=0
BVALUE_TYPE=0
## valueType 三个分类
# UntypedValue：未分类的值，表示一个未分类的浮点数值。    [1]
# GaugeValue：表示一个浮点数值，通常用于表示可增减的值，如服务器的当前连接数。   [2]
# CounterValue：表示一个浮点数值，通常用于表示只增不减的计数器，如请求数量。 [3]

# 指标
AKEY=/dev/sda
BKEY=/dev/sdb

AVALUE=23941.2
BVALUE=1034.5

# ①数据指标Name是什么 @②数据的labels @③HELP帮助信息 @④Name=FQname @⑤唯一的FQName(FQNAME + 空格 + XKEY) @⑥Metrics指标
# 非常重要的有 ② ⑤ ⑥
# disk@nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy,disk_free=/dev/sda@disk_free@disk_free /dev/sda@23941.1
# disk@nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy,disk_free=/dev/sdb@disk_free@disk_free /dev/sdb@1033.4

# disk@nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy,disk_free=/dev/sda@disk_free@/dev/sda disk free MB@disk_free /dev/sda@23941.1
echo "${AVALUE_TYPE}@${LABELS},${QNAME}=${AKEY}@${AFQNAME}@${AKEY} ${HELP}@${QNAME} ${AKEY}@${AVALUE}"
echo "${BVALUE_TYPE}@${LABELS},${QNAME}=${BKEY}@${BFQNAME}@${BKEY} ${HELP}@${QNAME} ${BKEY}@${BVALUE}"
echo ""
