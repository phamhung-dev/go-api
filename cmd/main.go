package main

import (
	"fmt"
	"go-api/component/appctx"
	"go-api/config/db"
	_ "go-api/migration"
	"go-api/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func main() {

	appCtx := appctx.NewAppContext(db.GetConnection())

	fmt.Println("Setting up routes")

	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/users")
			{
				user.GET("", ginuser.ListUser(appCtx))
				user.POST("", ginuser.CreateUser(appCtx))
				user.DELETE("/:id", ginuser.DeleteUser(appCtx))
			}
		}
	}

	fmt.Println("Setting up routes successfully")

	router.Run()
}
