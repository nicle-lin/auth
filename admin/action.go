package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)



type Action struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *DefaultAdmin) NewAction(action Action, multiOrm ...orm.Ormer) (int64, error) {
	if action.Name == "" {
		return 0, ErrActionNameEmpty
	}
	return insertAction(action, multiOrm...)
}

func (a *DefaultAdmin) UpdateAction(action Action, multiOrm ...orm.Ormer) error {
	if action.Id == 0 {
		return ErrInvalidActonId
	}
	if action.Name == "" {
		return ErrUsernameEmpty
	}
	return updateAction(action, multiOrm...)
}

func (a *DefaultAdmin) DeleteAction(action Action, multiOrm ...orm.Ormer) error {
	if action.Id == 0 {
		return ErrInvalidActonId
	}
	return deleteAction(action.Id, multiOrm...)
}

func (a *DefaultAdmin) ActionInfo(id int64, multiOrm ...orm.Ormer) (Action, error) {
	return actionInfo(id, multiOrm...)
}

func actionInfo(id int64, multiOrm ...orm.Ormer) (action Action, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableAction, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&action)
	return
}

func deleteAction(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableAction, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertAction(action Action, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name": action.Name,
	}
	id, ok := CheckExist(o, TableAction, data)
	if ok {
		return id, ErrActionAlreadyExist
	}

	insertData,err := Struct2Map(action)
	if err != nil{
		return 0,err
	}
	values, sql := InsertSql(TableAction, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateAction(action Action, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, action.Id)
	action.Id = 0
	updateData,err := Struct2Map(action)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableAction, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
