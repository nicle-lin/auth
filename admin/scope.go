package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)


type Scope struct {
	Id     int64  `json:"id,omitempty"`
	IdcId  int64  `json:"idc_id,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (a *DefaultAdmin) NewScope(scope Scope, multiOrm ...orm.Ormer) (int64, error) {
	if scope.Name == "" {
		return 0, ErrScopeNameEmpty
	}
	if scope.IdcId == 0 {
		return 0, ErrScopeIdcIdEmpty
	}
	return insertScope(scope, multiOrm...)
}

func (a *DefaultAdmin) UpdateScope(scope Scope, multiOrm ...orm.Ormer) error {
	if scope.Id == 0 {
		return ErrInvalidScopeId
	}
	if scope.Name == "" {
		return ErrScopeNameEmpty
	}
	if scope.IdcId == 0 {
		return ErrScopeIdcIdEmpty
	}
	return updateScope(scope, multiOrm...)
}

func (a *DefaultAdmin) DeleteScope(scope Scope, multiOrm ...orm.Ormer) error {
	if scope.Id == 0 {
		return ErrInvalidScopeId
	}
	return deleteScope(scope.Id, multiOrm...)
}

func (a *DefaultAdmin) ScopeInfo(id int64, multiOrm ...orm.Ormer) (Scope, error) {
	return ScopeInfo(id, multiOrm...)
}

func ScopeInfo(id int64, multiOrm ...orm.Ormer) (scope Scope, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableScope, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&scope)
	return
}

func deleteScope(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableScope, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertScope(scope Scope, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name":    scope.Name,
		"idc_id":  scope.IdcId,
		"user_id": scope.UserId,
	}
	id, ok := CheckExist(o, TableScope, data)
	if ok {
		return id, ErrScopeAlreadyExist
	}

	insertData,err := Struct2Map(scope)
	if err != nil{
		return 0, err
	}
	values, sql := InsertSql(TableScope, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateScope(scope Scope, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, scope.Id)
	scope.Id = 0
	updateData,err:= Struct2Map(scope)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableScope, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
