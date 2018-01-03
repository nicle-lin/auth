package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Node struct {
	Id   int64  `json:"id,omitempty"`
	Url  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *DefaultAdmin) NewNode(node Node, multiOrm ...orm.Ormer) (int64, error) {
	if node.Name == "" {
		return 0, ErrNodeNameEmpty
	}
	if node.Url == "" {
		return 0, ErrNodeURLEmpty
	}
	return insertNode(node, multiOrm...)
}

func (a *DefaultAdmin) UpdateNode(node Node, multiOrm ...orm.Ormer) error {
	if node.Id == 0 {
		return ErrInvalidNodeId
	}
	if node.Name == "" {
		return ErrNodeNameEmpty
	}
	if node.Url == "" {
		return ErrNodeURLEmpty
	}
	return updateNode(node, multiOrm...)
}

func (a *DefaultAdmin) DeleteNode(node Node, multiOrm ...orm.Ormer) error {
	if node.Id == 0 {
		return ErrInvalidNodeId
	}
	return deleteNode(node.Id, multiOrm...)
}

func (a *DefaultAdmin) NodeInfo(id int64, multiOrm ...orm.Ormer) (Node, error) {
	return nodeInfo(id, multiOrm...)
}

func (a *DefaultAdmin) NodeInfoByUrl(url string, multiOrm ...orm.Ormer) (Node, error) {
	return nodeInfoByUrl(url, multiOrm...)
}

func (a *DefaultAdmin) AllNodeInfo(multiOrm ...orm.Ormer) ([]Node, error) {
	return allNodeInfo(multiOrm...)
}

func allNodeInfo(multiOrm ...orm.Ormer) (nodes []Node, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
	`, TableNode)
	o := NewOrm(multiOrm)
	_, err = o.Raw(sql).QueryRows(&nodes)
	return
}

func nodeInfoByUrl(url string, multiOrm ...orm.Ormer) (node Node, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			url = ?
		LIMIT 1
	`, TableNode)
	o := NewOrm(multiOrm)
	err = o.Raw(sql,url).QueryRow(&node)
	return
}

func nodeInfo(id int64, multiOrm ...orm.Ormer) (node Node, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			id = %v
		LIMIT 1
		`, TableNode, id)
	o := NewOrm(multiOrm)
	err = o.Raw(sql).QueryRow(&node)
	return
}

func deleteNode(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableNode, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

func insertNode(node Node, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"name": node.Name,
		"url":  node.Url,
	}
	id, ok := CheckExist(o, TableNode, data)
	if ok {
		return id, ErrNodeAlreadyExist
	}

	insertData,err := Struct2Map(node)
	if err != nil{
		return 0, err
	}
	values, sql := InsertSql(TableNode, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateNode(node Node, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, node.Id)
	node.Id = 0
	updateData ,err := Struct2Map(node)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableNode, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
