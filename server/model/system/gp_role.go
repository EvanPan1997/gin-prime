package system

import (
	"database/sql"
	"main/global"
)

type GpRole struct {
	global.GpModel
	RoleId    string       `json:"role_id" gorm:"size:32;comment:'角色ID'"`
	RoleName  string       `json:"role_name" gorm:"size:64;comment:'角色名称'"`
	ParentId  string       `json:"parent_id" gorm:"size:32;comment:'上一级角色ID'"`
	IsEnabled sql.NullBool `json:"is_enabled" gorm:"comment:'是否启用'"`
	Remarks   string       `json:"remarks" gorm:"size:'256';comment:'备注'"`
}
