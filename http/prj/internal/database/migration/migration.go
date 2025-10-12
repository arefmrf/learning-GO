package migration

import (
	"fmt"
	"log"
	itemModel "prj/internal/modules/item/model"
	userModel "prj/internal/modules/user/model"
	"prj/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&itemModel.Item{}, userModel.User{})

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("migrate done.")
}
