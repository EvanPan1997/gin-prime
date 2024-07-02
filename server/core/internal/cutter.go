package internal

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Cutter 实现 io.Writer 接口
// 用于日志切割, strings.Join([]string{director,layout, formats..., level+".log"}, os.PathSeparator)
type Cutter struct {
	level     string // 日志级别(debug, info, warn, error, dpanic, panic, fatal)
	layout    string // 时间格式 2006-01-02 15:04:05
	suffix    string // 后缀 .log
	directory string // 日志文件夹
	file      *os.File
	mutex     *sync.RWMutex // 读写锁
}

type CutterOption func(*Cutter)

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

func NewCutter(level string, directory string, options ...CutterOption) *Cutter {
	c := &Cutter{
		level:     level,
		directory: directory,
		mutex:     new(sync.RWMutex),
	}
	for i := 0; i < len(options); i++ {
		options[i](c)
	}
	return c
}

func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()

	// 写入逻辑
	values := make([]string, 0, 3)
	values = append(values, c.directory)
	if c.layout != "" {
		values = append(values, time.Now().Format(c.layout))
	}
	values = append(values, c.level+".log")
	filename := filepath.Join(values...)
	director := filepath.Dir(filename)
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}

func (c *Cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.file != nil {
		return c.file.Sync()
	}
	return nil
}
