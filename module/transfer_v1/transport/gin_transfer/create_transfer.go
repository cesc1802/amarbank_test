package gin_transfer

import (
	"amarbank/mockdata"
	storage3 "amarbank/module/account_v1/storage"
	"amarbank/module/transfer_v1/dto"
	service2 "amarbank/module/transfer_v1/service"
	"amarbank/module/transfer_v1/storage"
	storage2 "amarbank/module/user_v1/storage"
	"amarbank/pkg/rest_response"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CreateTransfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreateTransferRequest

		if err := c.ShouldBind(&input); err != nil {
			panic(err)
		}

		store, _ := storage.NewTxnFileStorage(mockdata.DefaultPathToTxn, os.ModeAppend)
		userStore, _ := storage2.NewUSerFileStorage(mockdata.DefaultPathToUser, os.ModeAppend)
		accStore, _ := storage3.NewAccountFileStorage(mockdata.DefaultPathToAccount, os.ModeAppend)

		service := service2.NewCreateTxnService(store, accStore, userStore)

		if err := service.CreateTxn(c.Request.Context(), &input); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, rest_response.SimpleSuccessResponse(true))

	}
}
