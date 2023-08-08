package main

import (
	"fmt"
	"go-api/component/appctx"
	"go-api/component/tokenpvd"
	"go-api/component/uploadpvd"
	"go-api/db"
	"go-api/middleware"
	"go-api/migration"
	"go-api/module/upload/transport/ginupload"
	"go-api/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func main() {
	appCtx := appctx.NewAppContext(db.NewPostgresqlConnection(), tokenpvd.NewJWTProvider(), uploadpvd.NewCloudinaryProvider())

	migration.Migrate(appCtx)

	fmt.Println("Setting up routes")

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/register", ginuser.CreateUser(appCtx))
				auth.POST("/login", ginuser.Login(appCtx))
				auth.GET("/refresh-token", ginuser.RefreshToken(appCtx))
			}

			upload := v1.Group("/upload")
			{
				upload.POST("", ginupload.UploadFile(appCtx))
			}

			user := v1.Group("/users")
			{
				user.GET("/me", middleware.Authenticate(appCtx), ginuser.FindUser(appCtx))
				user.GET("", ginuser.ListUser(appCtx))
				user.PUT("/:id", ginuser.UpdateUser(appCtx))
				user.DELETE("/:id", ginuser.DeleteUser(appCtx))
			}
		}
	}

	fmt.Println("Setting up routes successfully")

	router.Run()
}
