package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	userbusiness "go-api/module/user/business"
	usermodel "go-api/module/user/model"
	userstorage "go-api/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := userstorage.NewStorage(db)
		business := userbusiness.NewCreateUserBusiness(storage)

		response, err := business.CreateUser(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
