package utils

import (
	"time"
)

func DateFormatter(startingDate string) string {
	const parseLayout = "Jan-02-06"
	const formatLayout = "02-Jan-2006"

	parsedDate, err := time.Parse(parseLayout, startingDate)
	if err != nil {
		return ""
	}

	formattedDate := parsedDate.Format(formatLayout)
	return formattedDate
}
