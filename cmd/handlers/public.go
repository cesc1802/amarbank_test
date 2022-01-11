package handlers

import (
	"amarbank/module/account_v1/transport/gin_account"
	"amarbank/module/transfer_v1/transport/gin_transfer"
	"amarbank/module/user_v1/transport/gin_user"
	"github.com/gin-gonic/gin"
)

func PublicAPI() func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.Use()

		routerV1 := engine.Group("/v1")

		userV1 := routerV1.Group("/users")
		{
			userV1.GET(":id", gin_user.GetUserByID())
			userV1.POST("", gin_user.CreateUser())

		}

		accountV1 := routerV1.Group("/accounts")
		{
			accountV1.POST("", gin_account.CreateAccount())
		}

		transferV1 := routerV1.Group("transfers")
		{
			transferV1.POST("", gin_transfer.CreateTransfer())
		}
	}
}
