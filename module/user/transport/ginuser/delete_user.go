package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	userbusiness "go-api/module/user/business"
	userstorage "go-api/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := common.DecodeUID(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := userstorage.NewStorage(db)
		business := userbusiness.NewDeleteUserBusiness(storage)

		if err := business.DeleteUser(c.Request.Context(), id); err != nil {
			panic(err)
		}

		response := "delete user successfully"

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
