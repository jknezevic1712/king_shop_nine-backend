package database

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID               int64
	Name             string
	Age              int32
	DateOfBirth      string
	Email            string
	AccountCreatedAt string
}

func readEnvVars() map[string]string {
	// Read() will try to read .env file in the directory that the function calling THIS function is located in
	envVars, err := godotenv.Read()
	if err != nil {
		log.Printf("Error loading .env file, info: %v", err.Error())
	}

	return envVars
}

// Establishes connection to the DB and returns `connection` and `error`
func connectToDB() *sql.DB {
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

	// * Populate DB with tables
	// err = populateDBWithTables()
	// if err != nil {
	// 	log.Printf("connectToDB: error while populating DB with tables, %v\n", err)
	// 	conn.Close()
	// }
	// log.Println("connectToDB: succesfully populated DB with tables")

	return conn
}

// Create tables in DB based on the queries defined
func CreateTables() error {
	conn := connectToDB()

	// ? Some table names are reserved so you need to wrap the table name in quotes
	createTable := `
		CREATE TABLE IF NOT EXISTS "users" (
				id SERIAL PRIMARY KEY,
				name TEXT NOT NULL,
				age INT NOT NULL,
				dateOfBirth TEXT NOT NULL,
				email TEXT NOT NULL,
				accountCreatedAt TEXT NOT NULL,
				UNIQUE (email)
		);
	`

	// * Execute query of creating table
	_, err := conn.Exec(createTable)
	if err != nil {
		log.Printf("PopulateDBWithTables: error while populating database, %v\n", err)
	}

	log.Println("PopulateDBWithTables: OK")
	return nil
}

// Adds user to the DB by passing `user` map as an argument, return error if it exists
func AddUser(newUser User) error {
	conn := connectToDB()

	// * Execute query of inserting a newUser into user table
	insertData := `
        INSERT INTO "users" (name, age, dateOfBirth, email, accountCreatedAt)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := conn.Exec(insertData, newUser.Name, newUser.Age, newUser.DateOfBirth, newUser.Email, newUser.AccountCreatedAt)
	if err != nil {
		log.Printf("AddUser: error while inserting a new user, %v\n", err)
		return err
	}
	log.Println("AddUser: succesfully inserted a new user")

	conn.Close()

	return nil
}

// Fetches all users from the DB, returns `User` array and error if it exists
func FetchUsers() []User {
	var users []User
	conn := connectToDB()

	rows, err := conn.Query("SELECT id, name, age, dateofbirth, email, accountcreatedat FROM users")
	if err != nil {
		log.Printf("FetchUsers: error executing query to the DB, %v\n", err)
	}

	// * Defer closing rows so that any resources it holds will be released when the function exits.
	defer rows.Close()

	_, err = rows.Columns()
	if err != nil {
		log.Printf("FetchUsers: error while returning column names, %v\n", err)
	}
	// log.Printf("FetchUsers: adsadsdsadas, %v\n", dataArr)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.DateOfBirth, &user.Email, &user.AccountCreatedAt); err != nil {
			log.Printf("FetchUsers: error while fetching users, %v\n", err)
		}

		users = append(users, user)
		log.Printf("Current users: %v\n", users)
	}

	if err := rows.Err(); err != nil {
		log.Printf("FetchUsers: error while fetching users, %v\n", err)
	}

	// for rows.Next() {
	// 	var ver string
	// 	rows.Scan(&ver)
	// 	log.Println(ver)
	// }

	conn.Close()

	return users
}

// Fetch user from the DB with the provided `id`, returns `User` map and error if it exists
func FetchUserByID(id int64) User {
	var user User
	conn := connectToDB()

	row := conn.QueryRow("SELECT * FROM users WHERE id = $1", id)

	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.DateOfBirth, &user.Email, &user.AccountCreatedAt); err != nil {
		log.Printf("FetchUserByID: error while fetching user with id %c, %v\n", id, err)
	}

	log.Printf("FetchUserByID: fetched user => %v\n", user)
	conn.Close()

	return user
}
