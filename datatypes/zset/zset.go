package zset

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var logger = log.New(os.Stdout, "\033[34m[info ]\033[0m", log.Lshortfile | log.Lmicroseconds | log.Ldate)

// func getConn(db int) *redis.Client{
// 	return redis.NewClient(&redis.Options{
// 		Addr: "127.0.0.1:6379",
// 		Password: "123456",
// 		DB: db,
// 	})
// }

// 普通查询---------------------------------------------------------------------------

/*
ZSCORE
根据key Member 获取分数
*/


/*
ZCOUNT key min max
*/

func ZCOUNTPRE(rdb *redis.Client) {
	for i := 1 ; i <= 10 ; i ++ {
		cmd := rdb.ZAdd(context.Background(), "ZSET:ZCOUNT", redis.Z{
			Score: float64(i),
			Member: "member" + fmt.Sprint(i),
		})

		logger.Println("cmd = ", cmd)
	}
}

/*
ZRANK
*/

func ZRANKPRE(rdb *redis.Client) {
	for i := 1 ; i <= 10 ; i ++ {
		cmd := rdb.ZAdd(context.Background(), "Z:RANK", redis.Z{
			Score: float64(i),
			Member: "mem" + fmt.Sprint(i),
		})

		logger.Println("cmd = ", cmd)
	}

	cmd := rdb.ZAdd(context.Background(), "Z:RANK", redis.Z{
		Score: 10,
		Member: "mem11",
	}, redis.Z{
		Score: 10,
		Member: "mem12",
	},redis.Z{
		Score: 10,
		Member: "mem13",
	})

	logger.Println("cmd = ", cmd)
}

// func DIFF(rdb *redis.Client, keys []string) *redis.StringSliceCmd {
// 	return rdb.ZDiff(context.Background(), keys)
// }

func UNIONPRE(rdb *redis.Client) {
	cmd1 := rdb.ZAdd(context.Background(), "ZSET:ZUNION:1", redis.Z{
		Score: 1,
		Member: "one",
	},redis.Z{
		Score: 2,
		Member: "two",
	},redis.Z{
		Score: 3,
		Member: "three",
	})

	logger.Println("cmd = ", cmd1)

	cmd2 := rdb.ZAdd(context.Background(), "ZSET:ZUNION:2", redis.Z{
		Score: 1,
		Member: "one",
	},redis.Z{
		Score: 2,
		Member: "two",
	})

	logger.Println("cmd = ", cmd2)
}

func UNION(rdb *redis.Client) {

	// zunion 2 ZSET:ZUNION:1 ZSET:ZUNION:2 weights 5 10 aggregate max withscores: [{10 one} {15 three} {20 two}]

	cmd1 := rdb.ZUnionWithScores(context.Background(), redis.ZStore{
		Keys: []string{"ZSET:ZUNION:1", "ZSET:ZUNION:2"},
		Weights: []float64{5, 10},
		Aggregate: "max",
	})

	logger.Println("max = ", cmd1)

	// zunion 2 ZSET:ZUNION:1 ZSET:ZUNION:2 weights 5 10 aggregate min withscores: [{5 one} {10 two} {15 three}]

	cmd2 := rdb.ZUnionWithScores(context.Background(), redis.ZStore{
		Keys: []string{"ZSET:ZUNION:1", "ZSET:ZUNION:2"},
		Weights: []float64{5, 10},
		Aggregate: "min",
	})

	logger.Println("min = ", cmd2)

	// zunion 2 ZSET:ZUNION:1 ZSET:ZUNION:2 weights 5 10 withscores: [{15 one} {15 three} {30 two}]

	cmd3 := rdb.ZUnionWithScores(context.Background(), redis.ZStore{
		Keys: []string{"ZSET:ZUNION:1", "ZSET:ZUNION:2"},
		Weights: []float64{5, 10},
	})
	
	logger.Println("sum = ", cmd3)

	// zunion 2 ZSET:ZUNION:1 ZSET:ZUNION:2 withscores: [{2 one} {3 three} {4 two}]

	cmd4 := rdb.ZUnionWithScores(context.Background(), redis.ZStore{
		Keys: []string{"ZSET:ZUNION:1", "ZSET:ZUNION:2"},
	})

	logger.Println("deafult = ", cmd4)
}


/* ------------------------------------------------------------------
MPOP

---------------------------------------------------------------------*/

func MPOPPRE(rdb *redis.Client) {
	for i := 1 ; i <= 10 ; i++ {
		for j := 1 ; j <= 10 ; j++ {
			cmd := rdb.ZAdd(context.Background(), "ZSET:MPOP:" + fmt.Sprint(i), redis.Z{
				Score: float64(i * j),
				Member: fmt.Sprint(i) + ":member:" + fmt.Sprint(j),
			})
			logger.Println("i = ", i, " j = ", j, " cmd = ", cmd)
		}
	}
}

// func MPOP(rdb *redis.Client) {
// 	rdb.ZMPop(context.Background(), ""
// }f
//
//



