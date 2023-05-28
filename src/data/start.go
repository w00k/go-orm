package data

import (
	"w00k/orm/src/models"

	"gorm.io/gorm"
)

// create table User and add data
func Init(db *gorm.DB) {
	addCountries(db)
	addUsers(db)
}

func addUsers(db *gorm.DB) {
	// Create table for `models.User`
	db.AutoMigrate(&models.User{})

	// db.Migrator().CreateTable(&models.User{})
	db.Migrator().CreateIndex(&models.User{}, "Id")

	// You can add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	// Add data
	db.Create(&models.User{Name: "John Doe", Number: 500, CountryCode: "US"})
	db.Create(&models.User{Name: "Jane Doe", Number: 1000, CountryCode: "CA"})
	db.Create(&models.User{Name: "Mary Smith", Number: 3000, CountryCode: "AR"})
	db.Create(&models.User{Name: "Peter Jones", Number: 60, CountryCode: "CL"})
	db.Create(&models.User{Name: "Susan Brown", Number: 3000, CountryCode: "CL"})
	db.Create(&models.User{Name: "David Williams", Number: 90, CountryCode: "AR"})
	db.Create(&models.User{Name: "Emily Johnson", Number: 120, CountryCode: "CA"})
	db.Create(&models.User{Name: "Michael Garcia", Number: 230, CountryCode: "US"})
	db.Create(&models.User{Name: "Ashley Gonzalez", Number: 500, CountryCode: "CA"})
	db.Create(&models.User{Name: "Daniel Rodriguez", Number: 1000, CountryCode: "CL"})
}

func addCountries(db *gorm.DB) {
	// Create table for `models.Countries`
	db.AutoMigrate(&models.Country{})

	// db.Migrator().CreateTable(&models.User{})
	db.Migrator().CreateIndex(&models.Country{}, "Id")

	// You can add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Country{})

	countries := []models.Country{
		{Name: "United States", Code: "US", Capital: "Washington, D.C."},
		{Name: "Canada", Code: "CA", Capital: "Ottawa"},
		{Name: "Argentina", Code: "AR", Capital: "Buenos Aires"},
		{Name: "Chile", Code: "CL", Capital: "Santiago"},
	}

	db.Create(countries)
}
