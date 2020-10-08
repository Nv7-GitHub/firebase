package db

import (
	"encoding/json"
	"errors"

	"github.com/r3labs/sse"
)

// StreamEvent contains the data in a stream event sent by firebase
type StreamEvent struct {
	Path    string
	AbsPath string
	Data    interface{}
}

// Stream will call a handler on any changes to the database
func (db *Db) Stream(path string, handler func(StreamEvent, error)) {
	db.App.Refresh()

	url := db.App.DatabaseURL + path + ".json"
	if db.App.HasServiceAccount {
		url += "?access_token=" + db.App.Token.AccessToken
	}

	client := sse.NewClient(url)

	client.SubscribeRaw(func(msg *sse.Event) {
		var eventData map[string]interface{}
		err := json.Unmarshal(msg.Data, &eventData)
		if err != nil {
			handler(*new(StreamEvent), err)
			return
		}
		_, hasError := eventData["error"]
		if hasError {
			handler(*new(StreamEvent), errors.New(eventData["error"].(string)))
			return
		}
		handler(StreamEvent{
			Path:    eventData["path"].(string),
			AbsPath: path + eventData["path"].(string),
			Data:    eventData["data"],
		}, nil)
	})
}
