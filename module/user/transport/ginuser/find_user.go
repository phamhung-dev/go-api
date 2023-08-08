package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(response))
	}
}
