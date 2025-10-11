package bootstrap

import (
	"prj/internal/database/migration"
	"prj/pkg/config"
	"prj/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	migration.Migrate()
}
