package models

type Country struct {
	Id      uint   `gorm:"primaryKey"`  //index
	Code    string `gorm:"uniqueIndex"` //index
	Name    string
	Capital string
}
