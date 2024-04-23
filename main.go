package main

import (
	"github.com/gin-gonic/gin"
	"go_arise/thaiid"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/ping", pingHandler)
	r.POST("/thai/ids/verify", thaiid.ValidateHandler)

	err := r.Run()
	if err != nil {
		return
	}
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
