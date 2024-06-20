package core

import (
	"main/global"
	"main/initialize"
)

func Initialize() {
	global.GpViper = initialize.GetViper()
	global.GpDb = initialize.Gorm()
	if global.GpDb != nil {
		db, _ := global.GpDb.DB()
		defer db.Close()

		initialize.RegisterTables()
	}
}
