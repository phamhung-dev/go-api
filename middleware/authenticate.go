package middleware

import (
	"errors"
	"go-api/common"
	"go-api/component/appctx"
	userbusiness "go-api/module/user/business"
	usermodel "go-api/module/user/model"
	userstorage "go-api/module/user/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeader(authorization string) (string, error) {
	parts := strings.Split(authorization, " ")

	// "Authorization": "Bearer <token>"
	if parts[0] != "Bearer" || len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader
	}

	return parts[1], nil
}

func Authenticate(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		tokenProvider := appctx.GetTokenProvider()

		payload, err := tokenProvider.ValidateAccessToken(token)

		if err == common.ErrRecordNotFound {
			panic(common.ErrEntityNotFound(usermodel.EntityName, err))
		}

		if err != nil {
			panic(common.ErrDB(err))
		}

		db := appctx.GetMainDBConnection()
		storage := userstorage.NewStorage(db)
		business := userbusiness.NewFindUserBussiness(storage)

		user, err := business.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.IsLocked {
			panic(common.ErrNoPermission(usermodel.ErrUserIsLocked))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

var (
	ErrWrongAuthHeader = common.NewCustomErrorResponse(
		errors.New("wrong authen header"),
		"wrong authen header",
		"ErrWrongAuthHeader",
	)
)
