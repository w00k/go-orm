package data

import (
	"w00k/orm/src/models"

	"gorm.io/gorm"
)

// create table User and add data
func Init(db *gorm.DB) {

	// Create table for `models.User`
	db.AutoMigrate(&models.User{})

	// db.Migrator().CreateTable(&models.User{})
	db.Migrator().CreateIndex(&models.User{}, "Id")

	// You can add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})

	// Add data
	db.Create(&models.User{Name: "John Doe", Number: 500})
	db.Create(&models.User{Name: "Jane Doe", Number: 1000})
	db.Create(&models.User{Name: "Mary Smith", Number: 3000})
	db.Create(&models.User{Name: "Peter Jones", Number: 60})
	db.Create(&models.User{Name: "Susan Brown", Number: 80})
	db.Create(&models.User{Name: "David Williams", Number: 90})
	db.Create(&models.User{Name: "Emily Johnson", Number: 120})
	db.Create(&models.User{Name: "Michael Garcia", Number: 230})
	db.Create(&models.User{Name: "Ashley Gonzalez", Number: 560})
	db.Create(&models.User{Name: "Daniel Rodriguez", Number: 980})
}
