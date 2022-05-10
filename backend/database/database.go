package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	migrations "backend/database/migrations"
)

func InitDB() {
	var err error

	//connect to postgres database
	dsn := "host=localhost user=postgres password=docker123# dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)

		panic("failed to connect database")
	}

	migrations.MigrationsDatabase(db)

}
