package db

import (
	"github.com/erbaner/be-core/pkg/db/model_struct"
	"github.com/erbaner/be-core/pkg/utils"
)

func (d *DataBase) BatchInsertTempCacheMessageList(MessageList []*model_struct.TempCacheLocalChatLog) error {
	if MessageList == nil {
		return nil
	}
	return utils.Wrap(d.conn.Create(MessageList).Error, "BatchInsertTempCacheMessageList failed")
}
func (d *DataBase) InsertTempCacheMessage(Message *model_struct.TempCacheLocalChatLog) error {

	return utils.Wrap(d.conn.Create(Message).Error, "InsertTempCacheMessage failed")

}
