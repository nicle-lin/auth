package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gsn/idc_itss/auth"
	"gsn/idc_itss/auth/admin"
)

func main() {
	user := "root"
	passwd := "geesunn123"
	host := "112.74.38.12"
	port := "3306"
	dbName := "gsn_idc_itss"
	auth.InitDatabase(user, passwd, host, port, dbName)
	au := auth.NewAuth()
	//err := InsertData(au)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//err := TestUser(au)
	//if err != nil{
	//	fmt.Println("TestUser:",err)
	//}

	err := TestLoginAndVerify(au)
	fmt.Println(err)


}
func TestLoginAndVerify(au auth.Auth) error{
	loginData := admin.Account{
		Name:"user1",
		Password:"admin",
	}
	userId, err := au.Login(loginData)
	if err != nil{
		fmt.Println(err)
		return err
	}
	fmt.Println("userId:",userId)

	err, ok := au.VerifyAuth(userId,"/monitor/server/new")
	if ok{
		fmt.Println("login success")
	}
	if err != nil{
		fmt.Println(err)
	}
	return nil
}


func TestUser(au auth.Auth) error{
	o := orm.NewOrm()
	o.Begin()
	defer o.Rollback()
	userData1 := admin.Account{
		Name:     "user55",
		Tel: "18825165055",
		Password: "admin",
	}
	userId1, err := au.NewUser(userData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("userId1:", userId1)

	account, err := au.UserInfo("user55",o)
	if err != nil{
		o.Rollback()
		return err
	}
	fmt.Printf("Account1:%+v\n",account)
	updateData := admin.Account{
		Id: userId1,
		Name: "user55_update",
		Password:"test",
	}
	err = au.UpdateUser(updateData,o)
	if err != nil{
		o.Rollback()
		return err
	}

	account, err = au.UserInfo("user55_update",o)
	if err != nil{
		o.Rollback()
		return err
	}
	fmt.Printf("Account2:%+v\n",account)
	o.Commit()

	err = au.DeleteUser(updateData)
	fmt.Println(err)

	return nil
}



func InsertData(au auth.Auth) error {
	o := orm.NewOrm()
	o.Begin()
	defer o.Rollback()
	groupData1 := admin.Group{
		Name: "group111",
	}
	groupId1, err := au.NewGroup(groupData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("GroupId1:", groupId1)

	groupData2 := admin.Group{
		Name: "group222",
	}
	groupId2, err := au.NewGroup(groupData2, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("GroupId2:", groupId2)

	departmentData1 := admin.Department{
		Name:"department111",
	}
	departmentId1, err := au.NewDepartment(departmentData1,o)
	if err != nil {
		o.Rollback()
		return err
	}

	departmentGroupData1 := admin.DepartmentGroup{
		DepartmentId:departmentId1,
		GroupId:groupId1,
	}

	departmentGroupData2 := admin.DepartmentGroup{
		DepartmentId:departmentId1,
		GroupId:groupId2,
	}

	_, err = au.NewDepartmentGroup(departmentGroupData1,o)
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = au.NewDepartmentGroup(departmentGroupData2,o)
	if err != nil {
		o.Rollback()
		return err
	}

	userData1 := admin.Account{
		Name:     "user111",
		Tel: "18825165111",
		Password: "admin",
	}
	userId1, err := au.NewUser(userData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("userId1:", userId1)

	userData2 := admin.Account{
		Name:     "user222",
		Tel: "18825165222",
		Password: "admin",
	}
	userId2, err := au.NewUser(userData2, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("userId2:", userId2)

	userData4 := admin.Account{
		Name:     "user444",
		Tel: "18825165444",
		Password: "admin",
	}
	userId4, err := au.NewUser(userData4, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("userId4:", userId4)

	userData3 := admin.Account{
		Name:     "user333",
		Tel: "18825165333",
		Password: "admin",
	}
	userId3, err := au.NewUser(userData3, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("userId3:", userId3)

	groupUserData1 := admin.GroupUser{
		UserId:  userId1,
		GroupId: groupId1,
	}
	groupUserId1, err := au.NewGroupUser(groupUserData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId1:", groupUserId1)

	groupUserData2 := admin.GroupUser{
		UserId:  userId2,
		GroupId: groupId1,
	}
	groupUserId2, err := au.NewGroupUser(groupUserData2, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId2:", groupUserId2)

	groupUserData3 := admin.GroupUser{
		UserId:  userId3,
		GroupId: groupId1,
	}
	groupUserId3, err := au.NewGroupUser(groupUserData3, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId3:", groupUserId3)

	groupUserData4 := admin.GroupUser{
		UserId:  userId1,
		GroupId: groupId2,
	}
	groupUserId4, err := au.NewGroupUser(groupUserData4, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId4:", groupUserId4)

	groupUserData5 := admin.GroupUser{
		UserId:  userId2,
		GroupId: groupId2,
	}
	groupUserId5, err := au.NewGroupUser(groupUserData5, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId5:", groupUserId5)

	groupUserData6 := admin.GroupUser{
		UserId:  userId3,
		GroupId: groupId2,
	}
	groupUserId6, err := au.NewGroupUser(groupUserData6, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId6:", groupUserId6)

	groupUserData7 := admin.GroupUser{
		UserId:  userId4,
		GroupId: groupId2,
	}
	groupUserId7, err := au.NewGroupUser(groupUserData7, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("groupUserId7:", groupUserId7)

	roleData1 := admin.Role{
		Name: "role111",
	}
	roleData2 := admin.Role{
		Name: "role222",
	}

	roleId1, err := au.NewRole(roleData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("roleId1:", roleId1)

	roleId2, err := au.NewRole(roleData2, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("roleId2:", roleId2)

	userRoleData1 := admin.UserRole{
		RoleId: roleId1,
		UserId: userId1,
	}

	userRoleData2 := admin.UserRole{
		RoleId: roleId1,
		UserId: userId2,
	}

	userRoleData3 := admin.UserRole{
		RoleId: roleId1,
		UserId: userId3,
	}
	userRoleData4 := admin.UserRole{
		RoleId: roleId2,
		UserId: userId1,
	}

	userRoleData5 := admin.UserRole{
		RoleId: roleId2,
		UserId: userId2,
	}

	userRoleData6 := admin.UserRole{
		RoleId: roleId2,
		UserId: userId3,
	}

	userRoleData7 := admin.UserRole{
		RoleId: roleId2,
		UserId: userId4,
	}

	userRole := map[string]admin.UserRole{
		"userRoleData1": userRoleData1,
		"userRoleData2": userRoleData2,
		"userRoleData3": userRoleData3,
		"userRoleData4": userRoleData4,
		"userRoleData5": userRoleData5,
		"userRoleData6": userRoleData6,
		"userRoleData7": userRoleData7,
	}
	for _, v := range userRole {
		_, err := au.NewUserRole(v, o)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	objectData1 := admin.Object{
		Name: "服务器11",
	}
	objectData2 := admin.Object{
		Name: "网络设备11",
	}

	objectId1, err := au.NewObject(objectData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("objectId1:", objectId1)
	objectId2, err := au.NewObject(objectData2, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("objectId2:", objectId2)

	scopeData1 := admin.Scope{
		IdcId:  1,
		UserId: userId1,
		Name:   "scope1_for_test11",
	}
	scopeId1, err := au.NewScope(scopeData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("scopeId1:", scopeId1)

	permissionData1 := admin.Permission{
		ScopeId:  scopeId1,
		ObjectId: objectId1,
	}
	permissionData2 := admin.Permission{
		ScopeId:  scopeId1,
		ObjectId: objectId2,
	}
	permissionId1, err := au.NewPermission(permissionData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	permissionId2, err := au.NewPermission(permissionData2, o)
	if err != nil {
		o.Rollback()
		return err
	}

	rolePermissionData1 := admin.RolePermission{
		PermissionId: permissionId1,
		RoleId:       roleId1,
	}
	rolePermissionData2 := admin.RolePermission{
		PermissionId: permissionId2,
		RoleId:       roleId2,
	}

	_, err = au.NewRolePermission(rolePermissionData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = au.NewRolePermission(rolePermissionData2, o)
	if err != nil {
		o.Rollback()
		return err
	}

	actionData1 := admin.Action{
		Name: "新增11",
	}
	actionData2 := admin.Action{
		Name: "查看11",
	}
	actionId1, err := au.NewAction(actionData1, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("actionId1:", actionId1)

	actionId2, err := au.NewAction(actionData2, o)
	if err != nil {
		o.Rollback()
		return err
	}
	fmt.Println("actionId2:", actionId2)

	objectActionData := map[string]admin.ObjectAction{
		"objectActionData1": {
			ObjectId: objectId1,
			ActionId: actionId1,
		},
		"objectActionData2": {
			ObjectId: objectId1,
			ActionId: actionId2,
		},
		"objectActionData3": {
			ObjectId: objectId2,
			ActionId: actionId1,
		},
		"objectActionData4": {
			ObjectId: objectId2,
			ActionId: actionId2,
		},
	}

	var (
		objectActionId1 int64
		objectActionId2 int64
		objectActionId3 int64
		objectActionId4 int64
	)
	objectActionId := map[string]*int64{
		"objectActionData1": &objectActionId1,
		"objectActionData2": &objectActionId2,
		"objectActionData3": &objectActionId3,
		"objectActionData4": &objectActionId4,
	}

	for k, v := range objectActionData {
		*objectActionId[k], err = au.NewObjectAction(v, o)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	urlData := map[string]admin.Node{
		"urlData1": {
			Url:  "/monitor/server/new11",
			Name: "创建服务器11",
		},
		"urlData2": {
			Url:  "/monitor/server/search11",
			Name: "查询服务器列表11",
		},
		"urlData3": {
			Url:  "/monitor/network/new11",
			Name: "创建网络设备11",
		},
		"urlData4": {
			Url:  "/monitor/network/search11",
			Name: "查询网络设备列表11",
		},
	}
	var (
		urlId1 int64
		urlId2 int64
		urlId3 int64
		urlId4 int64
	)
	urlId := map[string]*int64{
		"urlData1": &urlId1,
		"urlData2": &urlId2,
		"urlData3": &urlId3,
		"urlData4": &urlId4,
	}

	for k, v := range urlData {
		*urlId[k], err = au.NewNode(v, o)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	objectActionUrlData := map[string]admin.ObjectActionUrl{
		"objectActionUrlData1": {
			ObjectActionId: objectActionId1,
			UrlId:          urlId1,
		},
		"objectActionUrlData2": {
			ObjectActionId: objectActionId2,
			UrlId:          urlId2,
		},
		"objectActionUrlData3": {
			ObjectActionId: objectActionId3,
			UrlId:          urlId3,
		},
		"objectActionUrlData4": {
			ObjectActionId: objectActionId4,
			UrlId:          urlId4,
		},
	}
	for _, v := range objectActionUrlData {
		_, err := au.NewObjectActionUrl(v, o)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	return o.Commit()
}
