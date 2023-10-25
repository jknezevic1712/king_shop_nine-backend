package main

import "king-shop-nine/database"

func main() {
	database.CreateEmptyTables()
	database.PopulateUsersTableWithMockData()
	database.PopulateProductsTableWithMockData()
}
