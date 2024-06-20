package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"main/config"
)

var (
	GpViper  *viper.Viper
	GpConfig config.Server
	GpDb     *gorm.DB
	//GP_LOGGER *zap.Logger
)
