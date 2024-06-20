package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/global"
	"main/initialize/internal"
)

func GormMysql() *gorm.DB {
	cfg := global.GpConfig.Mysql
	if cfg.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       cfg.Dsn(), // DSN data source name
		DefaultStringSize:         191,       // string 类型字段的默认长度
		SkipInitializeWithVersion: false,     // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+cfg.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		return db
	}
}
