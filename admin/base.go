package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type ItemType uint8

const (
	TypeGroup ItemType = iota
	TypeGroupUser
	TypeUser
	TypeUserRole
	TypeRole
	TypeRolePermission
	TypePermission
	TypeObject
	TypeObjectAction
	TypeAction
	TypeObjectActionUrl
	TypeUrl
	TypeScope
	TypeDepartment
	TypeDepartmentGroup
)

const (
	KeyUser = "keyUser-"
	KeyNode = "AllKeyNode"
	KeyScope = "keyScope-"
)

var (
	ErrSystemPanic       = errors.New("系统异常")
	ErrSystemAssert      = errors.New("内部断言错误")
	ErrSystemMissingNode = errors.New("此路由未录入权限管理,请联系管理员")
	ErrPermissionDeny    = errors.New("您无权限操作")

	ErrRoleAlreadyExist       = errors.New("角色已经存在")
	ErrRoleNameEmpty          = errors.New("角色名为空")
	ErrInvalidRoleId          = errors.New("无效的角色Id")
	ErrMissingRoleId          = errors.New("缺少角色Id")
	ErrUserMissingRole        = errors.New("帐号没有绑定角色，请联系管理员")
	ErrRoleMissingPermission  = errors.New("角色没有绑定权限，请联系管理员")
	ErrObjectMissingAction    = errors.New("操作对象没有关联操作动作,请联系管理员")
	ErrObjectActionMissingUrl = errors.New("操作对象动作没有关联url,请联系管理员")

	ErrPermissionAlreadyExist = errors.New("许可已经存在")
	ErrInvalidPermissionId    = errors.New("无效的许可Id")

	ErrDepartmentAlreadyExist = errors.New("部门已经存在")
	ErrDepartmentNameEmpty    = errors.New("部门名为空")
	ErrInvalidDepartmentId    = errors.New("无效的部门Id")

	ErrDepartmentGroupAlreadyExist = errors.New("部门用户组已经存在")
	ErrInvalidDepartmentGroupId    = errors.New("无效的部门用户组Id")

	ErrActionAlreadyExist = errors.New("操作动作已经存在")
	ErrActionNameEmpty    = errors.New("操作动作名为空")
	ErrInvalidActonId     = errors.New("无效的操作动作Id")

	ErrGroupAlreadyExist = errors.New("用户组已经存在")
	ErrGroupNameEmpty    = errors.New("用户组名为空")
	ErrInvalidGroupId    = errors.New("无效的用户组Id")

	ErrGroupUserAlreadyExist = errors.New("用户组用户已经存在")
	ErrInvalidGroupUserId    = errors.New("无效的用户组Id")

	ErrObjectAlreadyExist = errors.New("操作对象已经存在")
	ErrObjectNameEmpty    = errors.New("操作对象名为空")
	ErrInvalidObjectId    = errors.New("无效的操作对象Id")

	ErrInvalidObjectActionId    = errors.New("无效的操作对象动作Id")
	ErrObjectActionAlreadyExist = errors.New("此操作对象已经绑定了此动作")

	ErrInvalidObjectActionUrlId    = errors.New("无效的操作对象动作Url的Id")
	ErrObjectActionUrlAlreadyExist = errors.New("此操作对象动作已经绑定了此Url")

	ErrInvalidRolePermissionId    = errors.New("无效的角色许可Id")
	ErrRolePermissionAlreadyExist = errors.New("此角色已经绑定了此许可")

	ErrScopeAlreadyExist = errors.New("操作范围已经存在")
	ErrScopeNameEmpty    = errors.New("操作范围名称为空")
	ErrScopeIdcIdEmpty   = errors.New("操作范围库区为空")
	ErrScopeUserIdEmpty  = errors.New("操作范围负责人为空")
	ErrInvalidScopeId    = errors.New("无效的操作范围Id")

	ErrNodeAlreadyExist = errors.New("权限节点已经存在")
	ErrNodeNameEmpty    = errors.New("权限节点名称为空")
	ErrNodeURLEmpty     = errors.New("权限节点URL为空")
	ErrInvalidNodeId    = errors.New("无效的权限节点Id")

	ErrUserAlreadyExist            = errors.New("用户已经存在")
	ErrPasswordEmpty               = errors.New("密码为空")
	ErrUsernameEmpty               = errors.New("用户名为空")
	ErrTelephoneEmpty              = errors.New("手机号为空")
	ErrInvalidUserId               = errors.New("无效的Id")
	ErrUsernameOrPasswordIncorrect = errors.New("用户名或密码错误")

	ErrInvalidUserRoleId    = errors.New("无效的用户角色Id")
	ErrUserRoleAlreadyExist = errors.New("此用户已经绑定了此角色")
)

type DefaultAdmin struct{}

type Admin interface {
	NewDepartment(department Department, multiOrm ...orm.Ormer) (int64, error)
	UpdateDepartment(department Department, multiOrm ...orm.Ormer) error
	DeleteDepartment(department Department, multiOrm ...orm.Ormer) error
	DepartmentInfo(id int64, multiOrm ...orm.Ormer) (Department, error)

	NewDepartmentGroup(gu DepartmentGroup, multiOrm ...orm.Ormer) (int64, error)
	UpdateDepartmentGroup(gu DepartmentGroup, multiOrm ...orm.Ormer) error
	DeleteDepartmentGroup(gu DepartmentGroup, multiOrm ...orm.Ormer) error
	DepartmentGroupInfo(departmentId int64, multiOrm ...orm.Ormer) ([]DepartmentGroup, error)

	NewGroup(group Group, multiOrm ...orm.Ormer) (int64, error)
	UpdateGroup(group Group, multiOrm ...orm.Ormer) error
	DeleteGroup(group Group, multiOrm ...orm.Ormer) error
	GroupInfo(id int64, multiOrm ...orm.Ormer) (Group, error)

	NewGroupUser(gu GroupUser, multiOrm ...orm.Ormer) (int64, error)
	UpdateGroupUser(gu GroupUser, multiOrm ...orm.Ormer) error
	DeleteGroupUser(gu GroupUser, multiOrm ...orm.Ormer) error
	GroupUserInfo(id int64, multiOrm ...orm.Ormer) ([]GroupUser, error)

	NewUser(account Account, multiOrm ...orm.Ormer) (int64, error)
	UpdateUser(account Account, multiOrm ...orm.Ormer) error
	DeleteUser(account Account, multiOrm ...orm.Ormer) error
	UserInfo(name string, multiOrm ...orm.Ormer) (Account, error)

	NewUserRole(ur UserRole, multiOrm ...orm.Ormer) (int64, error)
	UpdateUserRole(ur UserRole, multiOrm ...orm.Ormer) error
	DeleteUserRole(ur UserRole, multiOrm ...orm.Ormer) error
	UserRoleInfoByUserId(userId int64, multiOrm ...orm.Ormer) ([]UserRole, error)
	UserRoleInfoByRoleId(roleId int64, multiOrm ...orm.Ormer) ([]UserRole, error)

	NewRole(role Role, multiOrm ...orm.Ormer) (int64, error)
	UpdateRole(role Role, multiOrm ...orm.Ormer) error
	DeleteRole(role Role, multiOrm ...orm.Ormer) error
	RoleInfo(id int64, multiOrm ...orm.Ormer) (Role, error)

	NewRolePermission(rp RolePermission, multiOrm ...orm.Ormer) (int64, error)
	UpdateRolePermission(rp RolePermission, multiOrm ...orm.Ormer) error
	DeleteRolePermission(rp RolePermission, multiOrm ...orm.Ormer) error
	RolePermissionInfo(roleId int64, multiOrm ...orm.Ormer) ([]RolePermission, error)

	NewPermission(permission Permission, multiOrm ...orm.Ormer) (int64, error)
	UpdatePermission(permission Permission, multiOrm ...orm.Ormer) error
	DeletePermission(permission Permission, multiOrm ...orm.Ormer) error
	PermissionInfo(id int64, multiOrm ...orm.Ormer) ([]Permission, error)
	MultiPermission(rolesId []int64, multiOrm ...orm.Ormer) ([]Permission, error)

	NewObject(object Object, multiOrm ...orm.Ormer) (int64, error)
	UpdateObject(object Object, multiOrm ...orm.Ormer) error
	DeleteObject(object Object, multiOrm ...orm.Ormer) error
	ObjectInfo(id int64, multiOrm ...orm.Ormer) (Object, error)

	NewObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) (int64, error)
	UpdateObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) error
	DeleteObjectAction(ur ObjectAction, multiOrm ...orm.Ormer) error
	ObjectActionInfo(objectId int64, multiOrm ...orm.Ormer) ([]ObjectAction, error)

	NewAction(action Action, multiOrm ...orm.Ormer) (int64, error)
	UpdateAction(action Action, multiOrm ...orm.Ormer) error
	DeleteAction(action Action, multiOrm ...orm.Ormer) error
	ActionInfo(id int64, multiOrm ...orm.Ormer) (Action, error)

	NewObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) (int64, error)
	UpdateObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) error
	DeleteObjectActionUrl(ur ObjectActionUrl, multiOrm ...orm.Ormer) error
	ObjectActionUrlInfo(objectActionId int64, multiOrm ...orm.Ormer) ([]ObjectActionUrl, error)
	MultiObjectActionUrl(objectActionsId []int64, multiOrm ...orm.Ormer) ([]ObjectActionUrl, error)

	NewNode(node Node, multiOrm ...orm.Ormer) (int64, error)
	UpdateNode(node Node, multiOrm ...orm.Ormer) error
	DeleteNode(node Node, multiOrm ...orm.Ormer) error
	NodeInfo(id int64, multiOrm ...orm.Ormer) (Node, error)
	NodeInfoByUrl(url string, multiOrm ...orm.Ormer) (Node, error)
	AllNodeInfo(multiOrm ...orm.Ormer) ([]Node, error)

	NewScope(scope Scope, multiOrm ...orm.Ormer) (int64, error)
	UpdateScope(scope Scope, multiOrm ...orm.Ormer) error
	DeleteScope(scope Scope, multiOrm ...orm.Ormer) error
	ScopeInfo(id int64, multiOrm ...orm.Ormer) (Scope, error)
}

func NewAdmin() Admin {
	return &DefaultAdmin{}
}
