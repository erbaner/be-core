package main

import (
	"encoding/json"
	"time"

	"github.com/erbaner/be-core/pkg/log"
	"github.com/erbaner/be-core/pkg/network"
	"github.com/erbaner/be-core/pkg/server_api_params"
	"github.com/erbaner/be-core/pkg/utils"
	"github.com/erbaner/be-core/test"
)

var (
	//APIADDR = "http://43.155.69.205:10002"
	//WSADDR  = "ws://43.155.69.205:10001"
	//APIADDR = "https://chat-api-dev.opencord.so"
	//WSADDR  = "wss://chat-ws-dev.opencord.so"
	APIADDR = "http://125.124.195.201:10002"
	WSADDR  = "ws://125.124.195.201:10001"
	//APIADDR      = "http://113.108.8.93:10002"
	//WSADDR       = "ws://113.108.8.93:10001"
	REGISTERADDR = APIADDR + "/user_register"
	ACCOUNTCHECK = APIADDR + "/manager/account_check"
	TOKENADDR    = APIADDR + "/auth/user_token"
	SECRET       = "tuoyun"
	//SECRET       = "4zbF9Y6Fs1QJ0hsmpC3B676txZcCnjcZ"
	SENDINTERVAL = 20
)

const PlatformID = 1

type ResToken struct {
	Data struct {
		ExpiredTime int64  `json:"expiredTime"`
		Token       string `json:"token"`
		Uid         string `json:"uid"`
	}
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

func ggetToken(uid string) string {
	url := TOKENADDR
	var req server_api_params.UserTokenReq
	req.Platform = PlatformID
	req.UserID = uid
	req.Secret = SECRET
	req.OperationID = utils.OperationIDGenerator()
	r, err := network.Post2Api(url, req, "a")
	if err != nil {
		log.Error(req.OperationID, "Post2Api failed ", err.Error(), url, req)
		return ""
	}
	var stcResp ResToken
	err = json.Unmarshal(r, &stcResp)
	if stcResp.ErrCode != 0 {
		log.Error(req.OperationID, "ErrCode failed ", stcResp.ErrCode, stcResp.ErrMsg, url, req)
		return ""
	}
	log.Info(req.OperationID, "get token: ", stcResp.Data.Token)
	return stcResp.Data.Token
}

func gRunGetToken(strMyUid string) string {
	var token string
	for true {
		token = ggetToken(strMyUid)
		if token == "" {
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		} else {
			break
		}
	}
	return token
}
func main() {
	uid := "7788"
	//Gordon
	//uid:="1554321956297519104"
	//Gordon2
	//uid := "1583984945064968192"
	//uid := "3734595565"
	tokenx := gRunGetToken(uid)
	//tokenx := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI3MDcwMDgxNTMiLCJQbGF0Zm9ybSI6IkFuZHJvaWQiLCJleHAiOjE5NjY0MTJ1XjJZGWj5fB3mqC7p6ytxSarvxZfsABwIjoxNjUxMDU1MDU2fQ.aWvmJ_sQxXmT5nKwiM5QsF9-tfkldzOYZtRD3nrUuko"
	test.InOutDoTest(uid, tokenx, WSADDR, APIADDR)
	test.DoTestGetAdvancedHistoryMessageList()
	//test.DoTestSendMsg2GroupWithMessage(uid, "1623878302774460418", "2")
	//test.DoTestAddMessageReactionExtensions(1,"special handle")
	//time.Sleep(time.Second*5)
	//test.DoTestAddMessageReactionExtensions(2,"special handle")
	//time.Sleep(time.Second*5)
	//test.DoTestGetMessageListReactionExtensions("special handle")
	//test.DoTestSetAppBadge()
	//test.DoTestSearchLocalMessages()
	//test.DoTestGetAdvancedHistoryMessageList()
	println("start")
	//test.DoTestGetUserInDepartment()
	//test.DoTestGetDepartmentMemberAndSubDepartment()
	//test.DoTestDeleteAllMsgFromLocalAndSvr()
	//	test.DoTestGetDepartmentMemberAndSubDepartment()
	//test.DotestUploadFile()
	//test.DotestMinio()
	//test.DotestSearchFriends()
	//if *senderNum == 0 {
	//	test.RegisterAccounts(*onlineNum)
	//	return
	//}
	//
	//test.OnlineTest(*onlineNum)
	////test.TestSendCostTime()
	//test.ReliabilityTest(*singleSenderMsgNum, *intervalTime, 10, *senderNum)
	//test.DoTestSearchLocalMessages()
	//println("start")
	//test.DoTestSendImageMsg(strMyUidx, "17726378428")
	//test.DoTestSearchGroups()
	//test.DoTestGetHistoryMessage("")
	//test.DoTestGetHistoryMessageReverse("")
	//test.DoTestInviteInGroup()
	//test.DoTestCancel()
	//test.DoTestSendMsg2(strMyUidx, friendID)
	//test.DoTestGetAllConversation()

	//test.DoTestGetOneConversation("17726378428")
	//test.DoTestGetConversations(`["single_17726378428"]`)
	//test.DoTestGetConversationListSplit()
	//test.DoTestGetConversationRecvMessageOpt(`["single_17726378428"]`)

	//set batch
	//test.DoTestSetConversationRecvMessageOpt([]string{"single_17726378428"}, constant.NotReceiveMessage)
	//set one
	////set batch
	//test.DoTestSetConversationRecvMessageOpt([]string{"single_17726378428"}, constant.ReceiveMessage)
	////set one
	//test.DoTestSetConversationPinned("single_17726378428", false)
	//test.DoTestSetOneConversationRecvMessageOpt("single_17726378428", constant.NotReceiveMessage)
	//test.DoTestSetOneConversationPrivateChat("single_17726378428", false)
	//test.DoTestReject()
	//test.DoTestAccept()
	//test.DoTestMarkGroupMessageAsRead()
	//test.DoTestGetGroupHistoryMessage()
	//test.DoTestGetHistoryMessage("17396220460")
	time.Sleep(250000 * time.Millisecond)
	//b := utils.GetCurrentTimestampBySecond()
	i := 0
	for {
		//test.DoTestSendMsg2Group(strMyUidx, "42c9f515cb84ee0e82b3f3ce71eb14d6", i)
		i++
		time.Sleep(250 * time.Millisecond)
		//if i == 100 {
		//	break
		//}
		log.Warn("", "10 * time.Millisecond ###################waiting... msg: ", i)
	}
	//
	//log.Warn("", "cost time: ", utils.GetCurrentTimestampBySecond()-b)
	//return
	//i = 0
	//for {
	//	//test.DoTestSendMsg2Group(strMyUidx, "42c9f515cb84ee0e82b3f3ce71eb14d6", i)
	//	i++
	//	time.Sleep(1000 * time.Millisecond)
	//	if i == 10 {
	//		break
	//	}
	//	log.Warn("", "1000 * time.Millisecond ###################waiting... msg: ", i)
	//}
	//
	//i = 0
	//for {
	//	test.DoTestSendMsg2Group(strMyUidx, "42c9f515cb84ee0e82b3f3ce71eb14d6", i)
	//	i++
	//	time.Sleep(10000 * time.Millisecond)
	//	log.Warn("", "10000 * time.Millisecond ###################waiting... msg: ", i)
	//}

	//reliabilityTest()
	//	test.PressTest(testClientNum, intervalSleep, imIP)
}

//
//func main() {
//	testClientNum := 100
//	intervalSleep := 2
//	imIP := "43.128.5.63"

//
//	msgNum := 1000
//	test.ReliabilityTest(msgNum, intervalSleep, imIP)
//	for i := 0; i < 6; i++ {
//		test.Msgwg.Wait()
//	}
//
//	for {
//
//		if test.CheckReliabilityResult() {
//			log.Warn("CheckReliabilityResult ok, again")
//
//		} else {
//			log.Warn("CheckReliabilityResult failed , wait.... ")
//		}
//
//		time.Sleep(time.Duration(10) * time.Second)
//	}
//
//}

//func printCallerNameAndLine() string {
//	pc, _, line, _ := runtime.Caller(2)
//	return runtime.FuncForPC(pc).Name() + "()@" + strconv.Itoa(line) + ": "
//}

// myuid,  maxuid,  msgnum
