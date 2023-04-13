package common

//import (
//	"github.com/mitchellh/mapstructure"
//	"github.com/erbaner/be-core/open_im_sdk_callback"
//	"github.com/erbaner/be-core/pkg/db"
//	"github.com/erbaner/be-core/pkg/db/model_struct"
//)
//
//func GetGroupMemberListByGroupID(callback open_im_sdk_callback.Base, operationID string, db *db.DataBase, groupID string) []*model_struct.LocalGroupMember {
//	memberList, err := db.GetGroupMemberListByGroupID(groupID)
//	CheckDBErrCallback(callback, err, operationID)
//	return memberList
//}
//
//func MapstructureDecode(input interface{}, output interface{}, callback open_im_sdk_callback.Base, oprationID string) {
//	err := mapstructure.Decode(input, output)
//	CheckDataErrCallback(callback, err, oprationID)
//}
