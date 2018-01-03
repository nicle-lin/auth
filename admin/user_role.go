package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)


type UserRole struct {
	Id     int64 `json:"id,omitempty"`
	UserId int64 `json:"user_id,omitempty"`
	RoleId int64 `json:"role_id,omitempty"`
}

func (a *DefaultAdmin) NewUserRole(ur UserRole, multiOrm ...orm.Ormer) (int64, error) {
	if ur.UserId == 0 {
		return 0, ErrInvalidUserId
	}
	if ur.RoleId == 0 {
		return 0, ErrInvalidRoleId
	}

	return insertUserRole(ur, multiOrm...)
}

func (a *DefaultAdmin) UpdateUserRole(ur UserRole, multiOrm ...orm.Ormer) error {
	if ur.Id == 0 {
		return ErrInvalidUserRoleId
	}
	if ur.UserId == 0 {
		return ErrInvalidUserId
	}
	if ur.RoleId == 0 {
		return ErrInvalidRoleId
	}

	return updateUserRole(ur, multiOrm...)
}

func (a *DefaultAdmin) DeleteUserRole(ur UserRole, multiOrm ...orm.Ormer) error {
	if ur.Id == 0 {
		return ErrInvalidUserRoleId
	}
	return deleteUserRole(ur.Id, multiOrm...)
}

func (a *DefaultAdmin) UserRoleInfoByUserId(userId int64, multiOrm ...orm.Ormer) ([]UserRole, error) {
	return userRoleInfo(userId, TypeUser, multiOrm...)
}

func (a *DefaultAdmin) UserRoleInfoByRoleId(roleId int64, multiOrm ...orm.Ormer) ([]UserRole, error) {
	return userRoleInfo(roleId, TypeRole, multiOrm...)
}

func userRoleInfo(userRoleId int64, itemType ItemType, multiOrm ...orm.Ormer) (userRole []UserRole, err error) {
	var field = "role_id"
	if itemType == TypeUser {
		field = "user_id"
	}
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			%v = %v
		`, TableUserRole, field, userRoleId)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&userRole)
	return
}

func insertUserRole(ur UserRole, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"user_id": ur.UserId,
		"role_id": ur.RoleId,
	}
	id, ok := CheckExist(o, TableUserRole, data)
	if ok {
		return id, ErrUserRoleAlreadyExist
	}
	insertData, err := Struct2Map(ur)
	if err != nil{
		return 0, err
	}

	values, sql := InsertSql(TableUserRole, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateUserRole(ur UserRole, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, ur.Id)
	ur.Id = 0
	updateData,err := Struct2Map(ur)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableUserRole, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}

func deleteUserRole(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableUserRole, condition)
	_, err := o.Raw(sql).Exec()
	return err
}
