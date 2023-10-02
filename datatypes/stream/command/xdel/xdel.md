# xdel

## syntax

```shell
XDEL key id [id ...]
```

根据ID批量删除记录。
**Time complexity**, O(1)。删除具体的记录并且返回成功删除的个数。如果给出的ID中有不存在的ID，则返回的删除个数可能会小于给出的id个数。





## Understanding the low level details of entries deletion

redis流以内存高效的方式存储着：几十个流记录宏节点被线性打包起来，通过基数树进行索引。当删除一条记录时，该记录并没有从流中删除，而时被标记为已经删除。