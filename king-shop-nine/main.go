package main

import (
	"king-shop-nine/database"
)

func main() {
	database.CreateEmptyTables()
	database.PopulateUsersTableWithMockUsers()
}
