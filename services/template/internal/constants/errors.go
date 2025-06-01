package constants

import "fmt"

// business errors -> to frontend
var (
	ErrUserAlreadyExists = fmt.Errorf("user with such email already exists")
	ErrUserNotFound      = fmt.Errorf("user not found")
)
