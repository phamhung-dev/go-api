package migration

import (
	"fmt"
	"go-api/config/db"
	usermodel "go-api/module/user/model"
	"log"
)

func init() {
	fmt.Println("Migrating models...")

	if err := db.GetConnection().AutoMigrate(&usermodel.User{}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Migrated models successfully!")
	}
}
