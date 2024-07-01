package core

import (
	"fmt"
	"go.uber.org/zap"
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
	//cores := make([]zapcore.Core, 0, 7)
	//for i := 0; i <= 7; i++ {
	//	cores = append(cores, internal.NewZapCore())
	//}
	core := internal.NewZapCore()
	logger = zap.New(core)
	return logger
}
