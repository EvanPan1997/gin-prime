package core

import (
	"main/global"
)

func Initialize() {
	global.GpViper = GetViper()

	global.GpLogger = Zap()

	//global.GpDb = Gorm()
	//if global.GpDb != nil {
	//	RegisterTables()
	//
	//	db, _ := global.GpDb.DB()
	//	defer db.Close()
	//}
}
