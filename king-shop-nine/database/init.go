package database

import (
	"log"

	"king-shop-nine/utils"
)

var mockUsers = []User{
	{
		ID:               "1asda433sdaasd2",
		Name:             "Jakov K",
		Email:            "knezevic.jakov@gmail.com",
		AccountCreatedAt: utils.GetTimeNowMs(),
		Image:            "https://t2.gstatic.com/licensed-image?q=tbn:ANd9GcQdAnprsidzbOSZ4jI1SvcFeIEuFKwBLrILGo8tLCEA4ixMzfxUQfk6onBDhipea4sD",
	},
	{
		ID:               "2bvcv565vbbvc3a",
		Name:             "Patrick Jane",
		Email:            "pjane@mail.com",
		AccountCreatedAt: utils.GetTimeNowMs(),
		Image:            "https://media.wired.com/photos/593261cab8eb31692072f129/master/w_2560%2Cc_limit/85120553.jpg",
	},
	{
		ID:               "3zdfzdf34adadsa",
		Name:             "Kimball Cho",
		Email:            "kcho@mail.com",
		AccountCreatedAt: utils.GetTimeNowMs(),
		Image:            "https://cdn.britannica.com/89/149189-050-68D7613E/Bengal-tiger.jpg",
	},
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
				accountCreatedAt TEXT NOT NULL,
				image TEXT NOT NULL,
				UNIQUE (email)
		);
	`

	SessionTable := `
		CREATE TABLE IF NOT EXISTS "Session" (
				id TEXT NOT NULL PRIMARY KEY,
				sessionToken TEXT NOT NULL,
				userID TEXT NOT NULL,
				expires INT NOT NULL,
				UNIQUE (userID)
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
