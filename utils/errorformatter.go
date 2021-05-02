package errorformatter

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("you're already registered")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("invalid credentials")
	}

	return errors.New("invalid credentials")
}
