package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

/*
用户组
*/


type Group struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *DefaultAdmin) NewGroup(group Group, multiOrm ...orm.Ormer) (int64, error) {
	if group.Name == "" {
		return 0, ErrGroupNameEmpty
	}
	return insertGroup(group, multiOrm...)
}

func (a *DefaultAdmin) UpdateGroup(group Group, multiOrm ...orm.Ormer) error {
	if group.Id == 0 {
		return ErrInvalidGroupId
	}
	if group.Name == "" {
		return ErrGroupNameEmpty
	}
	return updateGroup(group, multiOrm...)
}

func (a *DefaultAdmin) DeleteGroup(group Group, multiOrm ...orm.Ormer) error {
	if group.Id == 0 {
		return ErrInvalidGroupId
	}
	return deleteGroup(group.Id, multiOrm...)
}

func (a *DefaultAdmin) GroupInfo(id int64, multiOrm ...orm.Ormer) (Group, error) {
	return GroupInfo(id, multiOrm...)
}

func GroupInfo(id int64, multiOrm ...orm.Ormer) (group Group, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableGroup, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&group)
	return
}

func deleteGroup(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableGroup, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertGroup(group Group, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name": group.Name,
	}
	id, ok := CheckExist(o, TableGroup, data)
	if ok {
		return id, ErrGroupAlreadyExist
	}

	insertData,err := Struct2Map(group)
	if err != nil{
		return 0, err
	}
	values, sql := InsertSql(TableGroup, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateGroup(group Group, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, group.Id)
	group.Id = 0
	updateData,err := Struct2Map(group)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableGroup, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
