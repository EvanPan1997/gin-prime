package main

import (
	"fmt"
	"main/core"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func init() {
	fmt.Println(
		"          _                        _              \n" +
			"   ____ _(_)___        ____  _____(_)___ ___  ___ \n" +
			"  / __ `/ / __ \\______/ __ \\/ ___/ / __ `__ \\/ _ \\\n" +
			" / /_/ / / / / /_____/ /_/ / /  / / / / / / /  __/\n" +
			" \\__, /_/_/ /_/     / .___/_/  /_/_/ /_/ /_/\\___/ \n" +
			"/____/             /_/                            ")
}

func main() {
	core.RunServer()
}
