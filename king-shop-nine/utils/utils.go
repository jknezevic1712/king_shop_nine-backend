package utils

import (
	"log"
	"strconv"
	"time"
)

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

// Convert current local time to a string
//
// @returns time now in ms as a string
func GetTimeNowMs() string {
	timeNowInt64 := time.Now().UnixMilli()
	timeNowMsInt := int(timeNowInt64)

	if int64(timeNowMsInt) != timeNowInt64 {
		log.Println("getTimeNowMs: overflow error while converting int64 to int")
		return ""
	}

	return strconv.Itoa(timeNowMsInt)
}
