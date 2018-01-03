package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type DepartmentGroup struct {
	Id           int64 `json:"id,omitempty"`
	DepartmentId int64 `json:"department_id,omitempty"`
	GroupId      int64 `json:"group_id,omitempty"`
}

func (a *DefaultAdmin) NewDepartmentGroup(gu DepartmentGroup, multiOrm ...orm.Ormer) (int64, error) {
	if gu.DepartmentId== 0 {
		return 0, ErrInvalidDepartmentId
	}
	if gu.GroupId == 0 {
		return 0, ErrInvalidGroupId
	}

	return insertDepartmentGroup(gu, multiOrm...)
}

func (a *DefaultAdmin) UpdateDepartmentGroup(gu DepartmentGroup, multiOrm ...orm.Ormer) error {
	if gu.Id == 0 {
		return ErrInvalidDepartmentGroupId
	}
	if gu.DepartmentId== 0 {
		return ErrInvalidDepartmentId
	}
	if gu.GroupId == 0 {
		return ErrInvalidDepartmentGroupId
	}

	return updateDepartmentGroup(gu, multiOrm...)
}

func (a *DefaultAdmin) DeleteDepartmentGroup(gu DepartmentGroup, multiOrm ...orm.Ormer) error {
	if gu.Id == 0 {
		return ErrInvalidDepartmentGroupId
	}
	return deleteDepartmentGroup(gu.Id, multiOrm...)
}

func (a *DefaultAdmin) DepartmentGroupInfo(departmentId int64, multiOrm ...orm.Ormer) ([]DepartmentGroup, error) {
	return departmentGroupInfo(departmentId, multiOrm...)
}

func departmentGroupInfo(departmentId int64, multiOrm ...orm.Ormer) (departmentGroup []DepartmentGroup, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			department_id = %v
		`, TableDepartmentGroup, departmentId)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&departmentGroup)
	return
}

func insertDepartmentGroup(depGroup DepartmentGroup, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"department_id": depGroup.DepartmentId,
		"group_id": depGroup.GroupId,
	}
	id, ok := CheckExist(o, TableDepartmentGroup, data)
	if ok {
		return id, ErrDepartmentGroupAlreadyExist
	}

	insertData,err := Struct2Map(depGroup)
	if err != nil{
		return 0, err
	}
	values, sql := InsertSql(TableDepartmentGroup, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateDepartmentGroup(depGroup DepartmentGroup, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, depGroup.Id)
	depGroup.Id = 0
	updateData,err := Struct2Map(depGroup)
	values, sql := UpdateSql(TableDepartmentGroup, updateData, condition)
	if err != nil{
		return err
	}
	_, err = o.Raw(sql, values).Exec()
	return err
}

func deleteDepartmentGroup(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableDepartmentGroup, condition)
	_, err := o.Raw(sql).Exec()
	return err
}
