# SDS

## Why need SDS

**为什么Redis不使用c中的字符串**

1. c语言中的字符串获取字符串的长度需要遍历整个字符串，效率太低
2. c语言中的字符串修改麻烦
3. 非二进制安全

因此，Redis设计了一种新的字符串 - Simple Dynamic String，简称SDS。

在Redis中，string类型的数据都会通过SDS进行存储。

## SDS Struct

```go
package main

// 为了节省内存 SDS 声明了五种结构体

const (
    len5 = iota
    len8
    len16
    len32
    len64
)

type sds8 struct {
    // 字符串的长度 不包含 '\0'
    len uint8
    // 申请的内存长度 不包含 '\0'
    alloc uint8
    // 当前结构体属于哪种
    flag uint8
    // 字符串存储区
    buf []byte
}

type sds5  struct {}
type sds16 struct {}
type sds32 struct {}
type sds64 struct {}

// 动态扩容字符串、内存预申请   
func (s *sds8) dynamicExpand(str string) {
    if str.size + sds8.len > sds8.alloc {
        if str.size < 1024 * 1024 * 8 {
                
        } else {

        }
    }
} 
```

## SDS Feature
1. 获取字符的长度时间复杂度为O(1)
2. 支持动态扩容
3. 减少内存分配次数
4. 二进制安全