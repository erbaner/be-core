package db

import (
	"github.com/erbaner/be-core/pkg/constant"
	"github.com/erbaner/be-core/pkg/db/model_struct"
	"github.com/erbaner/be-core/pkg/utils"
)

func (d *DataBase) GetJoinedWorkingGroupIDList() ([]string, error) {
	groupList, err := d.GetJoinedGroupListDB()
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var groupIDList []string
	for _, v := range groupList {
		if v.GroupType == constant.WorkingGroup {
			groupIDList = append(groupIDList, v.GroupID)
		}
	}
	return groupIDList, nil
}

func (d *DataBase) GetJoinedWorkingGroupList() ([]*model_struct.LocalGroup, error) {
	groupList, err := d.GetJoinedGroupListDB()
	var transfer []*model_struct.LocalGroup
	for _, v := range groupList {
		if v.GroupType == constant.WorkingGroup {
			transfer = append(transfer, v)
		}
	}
	return transfer, utils.Wrap(err, "GetJoinedSuperGroupList failed ")
}
