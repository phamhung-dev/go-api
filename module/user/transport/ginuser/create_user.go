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

		var data usermodel.UserCreateRequest

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(err))
			return
		}

		store := userstorage.NewStore(db)
		business := userbusiness.NewCreateUserBusiness(store)

		response, err := business.CreateUser(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
