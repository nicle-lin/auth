package auth

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gsn/idc_itss/auth/admin"
	memCache "gsn/idc_itss/auth/cache"
	"io"
	"os"
	"strings"
	"time"
)

/*
权限模块
*/

type Auth interface {
	VerifyAuth(userId int64, url string) (error, bool)
	Login(account admin.Account) (int64, error)
	Logout(userId int64) error
	SetNodeToCache() (map[string]admin.Node, error)
	SetUserAuthToCache(userId int64) (map[int64]admin.ObjectActionUrl, error)
	SetScopeToCache(userId int64)
	Cache() memCache.Cache
	admin.Admin
}

var defaultCache memCache.Cache

func init() {
	var err error
	defaultCache, err = memCache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		panic(err)
	}
}

type DefaultAuth struct {
	cache memCache.Cache
	admin.DefaultAdmin
}

func InitDatabase(user, password, host, port, dbName string) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", user, password, host, port, dbName))
}

func NewAuth() Auth {
	return &DefaultAuth{
		cache: defaultCache,
	}
}

// login successfully, set user/role/permission to cache
func (d *DefaultAuth) Login(account admin.Account) (int64, error) {
	user, err := d.UserInfo(account.Name)
	if err != nil {
		return 0, err
	}
	if user.Name == account.Name && admin.Md5sum(account.Password) == user.Password {

		_, err := d.SetUserAuthToCache(user.Id)
		if err != nil {
			return user.Id, err
		}
		return user.Id, nil
	}

	return 0, admin.ErrUsernameOrPasswordIncorrect
}

func (d *DefaultAuth) Logout(userId int64) error {
	keyUser := fmt.Sprintf("%v%v", admin.KeyUser, userId)
	return d.cache.Delete(keyUser)
}

func (d *DefaultAuth) VerifyAuth(userId int64, url string) (error, bool) {

	//step 1, check url
	nodes, err := d.SetNodeToCache()
	if err != nil {
		return err, false
	}
	v, ok := nodes[url]
	if !ok {
		return admin.ErrSystemMissingNode, false
	}

	//step 2, check permission
	userAuth, err := d.SetUserAuthToCache(userId)
	if err != nil {
		return err, false
	}

	// verifyAuth
	_, ok = userAuth[v.Id]
	if !ok {
		return admin.ErrPermissionDeny, false
	}
	return nil, true
}

func (d *DefaultAuth) SetScopeToCache(userId int64) {

}

func (d *DefaultAuth) SetNodeToCache() (map[string]admin.Node, error) {
	node := d.cache.Get(admin.KeyNode)
	if node == nil { //doesn't exist
		nodes, err := d.AllNodeInfo()
		if err != nil {
			return nil, err
		}
		nodeTmp := make(map[string]admin.Node)
		for _, v := range nodes {
			nodeTmp[v.Url] = v
		}
		err = d.cache.Put(admin.KeyNode, nodeTmp, 24*3600*time.Second)
		return nodeTmp, err
	} else {

		v, ok := node.(map[string]admin.Node)
		if !ok {
			return nil, admin.ErrSystemAssert
		}
		return v, nil
	}
}

func (d *DefaultAuth) SetUserAuthToCache(userId int64) (map[int64]admin.ObjectActionUrl, error) {
	key := fmt.Sprintf("%v%v", admin.KeyUser, userId)
	userAuth := d.cache.Get(key)
	if userAuth == nil { // doesn't exist
		// get it from database
		//get role
		userRoles, err := d.UserRoleInfoByUserId(userId)
		if err != nil {
			return nil, err
		}
		if len(userRoles) == 0 {
			return nil, admin.ErrUserMissingRole
		}

		var rolesId []int64
		for _, v := range userRoles {
			rolesId = append(rolesId, v.RoleId)
		}

		//get rolePermission
		rolePermissions, err := d.MultiRolePermission(rolesId)
		if err != nil {
			return nil, err
		}

		if len(rolePermissions) == 0 {
			return nil, admin.ErrRoleMissingPermission
		}

		var permissionsId []int64
		for _, v := range rolePermissions {
			permissionsId = append(permissionsId, v.PermissionId)
		}

		//get permission
		permissions, err := d.MultiPermission(permissionsId)
		if err != nil {
			return nil, err
		}
		if len(permissions) == 0 {
			return nil, admin.ErrRoleMissingPermission
		}
		var objectsId []int64
		for _, v := range permissions {
			objectsId = append(objectsId, v.ObjectId)
		}
		// get ObjectAction
		objectActions, err := d.MultiObjectAction(objectsId)
		if err != nil {
			return nil, err
		}
		if len(objectActions) == 0 {
			return nil, admin.ErrObjectMissingAction
		}

		var objectActionsId []int64
		for _, v := range objectActions {
			objectActionsId = append(objectActionsId, v.Id)
		}
		// get ObjectActionUrl
		objectActionUrls, err := d.MultiObjectActionUrl(objectActionsId)
		if err != nil {
			return nil, err
		}
		if len(objectActionUrls) == 0 {
			return nil, admin.ErrObjectActionMissingUrl
		}

		authInfo := make(map[int64]admin.ObjectActionUrl)
		for _, v := range objectActionUrls {
			authInfo[v.UrlId] = v
		}
		err = d.cache.Put(key, authInfo, 24*3600*time.Second)
		return authInfo, err
	} else {

		v, ok := userAuth.(map[int64]admin.ObjectActionUrl)
		if !ok {
			return nil, admin.ErrSystemAssert
		}
		return v, nil
	}
}

func (d *DefaultAuth) Cache() memCache.Cache {
	return d.cache
}

var RouterPath = "gsn/idc_itss/conf/routes"

//注册权限节点
func RegisterNode(url, name string) error {

	return nil
}

func readLinesFromFile(filename string) ([]string, error) {
	if filename == "" {
		filename = RouterPath
	}

	var r io.Reader = os.Stdin
	if filename != "-" {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}

	var lines []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// ignore empty lines
		if line == "" {
			continue
		}
		// strip comments
		if strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
