package internal

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Cutter 实现 io.Writer 接口
// 用于日志切割, strings.Join([]string{director,layout, formats..., level+".log"}, os.PathSeparator)
type Cutter struct {
	layout    string // 时间格式 2006-01-02 15:04:05
	suffix    string // 后缀 .log
	directory string // 日志文件夹
	file      *os.File
	mutex     *sync.RWMutex // 读写锁
}

const limitSize int64 = 4 * 1024 * 1024 * 1024

type CutterOption func(*Cutter)

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

func NewCutter(directory string, options ...CutterOption) *Cutter {
	c := &Cutter{directory: directory, mutex: new(sync.RWMutex)}
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
	currentDate := strings.ReplaceAll(time.Now().Format(c.layout), "-", "")
	for cnt := 0; ; cnt++ {
		var filename string
		if cnt == 0 {
			filename = currentDate + c.suffix
		} else {
			filename = currentDate + "-" + strconv.Itoa(cnt) + c.suffix
		}

		filePath := c.directory + "/" + filename

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				c.file, _ = os.Create(filePath)
				return c.file.Write(bytes)
			} else {
				return 0, err
			}
		}

		c.file, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return 0, err
		}

		if fileInfo.Size() < limitSize {
			return c.file.Write(bytes)
		}
	}
}

func (c *Cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.file != nil {
		return c.file.Sync()
	}
	return nil
}
