package business

import (
	"github.com/erbaner/be-core/open_im_sdk_callback"
	"github.com/erbaner/be-core/pkg/db/db_interface"
	"github.com/erbaner/be-core/pkg/log"
	"github.com/erbaner/be-core/pkg/utils"
)

type Business struct {
	listener open_im_sdk_callback.OnCustomBusinessListener
	db       db_interface.DataBase
}

func NewBusiness(db db_interface.DataBase) *Business {
	return &Business{
		db: db,
	}
}

func (b *Business) DoNotification(jsonDetailStr string, operationID string) {
	if b.listener == nil {
		log.NewDebug(operationID, "WorkMoments listener is null", jsonDetailStr)
		return
	}
	log.NewInfo(operationID, utils.GetSelfFuncName(), "json_detail: ", jsonDetailStr)
	b.listener.OnRecvCustomBusinessMessage(jsonDetailStr)
}
