# Set

## IntSet

+ 长度可变
+ 有序

```go
type intset struct {
    // 编码方式 
    // 支持 16位整数、32位整数、64位整数
    encoding uint32
    // 元素个数
    length uint32
    // 整数数组、保存集合数据
    // -128 - 127
    // contents []byte
    contents []int8
}
```

