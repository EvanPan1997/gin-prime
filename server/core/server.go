package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"main/global"
	"main/middleware"
)

func RunServer() {
	global.GpLogger.Info("info test")

	r := gin.Default()
	r.Use(middleware.Cors())

	server := endless.NewServer(":8080", r)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
