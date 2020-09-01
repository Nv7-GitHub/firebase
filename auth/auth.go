package auth

import (
	"time"

	"github.com/Nv7-Github/firebase"
)

// Auth has all the functions for authentication
type Auth struct {
	App *firebase.Firebase
}

// User has the response data of Auth requests
type User struct {
	IDToken      string
	RefreshToken string
	ExpiresIn    time.Duration

	OtherData OtherData
}

// OtherData contains all possible other data
type OtherData struct {
	Email      string
	LocalID    string
	TokenType  string
	UserID     string
	ProjectID  string
	Registered bool
}

// CreateAuth creates an Auth struct with the supplied app
func CreateAuth(app *firebase.Firebase) *Auth {
	return &Auth{app}
}
