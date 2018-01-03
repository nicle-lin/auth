package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

/*
用户组
*/


type Department struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *DefaultAdmin) NewDepartment(department Department, multiOrm ...orm.Ormer) (int64, error) {
	if department.Name == "" {
		return 0, ErrDepartmentNameEmpty
	}
	return insertDepartment(department, multiOrm...)
}

func (a *DefaultAdmin) UpdateDepartment(department Department, multiOrm ...orm.Ormer) error {
	if department.Id == 0 {
		return ErrInvalidDepartmentId
	}
	if department.Name == "" {
		return ErrDepartmentNameEmpty
	}
	return updateDepartment(department, multiOrm...)
}

func (a *DefaultAdmin) DeleteDepartment(department Department, multiOrm ...orm.Ormer) error {
	if department.Id == 0 {
		return ErrInvalidDepartmentId
	}
	return deleteDepartment(department.Id, multiOrm...)
}

func (a *DefaultAdmin) DepartmentInfo(id int64, multiOrm ...orm.Ormer) (Department, error) {
	return departmentInfo(id, multiOrm...)
}

func departmentInfo(id int64, multiOrm ...orm.Ormer) (department Department, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableDepartment, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&department)
	return
}

func deleteDepartment(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableDepartment, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertDepartment(department Department, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name": department.Name,
	}
	id, ok := CheckExist(o, TableDepartment, data)
	if ok {
		return id, ErrDepartmentAlreadyExist
	}

	insertData,err := Struct2Map(department)
	if err != nil{
		return 0,err
	}
	values, sql := InsertSql(TableDepartment, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateDepartment(department Department, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, department.Id)
	department.Id = 0
	updateData,err := Struct2Map(department)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableDepartment, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}