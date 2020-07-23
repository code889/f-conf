## 配置
#### 1. 介绍
    鉴于现有的配置使用不熟练
    重写一套用户友好的配置文件
    配置后不常用
    使用友好性牺牲性能
#### 2. 特点

    1. 配置特点
        a.ini
        [a] 对象名
        b=c 对象字段名-对象字段值
        [e]
        f=e
    2. 不可重复
        section、option
    3. 支持特性
        disable
        enable
        依次调用对应函数
        无法识别字段异常
    5. 覆盖
        字段倒序顺序覆盖
    4. 支持空值
        interface{}对应默认值
        
#### 3. 内容忽略
    空行
    注释
    配置行不许有注释
    
#### 4. 用户接口
    结构体
    字段对应section对应结构体变量
    处理结构体变量名和配置section映射关系
    导致配置修改内容变化
    
#### 5. python语言
    python直接构建为字典类型
    不需要映射
    
    