package main

import (
	"king-shop-nine/database"
	"king-shop-nine/utils"
	"log"
	"time"
)

// ! Types defined in separate packages are different even if they share name and underlying structure.
// type User struct {
// 	ID          int64
// 	Name        string
// 	Age         int32
// 	DateOfBirth string
// }

func main() {
	var user = database.User{
		ID:               1,
		Name:             "Jakov K",
		Age:              24,
		DateOfBirth:      utils.DateFormatter("Dec-17-98"),
		Email:            "knezevic.jakov@gmail.com",
		AccountCreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Creating tables
	database.CreateTables()

	// Adding a user
	addUserErr := database.AddUser(user)
	if addUserErr != nil {
		log.Printf("main: addUser ERROR, %v\n", addUserErr)
	} else {
		log.Printf("main: addUser OK\n")
	}

	// Fetching users
	fetchedUsers := database.FetchUsers()
	log.Printf("main: fetchedUsers => %v\n", fetchedUsers)

	// Fetching user by id
	fetchedUser := database.FetchUserByID(1)
	log.Printf("main: fetchUserByID => %v\n", fetchedUser)
}
