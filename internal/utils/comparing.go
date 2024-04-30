package utils

import (
	"errors"
	"strconv"
	"strings"
)

func IsYearRangeLessOrEqual(yearRange string, referenceYearRange string) bool {
	// Split the year ranges into start and end years
	currentStartYear, currentEndYear, err := ParseYearRange(yearRange)
	if err != nil {
		// Handle error, for example, return false or log the error
		return false
	}
	referenceStartYear, referenceEndYear, err := ParseYearRange(referenceYearRange)
	if err != nil {
		// Handle error
		return false
	}

	// Compare start years
	if currentStartYear < referenceStartYear {
		return true
	} else if currentStartYear > referenceStartYear {
		return false
	}

	// Start years are equal, compare end years
	return currentEndYear <= referenceEndYear
}

// Helper function to parse year range string (YYYY/YYYY) into separate years
func ParseYearRange(yearRange string) (int, int, error) {
	// Use string splitting and conversion
	parts := strings.Split(yearRange, "/")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid year range format")
	}
	startYear, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	endYear, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return startYear, endYear, nil
}
