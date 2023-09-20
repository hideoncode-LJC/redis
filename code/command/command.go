package command

import (
	"context"
	"fmt"
	"log"
	"github.com/redis/go-redis/v9"
)

func Keys(rdb *redis.Client, arg string) {


	data, err := rdb.Do(context.Background(), "KEYS", arg).Result()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("the data of ", arg + " = ", data)
}

func Get(rdb *redis.Client, arg string) {
	data, err := rdb.Get(context.Background(), arg).Result()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("key = ", arg, "data = ", data)
}