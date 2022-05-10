package database

import (
	"time"

	"gorm.io/gorm"
)

func init() {}

type User struct {
	gorm.Model
	ID     uint
	Name   string
	RoleID uint
}

type Role struct {
	gorm.Model
	ID       uint
	RoleName string
	Level    uint
}

type Project struct {
	gorm.Model
	ID          uint
	ProjectName string
}

type Task struct {
	gorm.Model
	ID          uint
	TaskName    string
	UserID      uint
	ProjectID   uint
	Description string
}

type Column struct {
	gorm.Model
	ID         uint
	ColumnName string
}

type History struct {
	gorm.Model
	ID           uint
	TaskID       uint
	ColumnFromID uint
	ColumnToID   uint
	UserID       uint
	CreatedAt    time.Time
}

type UserProject struct {
	gorm.Model
	UserID    uint
	ProjectID uint
}
