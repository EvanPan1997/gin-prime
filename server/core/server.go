package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"main/global"
	"main/middleware"
	"strconv"
)

func RunServer() {
	for i := 0; i < 100000; i++ {
		global.GpLogger.With(zap.String("key", strconv.Itoa(i))).Info("info test")
	}

	r := gin.Default()
	r.Use(middleware.Cors())

	server := endless.NewServer(":8080", r)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
