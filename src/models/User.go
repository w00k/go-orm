package models

type User struct {
	Id          uint   `gorm:"primaryKey"`  //index
	Name        string `gorm:"uniqueIndex"` //index
	Number      int
	CountryCode string
}
