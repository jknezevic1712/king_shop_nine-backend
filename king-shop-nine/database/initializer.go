package database

import (
	"database/sql"
	"log"

	"king-shop-nine/utils"
)

// Establish connection to the DB
//
// @returns *sql.DB
func ConnectToDB() *sql.DB {
	envVars := utils.ReadEnvVars()
	connection_string := envVars["DB_CONN_STRING"]

	// * Open connection to the DB
	conn, err := sql.Open("postgres", connection_string)
	if err != nil {
		log.Printf("connectToDB: error while opening connection to the DB, %v\n", err)
	}

	// * Ping DB so we know we have a succesful connection
	pingErr := conn.Ping()
	if pingErr != nil {
		log.Printf("connectToDB: ping unsuccesful, %v\n", pingErr)
		conn.Close()
	}
	log.Println("connectToDB: connected to db")

	return conn
}

// Create empty tables in DB based on the queries defined
//
// @returns error
func CreateEmptyTables() error {
	conn := ConnectToDB()

	// ? Some table names are reserved so you need to wrap the table name in quotes
	UsersTable := `
		CREATE TABLE IF NOT EXISTS "Users" (
				id TEXT NOT NULL PRIMARY KEY,
				name TEXT NOT NULL,
				email TEXT NOT NULL,
				"accountCreatedAt" TEXT NOT NULL,
				image TEXT NOT NULL,
				UNIQUE (email)
		);
	`

	SessionTable := `
		CREATE TABLE IF NOT EXISTS "Session" (
				id SERIAL PRIMARY KEY,
				"sessionToken" TEXT NOT NULL,
				"userID" TEXT NOT NULL,
				expires TEXT NOT NULL,
				UNIQUE ("userID")
		);
	`

	ProductsTable := `
		CREATE TABLE IF NOT EXISTS "Products" (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			"shortDescription" TEXT NOT NULL,
			description TEXT NOT NULL,
			category TEXT NOT NULL,
			subcategory TEXT NOT NULL,
			image TEXT NOT NULL,
			"dateAdded" TEXT NOT NULL,
			"ratingRate" REAL NOT NULL,
			"ratingCount" INTEGER NOT NULL,
			UNIQUE (title)
		);
	`

	// * Execute query of creating table
	_, err := conn.Exec(UsersTable)
	if err != nil {
		log.Printf("CreateEmptyTables: error while creating `Users` table, %v\n", err)
		return err
	}

	_, err = conn.Exec(SessionTable)
	if err != nil {
		log.Printf("CreateEmptyTables: error while creating `Session` table, %v\n", err)
		return err
	}

	_, err = conn.Exec(ProductsTable)
	if err != nil {
		log.Printf("CreateEmptyTables: error while creating `Products` table, %v\n", err)
		return err
	}

	log.Println("CreateEmptyTables: OK")
	return nil
}

// Populate `Users` table with mock data
//
// @returns error
func PopulateUsersTableWithMockData() error {
	for i := 0; i < len(utils.MockUsers); i++ {
		err := AddUser(utils.MockUsers[i])
		if err != nil {
			log.Printf("PopulateUsersTableWithMockData: error while populating `Users` table, %v\n", err)
			return err
		}
	}

	log.Println("PopulateUsersTableWithMockData: OK")
	return nil
}

// Populate `Products` table with mock data
//
// @returns error
func PopulateProductsTableWithMockData() error {
	for i := 0; i < len(utils.MockProducts); i++ {
		err := AddProduct(utils.MockProducts[i])
		if err != nil {
			log.Printf("PopulateProductsTableWithMockData: error while populating `Products` table, %v\n", err)
			return err
		}
	}

	log.Println("PopulateProductsTableWithMockData: OK")
	return nil
}
