package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() (db *gorm.DB) {
	dsn := "host=localhost user=postgres password=postgres dbname=orm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return
}
