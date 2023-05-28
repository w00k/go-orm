package query

import (
	"fmt"
	"w00k/orm/src/models"

	"gorm.io/gorm"
)

func GetFirst(db *gorm.DB) (user models.User) {
	result := db.First(&user)

	fmt.Printf("GetFirst user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("count records %d\n\n", result.Error)      // returns error or nil

	return
}

func GetLast(db *gorm.DB) (user models.User) {
	result := db.Last(&user)
	fmt.Printf("GetLast user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("count records %d\n\n", result.Error)      // returns error or nil
	return
}

func GetUserWhereIdIsValue(db *gorm.DB, value int) (user models.User) {
	result := db.First(&user, value)

	fmt.Printf("GetUserWhereIdIsValue user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("count records %d\n\n", result.Error)      // returns error or nil

	return
}

func GetUserWhereIdsAreValue(db *gorm.DB, value []int) (users []models.User) {
	result := db.Find(&users, value)

	for _, user := range users {
		fmt.Printf("GetUserWhereIdsAreValue user ::: Id %d ,Name %s ,Number %d\n", user.Id, user.Name, user.Number)
	}

	fmt.Printf("count records %d\n", result.RowsAffected) // returns count of records found
	fmt.Printf("count records %d\n\n", result.Error)      // returns error or nil

	return
}
