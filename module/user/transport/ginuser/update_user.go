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

func UpdateUser(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()

		id, err := common.DecodeUID(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := userstorage.NewStorage(db)
		business := userbusiness.NewUpdateUserBusiness(storage)

		response, err := business.UpdateUser(c.Request.Context(), id, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
