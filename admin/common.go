package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"encoding/json"
)

func Md5sum(plaintext string) (ciphertext string) {
	h := md5.New()
	h.Write([]byte(plaintext))
	ciphertext = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func CheckExist(o orm.Ormer, table string, data map[string]interface{}) (int64, bool) {
	var (
		columns []string
		args    []interface{}
		id      int64
	)
	for k, v := range data {
		columns = append(columns, k+"=? ")
		args = append(args, v)
	}
	sql := fmt.Sprintf(`
		SELECT id FROM %v WHERE %v LIMIT 1
	`, table, strings.Join(columns, " AND "))

	err := o.Raw(sql, args...).QueryRow(&id)
	if err != nil && err != orm.ErrNoRows {
		panic(err)
	}
	if id != 0 {
		return id, true
	}
	return 0, false
}

func NewOrm(multiOrm []orm.Ormer) (o orm.Ormer) {
	if len(multiOrm) == 0 {
		o = orm.NewOrm()
	} else if len(multiOrm) == 1 {
		o = multiOrm[0]
	} else {
		panic("只能传一个Ormer")
	}
	return
}

func Join(a []int64,sep string) string{
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), sep), "[]")
}

func Struct2Map(v interface{}) (map[string]interface{},error){
	str, err := json.Marshal(v)
	if err != nil{
		return nil,err
	}
	mapInfo := make(map[string]interface{})
	err = json.Unmarshal([]byte(str),&mapInfo)
	if err != nil{
		return nil,err
	}
	return mapInfo,nil
}