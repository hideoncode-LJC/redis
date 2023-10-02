# Zset

# Zset简介

# ZSET常用命令

## 查询类命令

### ZCARD

```shell
ZCARD key
```
根据`key`返回score-member对的个数。

返回`整数类型`，如果key不存在或者空集合，返回0。

其他情况返回error。

****

### ZCOUNT

```shell
ZCOUNT key min max
```

根据key和min、max，得到位于min、max之间元素(score-member对)的个数。

返回`整数`类型。



### ZSCORE

>Sorted Set Score

```
ZSCORE key member
```

根据`key`和`member`获取分数。

返回字符串类型，存储着浮点数类型。

如果member or key 不存在，返回nil。

### ZMSCORE

```shell
ZMSCORE key member [member ...]
```

返回`字符串列表类型`

+ 至少需要输入一个member，不输入返回error。
+ key不存在，输入的member都会返回nil值。
+ key存在，member不存在对应的member返回nil值。
+ key对应的value类型不是zset，会返回error。



### ZRANK

```shell
ZRANK key member [WITHSCORES]
```

函数功能

+ 输入key、member，得到member的rank值。
+ rank值根据member的分数计算，分数最低的member的rank为0。

函数返回
+ 




## RANGE类命令

### Index

+ ZRANGE
+ ZRANGESTORE

### ZRANGE

```shell
ZRANGE 
key start stop 
[BYSCORE | BYLEX] 
[REV] 
[LIMIT offset count]
[WITHSCORES]
```

## REM类命令

### INDEX

+ REM
+ REMRANGEBYLEX
+ REMRANGEBYRANK
+ REMRANGEBYSCORE

### REM



### 交集、并集、差集

**ZUNION**

>ZSET UNION

```shell
# command name
ZUNION 
# 输入的key的数量 
numkeys 
# 多个key
key [key ...]
# 默认WIEGHTS 1
[WEIGHTS weight [weight ...]]
# 默认为sum
[AGGREGATE <SUM | MIN | MAX>]
```

计算规则：计算多个key的并集。

WEIGHTS：weight的数量要和key的数量匹配。

AGGREGATE：是对结果集合中每个key的score求和、最大值、最小值。

没有`WITHSCORES`参数，返回member list，反之返回member-score list。

**ZUNIONSTORE**

```shell
# command name
ZUNION 
# 输入的key的数量 

# 目的key
dest

numkeys 
# 多个key
key [key ...]
# 默认WIEGHTS 1
[WEIGHTS weight [weight ...]]
# 默认为sum
[AGGREGATE <SUM | MIN | MAX>]
```

如果destkey已经存在，则会覆盖。

返回目的key中score-member对的个数。




**ZINTER**

```shell
# Command name
ZINTER
# key的数量
numkeys key [key ...]
# weight 
[WEIGHTS weight [weight ...]]
# sum、min、max
[AGGREGATE <SUM | MIN | MAX>]
# 返回分数 
# 不加该选项返回 member list
# 加该选项返回 score-member list
[WITHSCORES]
```




**ZINTERCARD**

```shell
ZINTERCARD numkeys key [key ...] [LIMIT limit]
```
返回交集中元素的个数，当不提供`LIMIT`参数时，表示无限制。
提供该参数时，当交集中的数量达到限制个数时，算法直接退出，提高了效率。


**ZINTERSTORE**

```shell
ZINTERSTORE 
destination 
numkeys key [key ...] 
[WEIGHTS weight [weight ...]] 
[AGGREGATE <SUM | MIN | MAX>]
```

dest已经存在会被覆盖。

返回score-member list，存入目标key。

**ZDIFF**

```shell
ZDIFF numkeys key [... key] [WITHSCORES]
```

1. 求几个key之间的差集。
2. 返回几个key之间差集的member，应列表形式返回。
3. 如果key不存在，被视为空集合。
4. 其他情况返回错误。

**ZDIFFSTORE**

同时返回score-member对，并存储在目标key中。


### POP 弹出元素

#### MPOP

> multiply POP element



#### BMPOP

>Block Multiply POP

```shell
BZMPOP 
timeout 
numkeys key [key ...] 
<MIN | MAX> 
[COUNT count]
```


### 



## ZSET 底层
