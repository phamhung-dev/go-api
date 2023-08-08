package middleware

import (
	"go-api/common"
	"go-api/component/appctx"

	"github.com/gin-gonic/gin"
)

func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr := common.ErrInternalServer(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
