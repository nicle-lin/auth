package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type RolePermission struct {
	Id           int64 `json:"id,omitempty"`
	RoleId       int64 `json:"role_id,omitempty"`
	PermissionId int64 `json:"permission_id,omitempty"`
}

func (a *DefaultAdmin) NewRolePermission(rp RolePermission, multiOrm ...orm.Ormer) (int64, error) {
	if rp.RoleId == 0 {
		return 0, ErrInvalidRoleId
	}
	if rp.PermissionId == 0 {
		return 0, ErrInvalidPermissionId
	}

	return insertRolePermission(rp, multiOrm...)
}

func (a *DefaultAdmin) UpdateRolePermission(rp RolePermission, multiOrm ...orm.Ormer) error {
	if rp.Id == 0 {
		return ErrInvalidRolePermissionId
	}
	if rp.RoleId == 0 {
		return ErrInvalidRoleId
	}
	if rp.PermissionId == 0 {
		return ErrInvalidPermissionId
	}

	return updateRolePermission(rp,multiOrm...)
}

func (a *DefaultAdmin) DeleteRolePermission(rp RolePermission, multiOrm ...orm.Ormer) error {
	if rp.Id == 0 {
		return ErrInvalidRolePermissionId
	}
	return deleteRolePermission(rp.Id)
}

func (a *DefaultAdmin) RolePermissionInfo(roleId int64, multiOrm ...orm.Ormer) ([]RolePermission, error) {
	return rolePermissionInfo(roleId, multiOrm...)
}

func (a *DefaultAdmin) MultiRolePermission(rolesId []int64, multiOrm ...orm.Ormer) ([]RolePermission, error) {
	return multiRolePermission(rolesId, multiOrm...)
}

func multiRolePermission(rolesId []int64, multiOrm ...orm.Ormer) (rolePermissions []RolePermission, err error) {
	if len(rolesId) < 0 {
		err = ErrRoleMissingPermission
		return
	}
	rolesIdstr := Join(rolesId, ",")
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			role_id IN  (%v)
		`, TableRolePermission, rolesIdstr)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&rolePermissions)
	return
}

func rolePermissionInfo(roleId int64, multiOrm ...orm.Ormer) (rolePermission []RolePermission, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			role_id = %v
		`, TableRolePermission, roleId)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&rolePermission)
	return
}

func insertRolePermission(rp RolePermission, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"role_id":       rp.RoleId,
		"permission_id": rp.PermissionId,
	}
	id, ok := CheckExist(o, TableRolePermission, data)
	if ok {
		return id, ErrRolePermissionAlreadyExist
	}

	insertData, err := Struct2Map(rp)
	if err != nil{
		return 0, err
	}

	values, sql := InsertSql(TableRolePermission, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateRolePermission(rp RolePermission, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, rp.Id)
	rp.Id = 0
	updateData,err := Struct2Map(rp)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableRolePermission, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}

func deleteRolePermission(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableRolePermission, condition)
	_, err := o.Raw(sql).Exec()
	return err
}
