package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"main/config"
)

var (
	GP_Viper  *viper.Viper
	GP_CONFIG config.Server
	GP_LOGGER *zap.Logger
)
