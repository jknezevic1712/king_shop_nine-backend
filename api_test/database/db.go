package database

import (
	"database/sql"
	"fmt"

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

// Establishes connection to the DB and returns `connection` and `error`
func connectToDB() *sql.DB {
	connection_string := "user=postgres password=go_api_test! host=db.szjzyutbvqnniaikanai.supabase.co port=5432 dbname=postgres"

	// * Open connection to the DB
	conn, err := sql.Open("postgres", connection_string)
	if err != nil {
		fmt.Printf("connectToDB: error while opening connection to the DB, %v\n", err)
		panic(err)
	}

	// * Ping DB so we know we have a succesful connection
	pingErr := conn.Ping()
	if pingErr != nil {
		fmt.Printf("connectToDB: ping unsuccesful, %v\n", pingErr)
		conn.Close()
		panic(err)
	}
	fmt.Println("connectToDB: connected to db")

	// * Populate DB with tables
	err = populateDBWithTables(conn)
	if err != nil {
		fmt.Printf("connectToDB: error while populating DB with tables, %v\n", err)
		conn.Close()
		panic(err)
	}
	fmt.Println("connectToDB: succesfully populated DB with tables")

	return conn
}

// Populates DB with empty tables
func populateDBWithTables(conn *sql.DB) error {
	// ? Some table names are reserved so you need to wrap the table name in quotes
	createTable := `
		CREATE TABLE IF NOT EXISTS "user" (
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
		panic(err)
	}
	return nil
}

// Adds user to the DB by passing `user` map as an argument, return error if it exists
func AddUser(newUser User) error {
	conn := connectToDB()

	// * Execute query of inserting a newUser into user table
	insertData := `
        INSERT INTO "user" (name, age, dateOfBirth, email, accountCreatedAt)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := conn.Exec(insertData, newUser.Name, newUser.Age, newUser.DateOfBirth, newUser.Email, newUser.AccountCreatedAt)
	if err != nil {
		fmt.Printf("AddUser: error while inserting a new user, %v\n", err)
		return err
	}
	fmt.Println("AddUser: succesfully inserted a new user")

	conn.Close()

	return nil
}

// Fetches all users from the DB, returns `User` array and error if it exists
func FetchUsers() []User {
	var users []User
	conn := connectToDB()

	rows, err := conn.Query("SELECT id, name, age, dateofbirth, email. accountcreatedat FROM user")
	if err != nil {
		fmt.Printf("FetchUsers: error while querying the DB, %v\n", err)
		panic(err)
	}

	// * Defer closing rows so that any resources it holds will be released when the function exits.
	defer rows.Close()

	dataArr, err := rows.Columns()
	if err != nil {
		fmt.Printf("FetchUsers: error while querying the DB, %v\n", err)
		panic(err)
	}
	fmt.Printf("FetchUsers: adsadsdsadas, %v\n", dataArr)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.DateOfBirth, &user.Email, &user.AccountCreatedAt); err != nil {
			fmt.Printf("FetchUsers: error while fetching users, %v\n", err)
			panic(err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("FetchUsers: error while fetching users, %v\n", err)
		panic(err)
	}

	// for rows.Next() {
	// 	var ver string
	// 	rows.Scan(&ver)
	// 	fmt.Println(ver)
	// }

	conn.Close()

	return users
}

// Fetch user from the DB with the provided `id`, returns `User` map and error if it exists
// func GetUserByID(id int64) (User, error) {

// }
