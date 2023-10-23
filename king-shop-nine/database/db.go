package database

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID               string
	Name             string
	Email            string
	AccountCreatedAt string
	Image            string
}

func readEnvVars() map[string]string {
	// Read() will try to read .env file in the directory that the function calling THIS function is located in
	envVars, err := godotenv.Read()
	if err != nil {
		log.Printf("Error loading .env file, info: %v", err.Error())
	}

	return envVars
}

// Establish connection to the DB
//
// @returns *sql.DB
func ConnectToDB() *sql.DB {
	envVars := readEnvVars()
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

// Adds user to the DB by passing `user` map as an argument
//
// @args newUser of User struct type
//
// @returns error
func AddUser(newUser User) error {
	conn := ConnectToDB()

	// * Execute query of inserting a newUser into user table
	insertData := `
        INSERT INTO "Users" (id, name, email, accountCreatedAt, image)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := conn.Exec(insertData, newUser.ID, newUser.Name, newUser.Email, newUser.AccountCreatedAt, newUser.Image)
	if err != nil {
		log.Printf("AddUser: error while inserting a new user, %v\n", err)
		return err
	}
	log.Println("AddUser: succesfully inserted a new user")

	conn.Close()
	return nil
}

// Fetch all users
//
// @returns []User
func FetchUsers() []User {
	var users []User
	conn := ConnectToDB()

	rows, err := conn.Query(`SELECT * FROM "Users"`)
	if err != nil {
		log.Printf("FetchUsers: error executing query to the DB, %v\n", err)
	}

	// * Defer closing rows so that any resources it holds will be released when the function exits.
	defer rows.Close()

	_, err = rows.Columns()
	if err != nil {
		log.Printf("FetchUsers: error while returning column names, %v\n", err)
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccountCreatedAt, &user.Image); err != nil {
			log.Printf("FetchUsers: error while fetching users, %v\n", err)
		}

		users = append(users, user)
		log.Printf("Current users: %v\n", users)
	}

	if err := rows.Err(); err != nil {
		log.Printf("FetchUsers: error while fetching users, %v\n", err)
	}

	conn.Close()
	return users
}

// Fetch user by ID
//
// @args userID of int64 type
//
// @returns User
func FetchUserByID(userID int64) User {
	var user User
	conn := ConnectToDB()

	row := conn.QueryRow(`SELECT * FROM "Users" WHERE id = $1`, userID)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.AccountCreatedAt, &user.Image); err != nil {
		log.Printf("FetchUserByID: error while fetching user with id %c, %v\n", userID, err)
	}

	log.Printf("FetchUserByID: fetched user => %v\n", user)
	conn.Close()
	return user
}
