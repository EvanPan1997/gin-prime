package core

import (
	"main/global"
	"main/initialize"
)

func Initialize() {
	global.GpViper = initialize.GetViper()

	global.GpLogger = initialize.Zap()

	global.GpDb = initialize.Gorm()
	if global.GpDb != nil {
		initialize.RegisterTables()

		db, _ := global.GpDb.DB()
		defer db.Close()
	}

	//global.GpLogger.Info("initialize success")
	//global.GpLogger.Error("initialize error")
}
