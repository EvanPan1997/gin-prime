package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"main/global"
	"os"
	"regexp"
	"strings"
)

// 优先级: 命令行 > 环境变量 > 默认值
func GetViper() *viper.Viper {
	var config string
	// 获取命令行参数
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
			// 通过gin框架的Mode获取配置文件
			switch gin.Mode() {
			case gin.DebugMode:
				config = global.ConfigDefaultFile
			case gin.ReleaseMode:
				config = global.ConfigReleaseFile
			case gin.TestMode:
				config = global.ConfigTestFile
			}
			if config != "" {
				fmt.Printf("您正在使用gin框架的%s模式, 配置文件路径为%s\n", gin.Mode(), config)
			} else {
				panic("config filename cannot be empty")
			}
		} else {
			// 正则校验环境变量
			re := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
			match := re.FindString(configEnv)
			if len(match) > 0 {
				fmt.Printf("非法环境变量, 值为: %s\n", configEnv)
				panic("Illegal Environment Variable")
			} else {
				config = "config." + strings.ToLower(configEnv) + ".yaml"
				fmt.Printf("您正在使用环境变量读取配置, 值为%s, 配置文件路径为%s\n", configEnv, config)
			}
		}
	} else {
		fmt.Printf("您正在使用命令行\"-c\"指令指定配置文件, 配置文件路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GpConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GpConfig); err != nil {
		panic(err)
	}
	return v
}
