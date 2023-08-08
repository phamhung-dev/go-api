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

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		var filter usermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var response []usermodel.User

		storage := userstorage.NewStorage(db)
		business := userbusiness.NewListUserBusiness(storage)

		response, err := business.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(response, paging, filter))
	}
}
