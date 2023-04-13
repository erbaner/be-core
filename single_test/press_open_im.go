package main

import (
	"flag"

	"github.com/erbaner/be-core/pkg/constant"
	"github.com/erbaner/be-core/pkg/log"
	"github.com/erbaner/be-core/test"
)

func main() {

	var senderNum *int          //Number of users sending messages
	var singleSenderMsgNum *int //Number of single user send messages
	var intervalTime *int       //Sending time interval, in millisecond
	senderNum = flag.Int("sn", 10, "sender num")
	singleSenderMsgNum = flag.Int("mn", 10, "single sender msg num")
	intervalTime = flag.Int("t", 1000, "interval time mill second")
	flag.Parse()
	constant.OnlyForTest = 1
	log.NewPrivateLog("", test.LogLevel)
	log.Warn("", "press test begin, sender num: ", *senderNum, " single sender msg num: ", *singleSenderMsgNum, " send msg total num: ", *senderNum**singleSenderMsgNum)
	test.PressTest(*singleSenderMsgNum, *intervalTime, *senderNum)
	select {}
}
