package main

import (
	"golang-gin/model"
	"golang-gin/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	u := user.NewUser()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Golang Gin",
		})
	})

	r.GET("/new/user", func(c *gin.Context) {
		c.JSON(200, u.AllData())
	})

	r.POST("/add/user", func(ctx *gin.Context) {
		var m model.Interface

		err := ctx.ShouldBindJSON(&m)

		if err != nil {
			ctx.JSON(400, gin.H{
				"Message": "Error",
			})
			return
		}

		u.AddData(m)
		ctx.JSON(200, m)

	})

	r.Run(":8000")
}
