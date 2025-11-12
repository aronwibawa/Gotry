package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var counter int = 0
	router.GET("/ping", func(ctx *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		ctx.String(200, "<h1>halo</h1>")
	})
	router.GET("/count", func(ctx *gin.Context) {

		ctx.String(http.StatusAccepted, "Success%d", counter)
		counter++
	})
	router.POST("/setCount", func(ctx *gin.Context) {
		counter, _ = strconv.Atoi(ctx.Query("counter"))
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
