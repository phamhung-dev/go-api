package ginuser

import (
	"go-api/common"
	"go-api/component/appctx"
	userbusiness "go-api/module/user/business"
	usermodel "go-api/module/user/model"
	userstorage "go-api/module/user/storage"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Login(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appctx.GetMainDBConnection()
		tokenProvider := appctx.GetTokenProvider()

		storage := userstorage.NewStorage(db)
		business := userbusiness.NewLoginBusiness(storage, tokenProvider)

		token, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		domain := os.Getenv("DOMAIN")

		c.SetCookie("refresh_token", token.RefreshToken, token.ExpiredIn, "/", domain, false, true)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
