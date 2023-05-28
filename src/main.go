package main

import (
	"fmt"
	"w00k/orm/src/connection"
	"w00k/orm/src/data"
	"w00k/orm/src/query"
)

func main() {

	fmt.Println("start")

	db := connection.Conn()

	data.Init(db)

	query.GetFirst(db)
	query.GetLast(db)

	// Get first 10 Users
	query.GetUserWhereIdIsValue(db, 10)

	// Get Users with Ids in {1, 2, 3}
	query.GetUserWhereIdsAreValue(db, []int{3, 4, 5})

	fmt.Println("done")

}
