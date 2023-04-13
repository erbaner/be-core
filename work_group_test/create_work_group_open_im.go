package main

import (
	"flag"

	"github.com/erbaner/be-core/pkg/log"
	"github.com/erbaner/be-core/test"
)

func main() {
	var groupMemberNumber *int
	groupMemberNumber = flag.Int("gmn", 1000, "group member number ")
	flag.Parse()
	log.NewPrivateLog("", test.LogLevel)
	log.Warn("", "CreateWorkGroup  start, group member number: ", *groupMemberNumber)
	*groupMemberNumber = *groupMemberNumber + 2

	test.CreateWorkGroup(*groupMemberNumber)
	log.Warn("", "CreateWorkGroup finish, group member number: ", *groupMemberNumber+1)

}
