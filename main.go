package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const URI = "https://dummy-project-d2032.firebaseio.com/"

func connect() (*firebase.App, error) {
	conf := &firebase.Config{
		DatabaseURL: URI,
	}
	opt := option.WithCredentialsFile("/Users/yudi/Downloads/dummy-project-d2032-firebase-adminsdk-j2o4e-59b02494d7.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, err
}

type User struct {
	DateOfBirth string `json:"date_of_birth,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}

func main() {
	ctx := context.Background()
	conn, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	client, err := conn.Database(ctx)
	if err != nil {
		fmt.Println(err)
	}

	ref := client.NewRef("user")

	// child := ref.Child("data")
	ref.Set(ctx, map[string]*User{
		"awesome": {
			DateOfBirth: "june 23, 1921",
			FullName:    "alan suryajana",
		},
		"gratettt": {
			DateOfBirth: "december 01, 1920",
			FullName:    "greatt so clenup",
		},
	})

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	// ref := client.NewRef("users")
	// var data map[string]interface{}
	// if err := ref.Get(ctx, &data); err != nil {
	// 	log.Fatalln("Error reading from database:", err)
	// }
	// fmt.Println(data)
}
