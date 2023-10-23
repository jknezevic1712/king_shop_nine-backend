package database

import (
	"database/sql"
	"log"
	"time"

	"king-shop-nine/utils"
)

var mockUsers = []User{
	{
		ID:               "1asda433sdaasd2",
		Name:             "Jakov K",
		Email:            "knezevic.jakov@gmail.com",
		AccountCreatedAt: utils.IntToString(time.Now().UnixMilli()),
		Image:            "https://t2.gstatic.com/licensed-image?q=tbn:ANd9GcQdAnprsidzbOSZ4jI1SvcFeIEuFKwBLrILGo8tLCEA4ixMzfxUQfk6onBDhipea4sD",
	},
	{
		ID:               "2bvcv565vbbvc3a",
		Name:             "Patrick Jane",
		Email:            "pjane@mail.com",
		AccountCreatedAt: utils.IntToString(time.Now().UnixMilli()),
		Image:            "https://media.wired.com/photos/593261cab8eb31692072f129/master/w_2560%2Cc_limit/85120553.jpg",
	},
	{
		ID:               "3zdfzdf34adadsa",
		Name:             "Kimball Cho",
		Email:            "kcho@mail.com",
		AccountCreatedAt: utils.IntToString(time.Now().UnixMilli()),
		Image:            "https://cdn.britannica.com/89/149189-050-68D7613E/Bengal-tiger.jpg",
	},
}

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

	log.Println("CreateEmptyTables: OK")
	return nil
}

// Populate `Users` table with mock users data`
//
// @returns error
func PopulateUsersTableWithMockUsers() error {
	for i := 0; i < len(mockUsers); i++ {
		err := AddUser(mockUsers[i])
		if err != nil {
			log.Printf("PopulateUsersTableWithMockUsers: error while populating `Users` table, %v\n", err)
			return err
		}
	}

	log.Println("PopulateUsersTableWithMockUsers: OK")
	return nil
}
