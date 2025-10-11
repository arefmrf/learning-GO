package migration

import (
	"fmt"
	"log"
	itemModel "prj/internal/modules/item/model"
	"prj/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&itemModel.Item{})

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("migrate done.")
}
