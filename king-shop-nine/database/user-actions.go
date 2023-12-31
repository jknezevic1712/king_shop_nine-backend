package database

import (
	"king-shop-nine/utils"
	"log"

	_ "github.com/lib/pq"
)

// Create user session
//
// @args userID string
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

// Update user session
//
// @args userID string
func updateUserSession(userID string) {
	conn := ConnectToDB()

	q := `
		UPDATE "Session"
		SET expires = $1, "sessionToken" = $2
		WHERE "userID" = $3
	`
	_, err := conn.Exec(q, utils.TimeAfter30Days(), utils.CreateJWTToken(), userID)
	if err != nil {
		log.Printf("updateUserSession: error updating user session, %v", err)
	}
}

// Sign up user
//
// @args newUser User
//
// @returns error
func SignUpUser(newUser utils.User) error {
	conn := ConnectToDB()

	q := `
		INSERT INTO "Users" (id, name, email, "accountCreatedAt", image)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := conn.Exec(q, newUser.ID, newUser.Name, newUser.Email, newUser.AccountCreatedAt, newUser.Image)
	if err != nil {
		log.Printf("SignUpUser: error while inserting a new user, %v\n", err)
		return err
	}

	createUserSession(newUser.ID)
	conn.Close()
	return nil
}

// Fetch all users
//
// @returns []User, error
func FetchUsers() ([]utils.User, error) {
	users := []utils.User{}
	conn := ConnectToDB()

	rows, err := conn.Query(`SELECT * FROM "Users"`)
	if err != nil {
		log.Printf("FetchUsers: error executing query to the DB, %v\n", err)
		return users, err
	}

	defer rows.Close()

	_, err = rows.Columns()
	if err != nil {
		log.Printf("FetchUsers: error while returning column names, %v\n", err)
		return users, err
	}

	for rows.Next() {
		user := utils.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccountCreatedAt, &user.Image); err != nil {
			log.Printf("FetchUsers: error while fetching users, %v\n", err)
			return users, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("FetchUsers: error while fetching users, %v\n", err)
		return users, err
	}

	conn.Close()
	return users, nil
}

// Fetch user by ID and optionally update user session if the `isLogin` argument is true
//
// @args userID string, isLogin bool
//
// @returns User, error
func FetchUserByID(userID string, isLogin bool) (utils.User, error) {
	user := utils.User{}
	conn := ConnectToDB()

	row := conn.QueryRow(`SELECT * FROM "Users" WHERE id = $1`, userID)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.AccountCreatedAt, &user.Image); err != nil {
		log.Printf("FetchUserByID: error while fetching user with id %v, %v\n", userID, err)
		return user, err
	}

	if isLogin {
		updateUserSession(user.ID)
	}

	conn.Close()
	return user, nil
}

// Login user via credentials
//
// @args userCredentials UserLoginCredentials
//
// returns (UserData, error)
func LoginUserViaCredentials(userCredentials utils.UserLoginCredentials) (utils.UserData, error) {
	userData := utils.UserData{}

	return userData, nil
}
