package main

import (
	"context"
	"fmt"
	"main/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("error Occured")
	}
}
