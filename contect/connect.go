package connect

import (
	"fmt"
	config "go-redis/config"

	"github.com/redis/go-redis/v9"
)

func GetRedisCLient() *redis.Client {
	// get redis config struct
	rdc := config.GetRdb()

	// output redis config struct
	fmt.Printf("%v\t%T\n", rdc, rdc)

	// return redis client
	return redis.NewClient(&redis.Options{
		Addr: rdc.IP + ":" + rdc.Port,
		Password: rdc.Password,
		DB: rdc.DB,
		})
}

// func main() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "192.168.81.133:6379",
// 		Password: "123456", // no password set
// 		DB:       0,        // use default DB
// 		PoolSize: 20,
// 	})

// 	val, err := rdb.Get("key").Result()

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	fmt.Println("value = ", val)

// }

// func InsertKeys() {
// 	for i := 1 ; i <= 10 ; i++ {
// 		rdb 
// 	}
// }


