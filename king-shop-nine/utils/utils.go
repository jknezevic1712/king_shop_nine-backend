package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Read .env file located in the directory of the function that calls this function
//
// @returns map[string]string of environment variables
func ReadEnvVars() map[string]string {
	// Read() will try to read .env file in the directory that the function calling THIS function is located in
	envVars, err := godotenv.Read()
	if err != nil {
		log.Printf("Error loading .env file, info: %v", err)
	}

	return envVars
}

// Format a date to according to a predefined format
//
// @args dateToFormat string, e.g. "Jan-25-84"
//
// @returns string
func DateFormatter(dateToFormat string) string {
	const parseLayout = "Jan-02-06"
	const formatLayout = "02-Jan-2006"

	parsedDate, err := time.Parse(parseLayout, dateToFormat)
	if err != nil {
		return ""
	}

	formattedDate := parsedDate.Format(formatLayout)
	return formattedDate
}

// Convert int64 to string
//
// @returns string
func IntToString(number int64) string {
	numberInt := int(number)
	if int64(numberInt) != number {
		log.Println("IntToString: overflow error while converting int64 to int")
		return ""
	}

	return strconv.Itoa(numberInt)
}

// Convert string to int
//
// @returns int
func StringToInt(text string) (int, error) {
	number, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Error during conversion")
		return 0, err
	}

	fmt.Println(number)
	return number, nil
}

// Get time after 30 days in milliseconds
//
// @returns string
func TimeAfter30Days() string {
	timeAfter30Days := time.Now().UnixMilli() + 2592000000
	return IntToString(timeAfter30Days)
}

// Create JWT token that returns signed string you can save to DB in a `sessionToken` field
//
// @returns string
func CreateJWTToken() string {
	envVars := ReadEnvVars()

	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(envVars["JWT_KEY"])
	t = jwt.New(jwt.SigningMethodHS256)
	s, _ = t.SignedString(key)

	return s
}

// Validate JWT token according to the `JWT_KEY` env variable
func ValidateJWTToken(tokenString string) {
	envVars := ReadEnvVars()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(envVars["JWT_KEY"]), nil
	})

	if token.Valid {
		fmt.Println("OK")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Printf("Token malformed => %v\n", err)
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		fmt.Printf("Invalid signature => %v\n", err)
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		fmt.Printf("Token is either expired or not active yet => %v\n", err)
	} else {
		fmt.Printf("Couldn't handle this token => %v\n", err)
	}
}

// Convert data to stringified JSON
//
// @returns string
func ToJSON(data any) string {
	res, err := json.Marshal(data)
	if err != nil {
		log.Printf("ToJSON: error while converting data to json, %v\n", err)
	}
	return string(res)
}
