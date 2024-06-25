package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"main/config"
)

var (
	GpViper  *viper.Viper
	GpConfig config.Server
	GpLogger *zap.Logger
	GpDb     *gorm.DB
	//GP_LOGGER *zap.Logger
)
