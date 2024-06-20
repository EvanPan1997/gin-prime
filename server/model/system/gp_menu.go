package system

import (
	"database/sql"
	"main/global"
)

type GpMenu struct {
	global.GpModel
	MenuId      string       `json:"menu_id" gorm:"size:32;comment:'菜单ID'"`
	MenuName    string       `json:"menu_name" gorm:"size:128;comment:'菜单名称'"`
	ParentId    string       `json:"parent_id" gorm:"size:32;comment:'上一个节点ID'"`
	ComponentId string       `json:"component_id" gorm:"size:128;comment:'组件ID'"`
	Path        string       `json:"path" gorm:"size:256;comment:'路径'"`
	Sequence    int          `json:"sequence" gorm:"comment:'顺序'"`
	IsEnabled   sql.NullBool `json:"is_enabled" gorm:"comment:'是否启用'"`
	Remarks     string       `json:"remarks" gorm:"size:500;comment:'备注'"`
}
