package main

import (
	"fmt"
	"sas-backend/app"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	a := &app.App{}
	a.Initialize()
	a.Run()
}
