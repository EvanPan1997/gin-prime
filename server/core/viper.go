package core

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
func getViper() *viper.Viper {
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
				fmt.Printf("You are using the %s mode of gin framework, the routine of config file is %s\n", gin.Mode(), config)
			} else {
				panic("config filename cannot be empty")
			}
		} else {
			// 正则校验环境变量
			re := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
			match := re.FindString(configEnv)
			if len(match) > 0 {
				fmt.Printf("environment variable is Illegal, value: %s\n", configEnv)
				panic("Illegal Environment Variable")
			} else {
				config = "config." + strings.ToLower(configEnv) + ".yaml"
				fmt.Printf("You are using environment variable, value is %s, routine of config file is %s\n", configEnv, config)
			}
		}
	} else {
		fmt.Printf("You are passing the value using the \"-c\" argument on the command line, routine of config file is %s\n", config)
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
		if err = v.Unmarshal(&global.GP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GP_CONFIG); err != nil {
		panic(err)
	}
	return v
}
