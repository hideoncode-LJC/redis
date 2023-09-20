package main

// 四位无符号整数
type int4 int

type redisObject struct {
	/*
	type 对象类型 无符号4位整数 0 - 15
	#define OBJ_STRING 0
	#define OBJ_LIST 1
	#define OBJ_SET 2
	#define OBJ_ZSET 3
	#define OBJ_HASH 4
	*/
	Type int4
	/*
	编码方式 五种类型的底层编码方式
	目前有11种编码法师
	*/
	encoding int4 

	// 24bit
	lru 
}