package zset

import (
	"redis/datatypes/connect"
	"testing"
)



func TestDiff(t *testing.T) {
	// DIFFPRE()


	// DIFFEXEC()
}

func TestUnion(t *testing.T) {
	// rdb := connect.GetConn(13)
	// UNIONPRE(rdb)
	// UNION(rdb)
}


func TestMPOPPRE(t *testing.T) {
	MPOPPRE(connect.GetConn(10))
}

func TestZCOUNTPRE(t *testing.T) {
	ZCOUNTPRE(connect.GetConn(5))
}

func TestZRANKPRE(t *testing.T) {
	ZRANKPRE(connect.GetConn(4))
}


