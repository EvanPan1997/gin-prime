package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"main/middleware"
)

func RunServer() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/", Func1(), Func2())

	server := endless.NewServer(":8080", r)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func Func1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("func1 before func2")
		c.Next()
		fmt.Println("func1 after func2")
	}
}

func Func2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("func2")
	}
}
