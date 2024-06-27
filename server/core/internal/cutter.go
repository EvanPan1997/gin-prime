package internal

import (
	"sync"
)

// Cutter 实现 io.Writer 接口
// 用于日志切割, strings.Join([]string{director,layout, formats..., level+".log"}, os.PathSeparator)
type Cutter struct {
	layout         string // 时间格式 2006-01-02 15:04:05
	fileNameFormat string // "2006-01-02"
	suffix         string // 后缀 .log
	directory      string // 日志文件夹
	//file       *os.File      // 文件句柄
	mutex *sync.RWMutex // 读写锁
}

type CutterOption func(*Cutter)

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

//// CutterWithFormats 格式化参数
//func CutterWithFormats(format ...string) CutterOption {
//	return func(c *Cutter) {
//		if len(format) > 0 {
//			c.formats = format
//		}
//	}
//}
