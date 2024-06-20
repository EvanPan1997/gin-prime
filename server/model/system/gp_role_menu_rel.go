package system

import "main/global"

type GpRoleMenuRel struct {
	global.GpModel
	RoleId string `json:"role_id" gorm:"size:32;comment:'角色ID'"`
	MenuId string `json:"menu_id" gorm:"size:32;comment:'菜单ID'"`
}
