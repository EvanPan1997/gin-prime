package core

import (
	"main/global"
)

func Initialize() {
	global.GpViper = GetViper()

	global.GpLogger = Zap() // 初始化日志库

	global.GpDb = Gorm()
	if global.GpDb != nil {
		RegisterTables()

		db, _ := global.GpDb.DB()
		defer db.Close()
	}
}
