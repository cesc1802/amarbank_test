package handlers

import "github.com/gin-gonic/gin"

func AdminAPI() func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.Group("/admin")
	}
}