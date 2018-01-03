package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Remark       string `json:"remark,omitempty"`
	CreateTime   int64  `json:"create_time,omitempty"`
	CreateUserId int64  `json:"create_user_id"`
	IsDeleted    int64  `json:"is_deleted"`
}

func (a *DefaultAdmin) NewRole(role Role, multiOrm ...orm.Ormer) (int64, error) {
	if role.Name == "" {
		return 0, ErrRoleNameEmpty
	}
	return insertRole(role, multiOrm...)
}

func (a *DefaultAdmin) UpdateRole(role Role, multiOrm ...orm.Ormer) error {
	if role.Id == 0 {
		return ErrInvalidRoleId
	}
	return updateRole(role, multiOrm...)
}

func (a *DefaultAdmin) DeleteRole(role Role, multiOrm ...orm.Ormer) error {
	if role.Id == 0 {
		return ErrInvalidRoleId
	}
	return deleteRole(role.Id, multiOrm...)
}

func (a *DefaultAdmin) RoleInfo(id int64, multiOrm ...orm.Ormer) (Role, error) {
	return roleInfo(id, multiOrm...)
}

func roleInfo(id int64, multiOrm ...orm.Ormer) (role Role, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableRole, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&role)
	return
}

// delete role
func deleteRole(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableRole, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

//写入role表
func insertRole(role Role, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name": role.Name,
	}
	id, ok := CheckExist(o, TableRole, data)
	if ok {
		return id, ErrRoleAlreadyExist
	}

	insertData,err := Struct2Map(role)
	if err != nil{
		return 0,err
	}
	values, sql := InsertSql(TableRole, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateRole(role Role, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, role.Id)
	role.Id = 0
	updateData,err := Struct2Map(role)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableRole, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
