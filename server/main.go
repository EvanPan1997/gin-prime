package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func init() {
	fmt.Println(
		"  ____ _             ____       _                \n" +
			" / ___(_)_ __       |  _ \\ _ __(_)_ __ ___   ___ \n" +
			"| |  _| | '_ \\ _____| |_) | '__| | '_ ` _ \\ / _ \\\n" +
			"| |_| | | | | |_____|  __/| |  | | | | | | |  __/\n" +
			" \\____|_|_| |_|     |_|   |_|  |_|_| |_| |_|\\___|")
}

func main() {
	r := gin.Default()
	server := endless.NewServer(":8080", r)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
