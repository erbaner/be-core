package full

import (
	"github.com/erbaner/be-core/open_im_sdk_callback"
	"github.com/erbaner/be-core/pkg/common"
	"github.com/erbaner/be-core/pkg/log"
	"github.com/erbaner/be-core/pkg/sdk_params_callback"
	"github.com/erbaner/be-core/pkg/utils"
)

func (u *Full) GetUsersInfo(callback open_im_sdk_callback.Base, userIDList string, operationID string) {
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", userIDList)
		var unmarshalParam sdk_params_callback.GetUsersInfoParam
		common.JsonUnmarshalAndArgsValidate(userIDList, &unmarshalParam, callback, operationID)
		result := u.getUsersInfo(callback, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonStringDefault(result))
	}()
}
