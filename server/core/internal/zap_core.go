package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/global"
	"os"
	"time"
)

type ZapCore struct {
	zapcore.Core
}

func NewZapCore() *ZapCore {
	z := &ZapCore{}
	syncer := z.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return z.Enabled(level)
	})
	z.Core = zapcore.NewCore(global.GpConfig.Zap.Encoder(), syncer, levelEnabler)
	return z
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		global.GpConfig.Zap.Directory,
		CutterWithLayout(time.DateOnly),
	)
	if global.GpConfig.Zap.LogInConsole {
		multiWriteSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiWriteSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *ZapCore) Enabled(l zapcore.Level) bool {
	level, err := zapcore.ParseLevel(global.GpConfig.Zap.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	return l >= level
}

func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	return z.Core.Write(entry, fields)
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
