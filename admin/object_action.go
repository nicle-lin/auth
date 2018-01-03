package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ObjectAction struct {
	Id       int64 `json:"id,omitempty"`
	ObjectId int64 `json:"object_id,omitempty"`
	ActionId int64 `json:"action_id,omitempty"`
}

func (a *DefaultAdmin) NewObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) (int64, error) {
	if ur.ObjectId == 0 {
		return 0, ErrInvalidObjectId
	}
	if ur.ActionId == 0 {
		return 0, ErrInvalidActonId
	}

	return insertObjectAction(ur, multiOrm...)
}

func (a *DefaultAdmin) UpdateObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) error {
	if ur.Id == 0 {
		return ErrInvalidObjectActionId
	}
	if ur.ObjectId == 0 {
		return ErrInvalidObjectId
	}
	if ur.ActionId == 0 {
		return ErrInvalidActonId
	}

	return updateObjectAction(ur, multiOrm...)
}

func (a *DefaultAdmin) DeleteObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) error {
	if ur.Id == 0 {
		return ErrInvalidObjectActionId
	}
	return deleteObjectAction(ur.Id, multiOrm...)
}

func (a *DefaultAdmin) ObjectActionInfo(objectId int64, multiOrm ...orm.Ormer) ([]ObjectAction, error) {
	return objectActionInfo(objectId, multiOrm...)
}

func (a *DefaultAdmin) MultiObjectAction(objectsId []int64, multiOrm ...orm.Ormer) ([]ObjectAction, error) {
	return multiObjectAction(objectsId, multiOrm...)
}

func multiObjectAction(objectsId []int64, multiOrm ...orm.Ormer) (objectActions []ObjectAction, err error) {
	if len(objectsId) < 0 {
		err = ErrObjectMissingAction
		return
	}
	objectsIdstr := Join(objectsId, ",")
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			object_id IN (%v)
		`, TableObjectAction, objectsIdstr)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&objectActions)
	return
}

func objectActionInfo(objectId int64, multiOrm ...orm.Ormer) (objectAction []ObjectAction, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			object_id = %v
		`, TableObjectAction, objectId)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&objectAction)
	return
}

func insertObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"object_id": ur.ObjectId,
		"action_id": ur.ActionId,
	}
	id, ok := CheckExist(o, TableObjectAction, data)
	if ok {
		return id, ErrObjectActionAlreadyExist
	}
	insertData, err := Struct2Map(ur)
	if err != nil{
		return 0, err
	}

	values, sql := InsertSql(TableObjectAction, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, ur.Id)
	ur.Id = 0
	updateData,err := Struct2Map(ur)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableObjectAction, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}

func deleteObjectAction(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableObjectAction, condition)
	_, err := o.Raw(sql).Exec()
	return err
}
