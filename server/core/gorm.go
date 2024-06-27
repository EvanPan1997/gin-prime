package core

import (
	"fmt"
	"gorm.io/gorm"
	"main/global"
	"main/model/system"
	"os"
)

func Gorm() *gorm.DB {
	switch global.GpConfig.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GpDb
	err := db.AutoMigrate(
		system.GpMenu{},
		system.GpMenuWhitelist{},
		system.GpRole{},
		system.GpRoleMenuRel{},
	)
	if err != nil {
		fmt.Println("register table failed, ", err)
		os.Exit(0)
	}
	fmt.Println("register table success")
}
