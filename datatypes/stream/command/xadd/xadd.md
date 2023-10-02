# XADD

## Syntax

```shell
# command
XADD 
# key
key 
# 添加该选项 key为空时，不会创建key-stream对 第一次插入时不能使用该参数
[NOMKSTREAM] 
#
[<MAXLEN | MINID> [= | ~] threshold [LIMIT count]] 
# 流记录ID 
<* | id> 
# 数据
field value [field value ...]
```

将stream节点添加到key上。

如果key不存在，key也会被创建。如果不想auto creation，可以选择`NOMOSTREAM`option。

一个stream节点由多个`field-value`对组成，按照用户给出的顺序进行存储。读取时的顺序与输入的顺序完全相同。

`XADD`是唯一可以向流中添加数据的命令。

### Specifying a Stream ID as an argument 指定流ID作为参数

一个流记录ID定义了流内一个给定的记录。

如果给定的ID参数为*，XADD命令会自动生成一个唯一ID。


### Capped Streams 限制流

`XADD`集成了与`XTRIM`相同的语义。XADD可以增加新的记录并且维持流的大小，从而有效地为流设置一个任意阈值。



### Return

**批量字符串返回**

这个命令返回增加的记录的ID。如果*作为ID参数传递，则该ID是自动生成的，否则该命令仅返回用户在插入期间指定的相同ID。然后跟着对应ID的多条记录。

当key不存在并且使用了`NOMKSTREAM`选项，返回nil。


### go-redis source

```go
type XAddArgs struct {
    // stream ID
	Stream     string
    // true 为添加该选项 false为不添加该选项
	NoMkStream bool
	MaxLen     int64 // MAXLEN N
	MinID      string
	// Approx causes MaxLen and MinID to use "~" matcher (instead of "=").
	Approx bool
	Limit  int64
    // 记录ID
	ID     string
    // 记录的值
	Values interface{}
}

func (c cmdable) XAdd(ctx context.Context, a *XAddArgs) *StringCmd {
	// 构建命令长度
	args := make([]interface{}, 0, 11)
	// 添加命令、流ID
	args = append(args, "xadd", a.Stream)
	// true 为真增加该参数 
	if a.NoMkStream {
		args = append(args, "nomkstream")
	}
	// maxlen、minid 二选一 Approx默认= 为true时
	switch {
	case a.MaxLen > 0:
		if a.Approx {
			args = append(args, "maxlen", "~", a.MaxLen)
		} else {
			args = append(args, "maxlen", a.MaxLen)
		}
	case a.MinID != "":
		if a.Approx {
			args = append(args, "minid", "~", a.MinID)
		} else {
			args = append(args, "minid", a.MinID)
		}
	}
	if a.Limit > 0 {
		args = append(args, "limit", a.Limit)
	}
	if a.ID != "" {
		args = append(args, a.ID)
	} else {
		args = append(args, "*")
	}
	args = appendArg(args, a.Values)

	cmd := NewStringCmd(ctx, args...)
	_ = c(ctx, cmd)
	return cmd
}
```