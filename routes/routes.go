package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/health_check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.Run(":8080")
}
