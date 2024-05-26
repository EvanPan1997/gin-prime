package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

func RunServer() {
	r := gin.Default()
	server := endless.NewServer(":8080", r)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
