package db

import (
	"github.com/Nv7-Github/firebase"
)

// Db has all the functions to interact with a firebase realtime database
type Db struct {
	App *firebase.Firebase
}

// CreateDatabase will return a database initialized with the app
func CreateDatabase(app *firebase.Firebase) *Db {
	return &Db{app}
}
