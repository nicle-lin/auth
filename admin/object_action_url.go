package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)


type ObjectActionUrl struct {
	Id             int64 `json:"id,omitempty"`
	ObjectActionId int64 `json:"object_action_id,omitempty"`
	UrlId          int64 `json:"url_id,omitempty"`
}

func (a *DefaultAdmin) NewObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) (int64, error) {
	if ur.ObjectActionId == 0 {
		return 0, ErrInvalidObjectActionId
	}
	if ur.UrlId == 0 {
		return 0, ErrInvalidNodeId
	}

	return insertObjectActionUrl(ur, multiOrm...)
}

func (a *DefaultAdmin) UpdateObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) error {
	if ur.Id == 0 {
		return ErrInvalidObjectActionUrlId
	}
	if ur.ObjectActionId == 0 {
		return ErrInvalidObjectActionId
	}
	if ur.UrlId == 0 {
		return ErrInvalidNodeId
	}

	return updateObjectActionUrl(ur,multiOrm...)
}

func (a *DefaultAdmin) DeleteObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) error {
	if ur.Id == 0 {
		return ErrInvalidObjectActionUrlId
	}
	return deleteObjectActionUrl(ur.Id, multiOrm...)
}

func (a *DefaultAdmin) ObjectActionUrlInfo(objectActionId int64, multiOrm ...orm.Ormer) ([]ObjectActionUrl, error) {
	return ObjectActionUrlInfo(objectActionId, multiOrm...)
}

func (a *DefaultAdmin) MultiObjectActionUrl(objectActionsId []int64, multiOrm ...orm.Ormer) ([]ObjectActionUrl, error) {
	return multiObjectActionUrl(objectActionsId, multiOrm...)
}

func multiObjectActionUrl(objectActionsId []int64, multiOrm ...orm.Ormer) (objectActionUrls []ObjectActionUrl, err error) {
	if len(objectActionsId) < 0 {
		err = ErrObjectActionMissingUrl
		return
	}
	objectActionUrlsIdstr := Join(objectActionsId, ",")
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			object_action_id IN (%v)
		`, TableObjectActionUrl, objectActionUrlsIdstr)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&objectActionUrls)
	return
}


func ObjectActionUrlInfo(objectActionId int64, multiOrm ...orm.Ormer) (objectActionUrl []ObjectActionUrl, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			object_action_id = %v
		`, TableObjectActionUrl, objectActionId)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&objectActionUrl)
	return
}

func insertObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"object_action_id": ur.ObjectActionId,
		"url_id":           ur.UrlId,
	}
	id, ok := CheckExist(o, TableObjectActionUrl, data)
	if ok {
		return id, ErrObjectActionUrlAlreadyExist
	}

	insertData, err := Struct2Map(ur)
	if err != nil{
		return 0, err
	}

	values, sql := InsertSql(TableObjectActionUrl, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, ur.Id)
	ur.Id = 0
	updateData,err := Struct2Map(ur)
	if err != nil {
		return err
	}
	values, sql := UpdateSql(TableObjectActionUrl, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}

func deleteObjectActionUrl(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableObjectActionUrl, condition)
	_, err := o.Raw(sql).Exec()
	return err
}
