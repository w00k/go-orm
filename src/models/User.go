package models

type User struct {
	Id     uint `gorm:"primaryKey"` //index
	Name   string
	Number int
}
