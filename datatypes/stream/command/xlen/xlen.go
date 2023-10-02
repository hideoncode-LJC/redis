package xlen

import (
	"context"
	"redis/datatypes/connect"
)

var rdb = connect.GetConn(6)

func XLEN() {
	cmd := rdb.XLen(context.Background(), "xadd:test1")

	connect.Printlogger("cmd = ", cmd)
}