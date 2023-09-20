package hyperloglog

import (
	connect "go-redis/contect"
	"testing"

	
)

func TestHelpHyperloglog(t *testing.T) {

	rdb := connect.GetRedisCLient()

	HelpHyperloglog(rdb)
}