# xtrim

## syntax

```shell
# command
XTRIM key
# 裁剪策略
<MAXLEN | MINID>
# 精确、模糊
[= | ~]
# 阈值 长度或者记录ID
threshold
# 限制删除的大小
[LIMIT count]
```

**Time complexity O(N)**，N是被丢弃的记录。
`XTRIM`裁剪流通过丢弃旧的记录，优先丢弃小的ID。

`XTRIM`拥有两种策略。

+ **MAXLEN**，只要流的长度一超过`MAXLEN`(threshold是一个给出的正整数)，就丢弃旧的记录。
+ **MINID**，threshold在这种策略中是一个stream ID，执行`TRIM`函数时，丢弃比threshold小的记录。

**精确裁剪**

`XTRIM`默认使用`=`号，it means exact trimming。

+ 裁剪后的流的长度严格小于阈值。
+ 裁剪后的流中的记录的ID严格小于阈值。

**近似精确裁剪**

`XTRIM`使用`~`号时，表示近似精确裁剪，是裁剪操作更高效。精确裁剪会丢掉性能。 

+ `MAXLEN`和`threshold`之间的`~`意味着用户请求修剪函数，流中的长度至少是阈值，可能会稍微大于阈值。当可以获得性能提升时，redis将提前停止修剪。(例如，无法完全删除一个宏节点，宏节点往往存储多个记录)。这会使修剪功能的效率大幅度提高，但是修剪后会比阈值多几十个额外的记录。
+ 使用`~`时控制命令工作量的另一种方法是`LIMIT`子句。当使用它时，它指定将被驱逐的最大条目数。当未指定`LIMIT`和`count`时，将隐式地使用宏节点中条目数的100倍作为计数。将计数值0指定为count会完全禁用限制机制。

## Return

**返回整数**

返回成功删除的记录的个数。