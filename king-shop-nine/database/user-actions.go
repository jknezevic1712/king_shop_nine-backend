package database

import (
	"king-shop-nine/utils"
	"log"

	_ "github.com/lib/pq"
)

// Create user session
//
// @args userID *string
func createUserSession(userID string) {
	conn := ConnectToDB()

	q := `
		INSERT INTO "Session" (expires, "sessionToken", "userID")
		VALUES ($1, $2, $3)
	`
	_, err := conn.Exec(q, utils.TimeAfter30Days(), utils.CreateJWTToken(), userID)
	if err != nil {
		log.Printf("createUserSession: error creating user session, %v", err)
	}
}

// Add user to the DB
//
// @args newUser User
//
// @returns error
func AddUser(newUser utils.User) error {
	conn := ConnectToDB()

	q := `
		INSERT INTO "Users" (id, name, email, "accountCreatedAt", image)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := conn.Exec(q, &newUser.ID, &newUser.Name, &newUser.Email, &newUser.AccountCreatedAt, &newUser.Image)
	if err != nil {
		log.Printf("AddUser: error while inserting a new user, %v\n", err)
		return err
	}
	createUserSession(newUser.ID)

	log.Println("AddUser: succesfully inserted a new user")

	conn.Close()
	return nil
}

// Fetch all users
//
// @returns string
func FetchUsers() string {
	var users []utils.User
	conn := ConnectToDB()

	rows, err := conn.Query(`SELECT * FROM "Users"`)
	if err != nil {
		log.Printf("FetchUsers: error executing query to the DB, %v\n", err)
	}

	defer rows.Close()

	_, err = rows.Columns()
	if err != nil {
		log.Printf("FetchUsers: error while returning column names, %v\n", err)
	}

	for rows.Next() {
		var user utils.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccountCreatedAt, &user.Image); err != nil {
			log.Printf("FetchUsers: error while fetching users, %v\n", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("FetchUsers: error while fetching users, %v\n", err)
	}

	conn.Close()
	return utils.ToJSON(users)
}

// Fetch user by ID
//
// @args userID string
//
// @returns string
func FetchUserByID(userID string) string {
	var user utils.User
	conn := ConnectToDB()

	row := conn.QueryRow(`SELECT * FROM "Users" WHERE id = $1`, userID)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.AccountCreatedAt, &user.Image); err != nil {
		log.Printf("FetchUserByID: error while fetching user with id %v, %v\n", userID, err)
	}

	conn.Close()
	return utils.ToJSON(user)
}
