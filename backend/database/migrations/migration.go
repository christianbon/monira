package migrations

import (
	tipe "monira/backend/database/type"

	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&tipe.User{}, &tipe.Role{}, &tipe.Project{}, &tipe.Task{}, &tipe.Column{}, &tipe.History{}, &tipe.UserProject{})
}
