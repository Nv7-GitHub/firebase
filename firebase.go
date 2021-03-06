package firebase

import (
	"context"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// Firebase contains the data required for the web API
type Firebase struct {
	DatabaseURL       string
	APIKey            string
	ServiceAccount    []byte
	Config            *jwt.Config
	Token             *oauth2.Token
	HasServiceAccount bool
	Prefix            string
	Headers           map[string]string
}

// Refresh checks if oauth2 token needs to be refreshed, refreshes if needed
func (f *Firebase) Refresh() error {
	if f.HasServiceAccount && time.Since(f.Token.Expiry) > 0 {
		var err error
		f.Token, err = f.Config.TokenSource(context.Background()).Token()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateApp creates a firebase app without a service account
func CreateApp(DatabaseURL, APIKey string) *Firebase {
	if DatabaseURL[len(DatabaseURL)-1] != '/' {
		DatabaseURL += "/"
	}
	return &Firebase{
		DatabaseURL:       DatabaseURL,
		APIKey:            APIKey,
		ServiceAccount:    make([]byte, 0),
		Config:            new(jwt.Config),
		Token:             new(oauth2.Token),
		HasServiceAccount: false,
		Prefix:            "",
		Headers:           make(map[string]string, 0),
	}
}

// CreateAppWithServiceAccount creates a firebase app with a service account
func CreateAppWithServiceAccount(DatabaseURL, APIKey string, ServiceAccount []byte) (*Firebase, error) {
	if DatabaseURL[len(DatabaseURL)-1] != '/' {
		DatabaseURL += "/"
	}

	conf, err := google.JWTConfigFromJSON(ServiceAccount, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return new(Firebase), err
	}

	tok, err := conf.TokenSource(context.Background()).Token()
	if err != nil {
		return new(Firebase), err
	}

	return &Firebase{
		DatabaseURL:       DatabaseURL,
		APIKey:            APIKey,
		ServiceAccount:    ServiceAccount,
		Config:            conf,
		Token:             tok,
		HasServiceAccount: true,
		Prefix:            "",
		Headers:           make(map[string]string, 0),
	}, nil
}
