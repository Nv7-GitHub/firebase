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

// SignUpWithEmailAndPassword lets you create a user with a specified email and password
func (a *Auth) SignUpWithEmailAndPassword(email string, password string) (*User, error) {
	client := &http.Client{}

	url := "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=" + a.App.APIKey

	data := map[string]interface{}{"email": email, "password": password, "returnSecureToken": true}
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
		OtherData: OtherData{
			Email:   respdata["email"].(string),
			LocalID: respdata["localId"].(string),
		},
	}, nil
}

// SignInWithEmailAndPassword lets you sign in a user with a specified email and password
func (a *Auth) SignInWithEmailAndPassword(email string, password string) (*User, error) {
	client := &http.Client{}

	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + a.App.APIKey

	data := map[string]interface{}{"email": email, "password": password, "returnSecureToken": true}
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
		OtherData: OtherData{
			Email:      respdata["email"].(string),
			LocalID:    respdata["localId"].(string),
			Registered: respdata["registered"].(bool),
		},
	}, nil
}

// ResetPassword allows you to send a password reset email to a user with a certain email
func (a *Auth) ResetPassword(email string) error {
	client := &http.Client{}

	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=" + a.App.APIKey

	data := map[string]interface{}{"email": email, "requestType": "PASSWORD_RESET"}
	reqdata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqdata)))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respdata map[string]interface{}
	err = json.Unmarshal(body, &respdata)
	if err != nil {
		return err
	}

	_, exists := respdata["error"]
	if exists {
		return errors.New(respdata["error"].(map[string]interface{})["message"].(string))
	}

	return nil
}
