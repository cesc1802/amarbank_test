package gin_account

import (
	"amarbank/mockdata"
	"amarbank/module/account_v1/dto"
	service2 "amarbank/module/account_v1/service"
	"amarbank/module/account_v1/storage"
	storage2 "amarbank/module/user_v1/storage"
	app_error "amarbank/pkg/apperror"
	"amarbank/pkg/random"
	"amarbank/pkg/rest_response"
	"amarbank/pkg/valueobject"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.CreateAccountRequest

		if err := c.ShouldBind(&input); err != nil {
			panic(app_error.NewCustomError(err, "", ""))
		}

		store, _ := storage.NewAccountFileStorage(mockdata.DefaultPathToAccount, os.ModeAppend)
		userStore, _ := storage2.NewUSerFileStorage(mockdata.DefaultPathToUser, os.ModeAppend)
		pinGen := valueobject.NewPIN(random.Random(6), random.Random(4))

		service := service2.NewCreateAccountService(store, userStore, pinGen)

		if err := service.CreateAccount(c.Request.Context(), &input); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, rest_response.SimpleSuccessResponse(true))
	}
}
