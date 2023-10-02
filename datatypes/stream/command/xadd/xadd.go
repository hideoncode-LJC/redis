package xadd

import (
	"context"
	"fmt"
	"redis/datatypes/connect"
	"sync"

	"github.com/redis/go-redis/v9"
)

var rdb = connect.GetConn(6)

var wg sync.WaitGroup

func addData(n int) {

	values := make([]string, 0, 200)

	for i := 1 ; i <= n ; i++ {
		values = append(values, "feild" + fmt.Sprint(i))
		values = append(values, "value" + fmt.Sprint(i))
	}

	connect.Printlogger("values = ", len(values))

	cmd := rdb.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "xadd:test1",
		// NoMkStream: true,
		MaxLen: 98,
		//MinID: "",
		Approx: true,
		ID: "*",
		Values: values,
	})

	connect.Printlogger("cmd = ", cmd)

	wg.Done()
}

func XADD() {

	wg.Add(100)
	
	for i := 1 ; i <= 100 ; i++ {
		go addData(i)
	}
	
	wg.Wait()
}