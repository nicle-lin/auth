package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
这是权限模块
created time: 2017-12-20 15:06
*/

type Account struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	Tel           string `json:"tel,omitempty"`
	Email         string `json:"email,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	LoginCount    int64  `json:"login_count,omitempty"`
	CreateTime    int64  `json:"create_time,omitempty"`
	LastLoginTime int64  `json:"last_login_time,omitempty"`
	IsFrozen      int64  `json:"is_frozen,omitempty"`
	IsDeleted     int64  `json:"is_deleted,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

func (a *DefaultAdmin) NewUser(account Account, multiOrm ...orm.Ormer) (int64, error) {
	if account.Name == "" {
		return 0, ErrUsernameEmpty
	}
	if account.Password == "" {
		return 0, ErrPasswordEmpty
	}
	if account.Tel == ""{
		return 0,ErrTelephoneEmpty
	}
	account.Password = Md5sum(account.Password)
	account.CreateTime = time.Now().Unix()
	return insertUser(account, multiOrm...)
}

func (a *DefaultAdmin) UpdateUser(account Account, multiOrm ...orm.Ormer) error {
	if account.Id == 0 {
		return ErrInvalidUserId
	}
	if account.Password != "" {
		account.Password = Md5sum(account.Password)
	}
	return updateUser(account, multiOrm...)
}

func (a *DefaultAdmin) DeleteUser(account Account, multiOrm ...orm.Ormer) error {
	if account.Id == 0 {
		return ErrInvalidUserId
	}
	return deleteUser(account.Id, multiOrm...)
}

func (a *DefaultAdmin) UserInfo(name string, multiOrm ...orm.Ormer) (Account, error) {
	return userInfo(name, multiOrm...)
}

func userInfo(username string, multiOrm ...orm.Ormer) (account Account, err error) {
	sql := fmt.Sprintf(`
		SELECT
			*
		FROM
			%v
		WHERE
			name = ?
		LIMIT 1
		`, TableUser)
	o := NewOrm(multiOrm)
	err = o.Raw(sql,username).QueryRow(&account)
	return
}

// delete user
func deleteUser(id int64, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf("id=%v", id)
	sql := DeleteSql(TableUser, condition)
	_, err := o.Raw(sql).Exec()
	return err
}

// write it to the table auth_user
func insertUser(account Account, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm)
	//check if exist
	data := map[string]interface{}{
		"tel": account.Tel,
	}
	id, ok := CheckExist(o, TableUser, data)
	if ok {
		return id, ErrUserAlreadyExist
	}

	insertData,err:= Struct2Map(account)
	if err != nil{
		return 0, err
	}

	values, sql := InsertSql(TableUser, insertData)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func updateUser(account Account, multiOrm ...orm.Ormer) error {
	o := NewOrm(multiOrm)
	condition := fmt.Sprintf(`id = %v`, account.Id)
	account.Id = 0
	updateData ,err := Struct2Map(account)
	if err != nil{
		return err
	}
	values, sql := UpdateSql(TableUser, updateData, condition)
	_, err = o.Raw(sql, values).Exec()
	return err
}
