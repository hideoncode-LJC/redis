# underlying data structure

> 记录 underlying dataStructure 的结构和知识

**Redis为什么那么快**

1. *Redis* 是内存数据库，不需要进行磁盘IO，查找效率快。
2. *Redis* 使用的数据结构能使我们对数据进行CURD时，*Redis*能更高效的实现。

Redis 数据结构并不是指 String（字符串）对象、List（列表）对象、Hash（哈希）对象、Set（集合）对象和 Zset（有序集合）对象，因为这些是 Redis 键值对中值的数据类型，也就是数据的保存形式，这些对象的底层实现的方式就用到了数据结构。

![Alt text](image.png)

**键值对数据库是怎么实现的？**



## SDS

## list

### listnode design

```go
package list

type any int

type ListNode struct {
	// 前置节点
	pre *ListNode
	// 后驱节点
	next *ListNode
	// 指向任何类型的指针
	value *any
}

type List struct {
	// head
	head *ListNode
	// tail
	tail *ListNode
	// dup 节点值复制
	dup func(*ListNode, *ListNode)
	// free 释放节点
	free func(*ListNode)
	// match
	match func(*ListNode, *ListNode) int
	// length 64bit整数
	len uint64
}

```

### Advantage 
+ 获取前驱节点、后置节点、头节点、尾节点、链表的长度 时间复杂度都为O(1)
+ 这两个指针都可以指向 NULL，所以链表是无环链表？？
+ listNode 链表节使用 void* 指针保存节点值，并且可以通过 list 结构的 dup、free、match 函数指针为节点设置该节点类型特定的函数，因此链表节点可以保存各种不同类型的值；?????



### disadvantage

+ 链表结构每个节点之间内存都是不连续的，无法更好的利用缓存。

+ 存储密度低

> 利用缓存最好的结构是数组，数组每个元素之间地址是连续的。

### 使用场景

当value为基本数据类型List时，如果数据量多，则使用List存储。当数据量较少时，使用`压缩列表`存储。

压缩列表的优势是节省内存空间，内存紧凑。
但是压缩列表存在性能问题。

所以 Redis 在 3.2 版本设计了新的数据结构 quicklist，并将 List 对象的底层数据结构改由 quicklist 实现。

然后在 Redis 5.0 设计了新的数据结构 listpack，沿用了压缩列表紧凑型的内存布局，最终在最新的 Redis 版本，将 Hash 对象和 Zset 对象的底层数据结构实现之一的压缩列表，替换成由 listpack 实现。

## ziplist

### why need ziplist Advantage

1. 链表结构每个节点之间内存不连续，无法利用CPU cache。而ziplist被设计成内存紧凑型结构，可能高效率的利用CPU cache。
2. ziplist对数据进行`编码`，节省内存开销。

### Disadvantage

1. 不能保存过多的元素，过多的元素会使查找效率降低。
2. 新增或者修改某个元素时，会重新分配内存，甚至会引发连锁更新的问题。 连锁更新？？？？
3. 因此，在旧版Redis中，List、Hash、Zset 基本类型存储的数据不多时底层数据结构使用ziplist


### zip list design

