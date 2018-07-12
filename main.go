package main

import (
	"github.com/gin-gonic/gin"
	"EmailApi/Middlewares"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(Middlewares.CORSMiddleware())//
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9010")
}

