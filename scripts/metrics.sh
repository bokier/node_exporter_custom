#!/bin/bash

# 指标名称
METRIC="disk"
# 指标lables,可写多个(最小一个)
LABELS="nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy"
# fqName真实展示的名称
FQNAME="disk_free"
# help帮助简介
HELP="disk free MB"


# 指标
AKEY=/dev/sda
BKEY=/dev/sdb

AVALUE=23941.1
BVALUE=1033.4

# ①数据指标Name是什么 @②数据的labels @③HELP帮助信息 @④Name=FQname @⑤唯一的FQName(FQNAME + 空格 + XKEY) @⑥Metrics指标
# 非常重要的有 ② ⑤ ⑥
# disk@nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy,disk_free=/dev/sda@disk_free@disk_free /dev/sda@23941.1
# disk@nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy,disk_free=/dev/sdb@disk_free@disk_free /dev/sdb@1033.4

# disk@nodeName=192.168.0.100,serverName=diskMetrics,from=zbwyy,disk_free=/dev/sda@disk_free@/dev/sda disk free MB@disk_free /dev/sda@23941.1
echo "${METRIC}@${LABELS},${FQNAME}=${AKEY}@${FQNAME}@${AKEY} ${HELP}@${FQNAME} ${AKEY}@${AVALUE}"
echo "${METRIC}@${LABELS},${FQNAME}=${BKEY}@${FQNAME}@${BKEY} ${HELP}@${FQNAME} ${BKEY}@${BVALUE}"
echo ""
