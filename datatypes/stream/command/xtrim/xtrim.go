package xtrim

import (
	"context"
	"redis/datatypes/connect"
)

var rdb = connect.GetConn(6)

func XTRIM() {
	cmd := rdb.XTrimMaxLen(context.Background(), "xadd:test1", 50)
	connect.Printlogger("cmd = ", cmd)
}