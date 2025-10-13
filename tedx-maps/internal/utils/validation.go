package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateCoordinates(lat, lon float64) error {
	if lat < -90 || lat > 90 || lon < -180 || lon > 180 {
		return errors.New("invalid coordinates")
	}
	return nil
}
