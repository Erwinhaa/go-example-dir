package config

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	router := gin.New()
	return router
}
