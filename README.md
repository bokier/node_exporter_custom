<font color="red">当前时间: 2024/04/18 </font>

## 一、说明
    用于自定义的node_exporter监控扩展.
    实现自定义key和value的指标监控. 初步想法是脚本获取值的方式
    <让prometheus脚本化>

### custom.yaml配置文件说明
    apiEnable: "true"  // 从api获取指标是否启用[true/false]
    bashEnable: "true" // 从bash获取指标是否启用[true/false]

## 二、版本说明
### v0.0.1 2024/04/24
    此版本可以测试使用，基本功能已经实现
    后续完善和添加从api取值方式

## 三、启动参数
### 默认启动
    ./node_exporter_custom &>/dev/null &&
    端口 20240
    由于脚本问题导致请求出错的话，程序会退出

### 脚本说明
    v0.0.1版本：
    < sh脚本数据返回格式强一致性 >
    ① 输出string
    ② 共6个参数，以@符号隔开

    具体参数:
    ① 指标名称，可以自定义有含义即可 - <string>
    ② 标签labels，重要 - <[]string>  // name=node01,ip=10.0.0.1
    ③ FQName，重要 - <string>  // 不能含有空格，且在一个实例中唯一
    ④ HELP提示信息，用于描述这个指标是干嘛的 - <string> 
    ⑤ 一个不重要的东西，可以任意写 -<string>
    ⑥ 指标，重要 -<string>
    
    提示：
    可以多个脚本，或者一个脚本中输出多条信息
    但是必须要每个内容都有

## 四、正在开发
    v0.0.2版本 完善关于bash获取方式和规范
    1. 关闭http请求记录日志 ✅
    2. 判断数据指标获取内容有无问题