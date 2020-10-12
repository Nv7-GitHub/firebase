package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Set will set the data at a path to an array of bytes
func (db *Db) Set(path string, data []byte) error {
	db.App.Refresh()

	client := &http.Client{}

	url := db.App.Prefix + db.App.DatabaseURL + path + ".json"
	if db.App.HasServiceAccount {
		url += "?access_token=" + db.App.Token.AccessToken
	}

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	for k, v := range db.App.Headers {
		req.Header.Set(k, v)
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
		return errors.New(respdata["error"].(string))
	}

	return nil
}

// SetData will set the data at a path to an interface{}
func (db *Db) SetData(path string, data interface{}) error {
	setdata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return db.Set(path, setdata)
}
