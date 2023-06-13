package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	userbusiness "go-api/module/user/business"
	userstorage "go-api/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(err))
			return
		}

		store := userstorage.NewStore(db)
		business := userbusiness.NewDeleteUserBusiness(store)

		if err := business.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("delete user successfully"))
	}
}
