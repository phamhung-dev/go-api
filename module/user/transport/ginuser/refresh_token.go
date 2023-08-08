package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	userbusiness "go-api/module/user/business"
	userstorage "go-api/module/user/storage"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RefreshToken(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appctx.GetMainDBConnection()
		tokenProvider := appctx.GetTokenProvider()

		storage := userstorage.NewStorage(db)
		business := userbusiness.NewRefreshTokenBusiness(storage, tokenProvider)

		token, err := business.RefreshToken(c.Request.Context(), refreshToken)

		if err != nil {
			panic(err)
		}

		domain := os.Getenv("DOMAIN")

		c.SetCookie("refresh_token", token.RefreshToken, token.ExpiredIn, "/", domain, false, true)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
