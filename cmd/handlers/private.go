package handlers

import "github.com/gin-gonic/gin"

func PrivateAPI() func(engine *gin.Engine) {
	return func(engine *gin.Engine) {

		//TODO: setup some until middleware like Authentication and Authorization
		engine.Use()
	}
}