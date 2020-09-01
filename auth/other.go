package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SignInWithCustomToken allows you to exchange a custom Auth token for an ID and refresh token
func (a *Auth) SignInWithCustomToken(token string) (*User, error) {
	client := &http.Client{}

	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=" + a.App.APIKey

	data := map[string]interface{}{"token": token, "returnSecureToken": true}
	reqdata, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqdata)))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respdata map[string]interface{}
	err = json.Unmarshal(body, &respdata)
	if err != nil {
		return nil, err
	}

	_, exists := respdata["error"]
	if exists {
		return nil, errors.New(respdata["error"].(map[string]interface{})["message"].(string))
	}

	timeLength, err := strconv.Atoi(respdata["expiresIn"].(string))
	if err != nil {
		return nil, err
	}

	return &User{
		IDToken:      respdata["idToken"].(string),
		RefreshToken: respdata["refreshToken"].(string),
		ExpiresIn:    time.Duration(timeLength),
		OtherData:    OtherData{},
	}, nil
}

// RefreshTokenToIDToken allows you to exchange a refresh token for an ID token
func (a *Auth) RefreshTokenToIDToken(refreshToken string) (*User, error) {
	client := &http.Client{}

	url := "https://securetoken.googleapis.com/v1/token?key=" + a.App.APIKey

	reqdata := "grant_type=refresh_token&refresh_token=" + refreshToken

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqdata)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respdata map[string]interface{}
	err = json.Unmarshal(body, &respdata)
	if err != nil {
		return nil, err
	}

	_, exists := respdata["error"]
	if exists {
		return nil, errors.New(respdata["error"].(map[string]interface{})["message"].(string))
	}

	timeLength, err := strconv.Atoi(respdata["expires_in"].(string))
	if err != nil {
		return nil, err
	}

	return &User{
		IDToken:      respdata["id_token"].(string),
		RefreshToken: respdata["refresh_token"].(string),
		ExpiresIn:    time.Duration(timeLength),
		OtherData: OtherData{
			TokenType: respdata["token_type"].(string),
			UserID:    respdata["user_id"].(string),
			ProjectID: respdata["project_id"].(string),
		},
	}, nil
}
