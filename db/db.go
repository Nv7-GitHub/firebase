package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

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

// Get will get an array of bytes of the data at a certain path.
func (db *Db) Get(path string) ([]byte, error) {
	db.App.Refresh()

	url := db.App.DatabaseURL + path + ".json"
	if db.App.HasServiceAccount {
		url += "?access_token=" + db.App.Token.AccessToken
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetData will get an interface{} of the data at a certain path.
func (db *Db) GetData(path string) (interface{}, error) {
	getdata, err := db.Get(path)
	if err != nil {
		return nil, err
	}

	var data interface{}
	json.Unmarshal(getdata, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Set will set the data at a path to an array of bytes
func (db *Db) Set(path string, data []byte) error {
	db.App.Refresh()

	client := &http.Client{}

	url := db.App.DatabaseURL + path + ".json"
	if db.App.HasServiceAccount {
		url += "?access_token=" + db.App.Token.AccessToken
	}

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(data)))
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
	json.Unmarshal(body, &respdata)
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
