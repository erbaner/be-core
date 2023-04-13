package common

import (
	api "github.com/erbaner/be-core/pkg/server_api_params"
	"github.com/erbaner/be-core/pkg/utils"
	"github.com/golang/protobuf/proto"
)

func UnmarshalTips(msg *api.MsgData, detail proto.Message) error {
	var tips api.TipsComm
	if err := proto.Unmarshal(msg.Content, &tips); err != nil {
		return utils.Wrap(err, "")
	}
	if err := proto.Unmarshal(tips.Detail, detail); err != nil {
		return utils.Wrap(err, "")
	}
	return nil
}
