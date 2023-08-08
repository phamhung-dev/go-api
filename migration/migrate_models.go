package migration

import (
	"fmt"
	"go-api/component/appctx"
	usermodel "go-api/module/user/model"
	"log"
)

func Migrate(appctx appctx.AppContext) {
	fmt.Println("Migrating...")

	db := appctx.GetMainDBConnection()

	if err := db.AutoMigrate(&usermodel.User{}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Migrating successfully")
	}
}
