package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// App ...
type App struct {
	ID    string `json:"id"`
	Title string `json:"title,omitempty"`
}

// BiggerApp ...
type BiggerApp struct {
	App
	Author string `json:"author,omitempty"`
}

// Application ...
type Application struct {
	ID string `json:"id"`
}

// Org ...
type Org struct {
	Name string `json:"name"`
}

// AppWithOrg ...
type AppWithOrg struct {
	App
	Org
}

func main() {
	data := []byte(`
    {
        "id": "k34rAT4",
        "title": "My Awesome App"
    }
`)
	var app App
	//from json to entity
	err := json.Unmarshal(data, &app)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("From data json:", app)
	// reversed: from entity to json
	appByte, err := json.Marshal(app)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("From struc to json:", string(appByte))

	// with omited fields
	biggerApp1 := BiggerApp{app, "Me"}
	biggerApp2 := BiggerApp{App: app}

	biggerApp1JSON, err := json.Marshal(biggerApp1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("biggerApp1:", string(biggerApp1JSON))
	biggerApp2JSON, err := json.Marshal(biggerApp2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("biggerApp1:", string(biggerApp2JSON))

	//recieved from bytes
	var biggerAppJSON BiggerApp
	err = json.Unmarshal(data, &biggerAppJSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(biggerAppJSON)
	data = []byte(`
    {
        "id": "k34rAT4",
        "name": "My Awesome Org"
    }
`)

	var appWithOrg AppWithOrg
	err = json.Unmarshal(data, &appWithOrg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appWithOrg)

	byteJSON, err := json.Marshal(appWithOrg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteJSON))
}
