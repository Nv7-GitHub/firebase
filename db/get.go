package db

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
