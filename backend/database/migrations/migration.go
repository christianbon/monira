package migrations

import (
	model "backend/database/models"

	"gorm.io/gorm"
)

func MigrationsDatabase(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Project{}, &model.Task{}, &model.Column{}, &model.History{}, &model.UserProject{})
}
