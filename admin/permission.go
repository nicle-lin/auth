package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

/*
许可
*/

type Permission struct {
	Id       int64 `json:"id,omitempty"`
	ObjectId int64 `json:"object_id,omitempty"`
	ScopeId  int64 `json:"scope_id,omitempty"`
}

func (a *DefaultAdmin) NewPermission(permission Permission, multiOrm ...orm.Ormer) (int64, error) {
	if permission.ObjectId == 0 {
		return 0, ErrInvalidObjectId
	}
	if permission.ScopeId == 0 {
		return 0, ErrInvalidScopeId
	}
	return insertPermission(permission, multiOrm...)
}

func (a *DefaultAdmin) UpdatePermission(permission Permission, multiOrm ...orm.Ormer) error {
	if permission.Id == 0 {
		return ErrInvalidActonId
	}
	if permission.ObjectId == 0 {
		return ErrInvalidObjectId
	}
	if permission.ScopeId == 0 {
		return ErrInvalidScopeId
	}
	return updatePermission(permission, multiOrm...)
}

func (a *DefaultAdmin) DeletePermission(permission Permission, multiOrm ...orm.Ormer) error {
	if permission.Id == 0 {
		return ErrInvalidPermissionId
	}
	return deletePermission(permission.Id, multiOrm...)
}

func (a *DefaultAdmin) PermissionInfo(id int64, multiOrm ...orm.Ormer) ([]Permission, error) {
	return permissionInfo(id, multiOrm...)
}

func (a *DefaultAdmin) MultiPermission(ids []int64, multiOrm ...orm.Ormer) ([]Permission, error) {
	return multiPermission(ids)
}

func multiPermission(ids []int64, multiOrm ...orm.Ormer) (permissions []Permission, err error) {
	if len(ids) < 0 {
		err = ErrRoleMissingPermission
		return
	}
	idStr := Join(ids, ",")
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id IN (%v)
	`, TablePermission, idStr)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&permissions)
	return
}

func permissionInfo(id int64, multiOrm ...orm.Ormer) (permissions []Permission, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		`, TablePermission, id)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&permissions)
	return
}

func deletePermission(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TablePermission, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertPermission(permission Permission, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"object_id": permission.ObjectId,
		"scope_id":  permission.ScopeId,
	}
	id, ok := CheckExist(o, TablePermission, data)
	if ok {
		return id, ErrPermissionAlreadyExist
	}

	insertData ,err := Struct2Map(permission)
	if err != nil{
		return 0,err
	}
	values, sql := InsertSql(TablePermission, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updatePermission(permission Permission, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, permission.Id)
	permission.Id = 0
	updateData,err  := Struct2Map(permission)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TablePermission, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
