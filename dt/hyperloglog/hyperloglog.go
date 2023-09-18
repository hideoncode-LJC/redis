package hyperloglog

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// func main() {
// 	// 得到当前redis占用的内存
// 	rdb := getRdb()

// 	fmt.Printf("%v\n", rdb)

// 	data, err := rdb.Do(context.Background(), "INFO", "memory").Result()

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%v\t%T\n", data, data)

// 	// dataJson := json.Unmarshal([]byte(data), v)
// }

// func getRdb() *redis.Client {
// 	rdc := config.GetRdb()

// 	fmt.Println(rdc)

// }

// 获取 hyperloglog 相关操作
func HelpHyperloglog(rdb *redis.Client) {
	data, err := rdb.Do(context.Background(), "KEYS", "*").Result()
	
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}

// func PFADD()

