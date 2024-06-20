package internal

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"main/config"
	"main/global"
	"os"
	"time"
)

type _gorm struct{}

var Gorm = new(_gorm)

func (g *_gorm) Config() *gorm.Config {
	var general config.GeneralDB
	switch global.GpConfig.System.DbType {
	case "mysql":
		general = global.GpConfig.Mysql.GeneralDB
	default:
		general = global.GpConfig.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
