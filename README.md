# Firebase
The firebase library supplies lightweight, easy-to-use bindings to the Firebase REST HTTP API. It was created because of a lack of all-in-one packages that allowed Authentication, and Database management.

# Installation
You can install the latest version of this library with
```bash
go get -v -u github.com/Nv7-Github/firebase
```

# Usage
The first thing to do is to create a `Firebase` object. This has all the dat required for using the Firebase project. You do this with the `CreateApp` and `CreateAppWithServiceAccount` methods. This is how to use the `CreateApp` method:
```go
package main

import "github.com/Nv7-Github/firebase"

func main() {
    app := firebase.CreateApp("https://[PROJECT_ID].firebaseio.com", "[API_KEY]")
}
```
And this is how the `CreateAppWithServiceAccount` method is used:
```go
package main

import "github.com/Nv7-Github/firebase"

var serviceAccount []byte = []byte(`
{
	// ServiceAccount here, get rid of this comment
}
`)

func main() {
    app, err := firebaseCreateAppWithServiceAccount("https://[PROJECT_ID].firebaseio.com", "[API_KEY]", serviceAccount)
    if err != nil {
        panic(err)
    }
}
```
## Realtime Database
The realtime database is the only one currently supported. It supports writing data to a path, reading data from a path, and streaming data from a path. To do this, first you need to create a Database object. In this example I am using the basic app, but an app with a service account works the same. 
```go
package main

import (
    "github.com/Nv7-Github/firebase"
    "github.com/Nv7-Github/firebase/db"
)

func main() {
    app := firebase.CreateApp("https://[PROJECT_ID].firebaseio.com", "[API_KEY]")
    database := db.CreateDatabase(app)
}
```
### Reading
To read from the database, you use the `Get` and `GetData` methods. To read the data from path `/learning`, you could use the `Get` method so that, if you know the type you can Unmarshal the data into the right type.
```go
var learningData map[string]bool
data, err := database.Get("/learning")
if err != nil {
    panic(err)
}
err := json.Unmarshal(data, &learningData)
if err != nil {
    panic(err)
}
// Now you can have a map[string]bool
```
Or, if you want to typecast it, you can do this:
```go
data, err := database.GetData("/learning")
if err != nil {
    panic(err)
}
learningData := make(map[string]bool, 0)
learningData["learning"] = data.(map[string]interface{})["learning"].(bool)
```

### Writing
To write to the database, you use the `Set` and `SetData` methods. To change the data at path `/learning`, to `{"learning": true}` you would marshal data and then set it:
```go
database.Set("/learning", []byte(`{"learning": true}`))
```
Or you could use `SetData`:

```go
database.SetData("/learning", map[string]bool{"learning": true})
```


# Q & A
## Why not use the firebase-admin-sdk?
The Firebase admin SDK, as said in the name, is an Admin SDK. It doesn't allow authenticating with passwords, only changing accounts. In addition, it is not very lightweight. On my laptop (Mac OS), as of when this was written, programs using Only the Firebase SDK were around 21.6 MB. On the other hand, programs written using this library are around 4.3 MB.
