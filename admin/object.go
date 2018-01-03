package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

/*
操作对象集
*/

type Object struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *DefaultAdmin) NewObject(object Object, multiOrm ...orm.Ormer) (int64, error) {
	if object.Name == "" {
		return 0, ErrObjectNameEmpty
	}
	return insertObject(object, multiOrm...)
}

func (a *DefaultAdmin) UpdateObject(object Object, multiOrm ...orm.Ormer) error {
	if object.Id == 0 {
		return ErrInvalidObjectId
	}
	if object.Name == "" {
		return ErrObjectNameEmpty
	}
	return updateObject(object, multiOrm...)
}

func (a *DefaultAdmin) DeleteObject(object Object, multiOrm ...orm.Ormer) error {
	if object.Id == 0 {
		return ErrInvalidObjectId
	}
	return deleteObject(object.Id, multiOrm...)
}

func (a *DefaultAdmin) ObjectInfo(id int64, multiOrm ...orm.Ormer) (Object, error) {
	return objectInfo(id, multiOrm...)
}

func objectInfo(id int64, multiOrm ...orm.Ormer) (object Object, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableObject, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&object)
	return
}

func deleteObject(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableObject, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertObject(object Object, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name": object.Name,
	}
	id, ok := CheckExist(o, TableObject, data)
	if ok {
		return id, ErrObjectAlreadyExist
	}

	insertData, err := Struct2Map(object)
	if err != nil {
		return 0, err
	}
	values, sql := InsertSql(TableObject, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateObject(object Object, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, object.Id)
	object.Id = 0
	updateData, err := Struct2Map(object)
	if err != nil {
		return err
	}
	values, sql := UpdateSql(TableObject, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
