package system

import "main/global"

type GpMenuWhitelist struct {
	global.GpModel
	MenuId string `json:"menu_id" gorm:"size:32;unique_index:unique_menu_idx"`
}
