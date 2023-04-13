package indexdb

import (
	"github.com/erbaner/be-core/pkg/db/model_struct"
	"github.com/erbaner/be-core/pkg/utils"
)

type LocalCacheMessage struct {
}

func (i *LocalCacheMessage) BatchInsertTempCacheMessageList(MessageList []*model_struct.TempCacheLocalChatLog) error {
	_, err := Exec(utils.StructToJsonString(MessageList))
	return err
}

func (i *LocalCacheMessage) InsertTempCacheMessage(Message *model_struct.TempCacheLocalChatLog) error {
	_, err := Exec(utils.StructToJsonString(Message))
	return err
}
