package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	usermodel "go-api/module/user/model"
	userstorage "go-api/module/user/storage"
	userbusiness "go-api/module/user/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleSuccessResponse(err))
			return
		}

		paging.Fulfill()

		var filter usermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleSuccessResponse(err))
			return
		}

		var data []usermodel.User

		store := userstorage.NewStore(db)
		business := userbusiness.NewListUserBusiness(store)

		data, err := business.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}