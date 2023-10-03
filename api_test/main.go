package main

import (
	"api_test/database"
	"fmt"
)

// ! Types defined in separate packages are different even if they share name and underlying structure.
// type User struct {
// 	ID          int64
// 	Name        string
// 	Age         int32
// 	DateOfBirth string
// }

func main() {
	// "Dec-17-98"

	// newDate := utils.DateFormatter("Dec-17-98")
	// if newDate == "" || err != nil {
	// 	fmt.Printf("main: error with creating a new date, %v", err)
	// 	return
	// }

	// fmt.Println("main: created a new date", newDate)

	// var user = database.User{
	// 	ID:               1,
	// 	Name:             "Jakov K",
	// 	Age:              24,
	// 	DateOfBirth:      utils.DateFormatter("Dec-17-98"),
	// 	Email:            "knezevic.jakov@gmail.com",
	// 	AccountCreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// }

	// err := database.AddUser(user)
	// if err != nil {
	// 	// fmt.Printf("main: unsuccesfully added the new user, %v\n", err)
	// 	log.Fatalf("main: error with query 'database.AddUser', %v", err)
	// }

	data := database.FetchUsers()
	fmt.Printf("main: fetched users from the DB, %v\n", data)
}
