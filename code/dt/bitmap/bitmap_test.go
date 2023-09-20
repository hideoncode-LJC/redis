package bitmap

import (
	connect "go-redis/contect"
	"testing"
	"fmt"
	command "go-redis/command"
)

func TestSetBit(t *testing.T) {

	rdb := connect.GetRedisCLient()
	SetBit(rdb)

	fmt.Println("查询数据")

	for i := 1; i < 13; i++ {
		command.Get(rdb, "zhangsan-" + fmt.Sprint(i))
	}
}