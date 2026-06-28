package routes

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine, db interface{}) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
