package main

import (
	"flag"

	"github.com/erbaner/be-core/pkg/log"
	"github.com/erbaner/be-core/test"
)

func main() {
	var onlineNum *int //Number of online users
	onlineNum = flag.Int("on", 10, "online num")
	flag.Parse()
	log.Warn("", "online test start, online num: ", *onlineNum)
	test.OnlineTest(*onlineNum)
	log.Warn("", "online test finish, online num: ", *onlineNum)
	select {}
}
