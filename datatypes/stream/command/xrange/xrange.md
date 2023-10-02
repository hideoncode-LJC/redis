# xrange

## syntax

```shell
# command
XRANGE 
# stream key
key 
# 通过ID确定范围
start end 
# 选择获取几个记录
[COUNT count]
```

该命令返回匹配的流记录根据一个给定的ID范围内。
该范围通过start(minimum)、end(maximum)来具体确定。