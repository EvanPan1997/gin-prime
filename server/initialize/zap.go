package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/global"
	"main/utils"
	"os"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GpConfig.Zap.Directory); !ok {
		fmt.Printf("create %v directory\n", global.GpConfig.Zap.Directory)
		_ = os.Mkdir(global.GpConfig.Zap.Directory, os.ModePerm)
	}
	logger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(zapcore.Lock(os.Stdout)),
			zap.NewAtomicLevelAt(zapcore.DebugLevel)),
		zap.AddCaller(),
		//zap.AddCallerSkip(1),
	)
	return logger
}
