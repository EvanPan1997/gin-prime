package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/core/internal"
	"main/global"
	"main/utils"
	"os"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GpConfig.Zap.Directory); !ok {
		fmt.Printf("create %v directory\n", global.GpConfig.Zap.Directory)
		_ = os.Mkdir(global.GpConfig.Zap.Directory, os.ModePerm)
	}
	//logger = zap.New(
	//	zapcore.NewCore(
	//		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
	//		zapcore.AddSync(zapcore.Lock(os.Stdout)),
	//		zap.NewAtomicLevelAt(zapcore.DebugLevel)),
	//	zap.AddCaller(),
	//)

	levels := global.GpConfig.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}

	logger = zap.New(zapcore.NewTee(cores...))
	if global.GpConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
