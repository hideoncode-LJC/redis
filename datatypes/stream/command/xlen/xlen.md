# xlen

## syntax

```shell
XLEN key
```

返回流中记录的个数。
如果给出的key不存在也返回0，就像流中不存在记录一样。
所以当XLEN返回0时，无法确定该流是空的还是key不存在。
所以，判断一个流是否存在，需要使用type、exists判断。
当流中所有元素被删除时，流也不会自动被删除，可能会存在消费组和该流联系起来。

## Return

**返回整数**

返回记录的个数。

## go-redis source

```go
func (c cmdable) XLen(ctx context.Context, stream string) *IntCmd {
	cmd := NewIntCmd(ctx, "xlen", stream)
	_ = c(ctx, cmd)
	return cmd
}

type IntCmd struct {
	baseCmd
	val int64
}
```

