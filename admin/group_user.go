package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type GroupUser struct {
	Id      int64 `json:"id,omitempty"`
	UserId  int64 `json:"user_id,omitempty"`
	GroupId int64 `json:"group_id,omitempty"`
}

func (a *DefaultAdmin) NewGroupUser(gu GroupUser, multiOrm ...orm.Ormer) (int64, error) {
	if gu.UserId == 0 {
		return 0, ErrInvalidUserId
	}
	if gu.GroupId == 0 {
		return 0, ErrInvalidGroupId
	}

	return insertGroupUser(gu, multiOrm...)
}

func (a *DefaultAdmin) UpdateGroupUser(gu GroupUser, multiOrm ...orm.Ormer) error {
	if gu.Id == 0 {
		return ErrInvalidGroupUserId
	}
	if gu.UserId == 0 {
		return ErrInvalidUserId
	}
	if gu.GroupId == 0 {
		return ErrInvalidGroupUserId
	}

	return updateGroupUser(gu, multiOrm...)
}

func (a *DefaultAdmin) DeleteGroupUser(gu GroupUser, multiOrm ...orm.Ormer) error {
	if gu.Id == 0 {
		return ErrInvalidGroupUserId
	}
	return deleteGroupUser(gu.Id, multiOrm...)
}

func (a *DefaultAdmin) GroupUserInfo(id int64, multiOrm ...orm.Ormer) ([]GroupUser, error) {
	return groupUserInfo(id, multiOrm...)
}

func groupUserInfo(groupId int64, multiOrm ...orm.Ormer) (groupUser []GroupUser, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			group_id = %v
		`, TableGroupUser, groupId)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&groupUser)
	return
}

func insertGroupUser(gu GroupUser, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"user_id":  gu.UserId,
		"group_id": gu.GroupId,
	}
	id, ok := CheckExist(o, TableGroupUser, data)
	if ok {
		return id, ErrGroupUserAlreadyExist
	}

	insertData,err := Struct2Map(gu)
	if err != nil{
		return 0, err
	}
	values, sql := InsertSql(TableGroupUser, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateGroupUser(gu GroupUser, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, gu.Id)
	gu.Id = 0
	updateData ,err := Struct2Map(gu)
	if err != nil{
		return  err
	}
	values, sql := UpdateSql(TableGroupUser, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}

func deleteGroupUser(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableGroupUser, condition)
	_, err := o.Raw(sql).Exec()
	return err
}
