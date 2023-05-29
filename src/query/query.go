package query

import (
	"fmt"
	"w00k/orm/src/models"

	"gorm.io/gorm"
)

// get the first element
func GetFirst(db *gorm.DB) (user models.User) {
	result := db.First(&user)

	fmt.Printf("GetFirst user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil

	return
}

// get the last element
func GetLast(db *gorm.DB) (user models.User) {
	result := db.Last(&user)

	fmt.Printf("GetLast user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil
	return
}

// get the element in position "value"
func GetUserWhereIdIsValue(db *gorm.DB, value int) (user models.User) {
	result := db.First(&user, value)

	fmt.Printf("GetUserWhereIdIsValue user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil

	return
}

// get the element in "value" array example {3, 4, 5}
func GetUserWhereIdsAreValue(db *gorm.DB, value []int) (users []models.User) {
	result := db.Find(&users, value)

	for _, user := range users {
		fmt.Printf("GetUserWhereIdsAreValue user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	}

	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil

	return
}

// query with where of one table
func GetUserById(db *gorm.DB, id int) (user models.User) {

	result := db.First(&user, "id = ?", id)

	fmt.Printf("GetUserById user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil

	return
}

// get elements in order desc
func GetDiferentNumber(db *gorm.DB) (users []models.User) {
	result := db.Distinct("number").Order("number desc").Find(&users)

	for _, user := range users {
		fmt.Printf("GetDiferentNumber user ::: Number %d\n", user.Number)
	}

	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil
	return
}

// get join tables in new struct
func GetUserJoinCountry(db *gorm.DB) (userJoinCountries []models.UserJoinCountry) {
	result := db.Model(&models.User{}).Select("users.id, users.name, users.number, countries.name as country").Joins("join countries on countries.code = users.country_code").Scan(&userJoinCountries)

	for _, userJoinCountry := range userJoinCountries {
		fmt.Printf("userJoinCountries ::: %d %s %d %s\n", userJoinCountry.Id, userJoinCountry.Name, userJoinCountry.Number, userJoinCountry.Country)
	}

	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil
	return
}

// get join table in new struct with where
func GetUserJoinCountryWhereCountryIsCA(db *gorm.DB) (userJoinCountries []models.UserJoinCountry) {
	result := db.Model(&models.User{}).Select("users.id, users.name, users.number, countries.name as country").Joins("join countries on countries.code = users.country_code").Where("users.country_code = ?", "CA").Scan(&userJoinCountries)

	for _, userJoinCountry := range userJoinCountries {
		fmt.Printf("GetUserJoinCountryWhereCountryIsCA ::: %d %s %d %s\n", userJoinCountry.Id, userJoinCountry.Name, userJoinCountry.Number, userJoinCountry.Country)
	}

	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("error %d\n\n", result.Error)              // returns error or nil
	return
}
