package bitmap

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/redis/go-redis/v9"
)

func SetBit(rdb *redis.Client) {
	fmt.Println("插入数据")
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 1; i < 13; i++ {
		for j := 1; j < days[i]; j++ {
			flag := rand.Int()
			if flag%2 == 0 {
				data, err := rdb.Do(context.Background(), "SETBIT", "zhangsan-"+fmt.Sprint(i), fmt.Sprint(j), "1").Result()
				if err != nil {
					log.Fatalln(err)
				} else {
					fmt.Println("data = ", data)
				}
			} else {
				data, err := rdb.Do(context.Background(), "SETBIT", "zhangsan-" + fmt.Sprint(i), fmt.Sprint(j), "0").Result()
				if err != nil {
					log.Fatalln(err)
				} else {
					fmt.Println("data = ", data)
				}
			}
		}
	}

	
}

// Returns the bit value at offset in the string value stored at key
// 获取某一个位置的bit值
// GETBIT key offset  offset 表示位置
func GetBit() {
	
}

