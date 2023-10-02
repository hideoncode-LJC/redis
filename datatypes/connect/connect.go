package connect

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func GetConn(db int) *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "123456",
		DB: db,
	})
}

func Getlogger() *log.Logger{
	return log.New(os.Stdout, "\033[34m[info]\033[0m", log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

func Printlogger(t ...any) {
	logger := Getlogger()
	logger.Println(t[0], t[1])
}