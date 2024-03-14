package main

import (
	"fmt"

	"github.com/juanhenaoparra/go-tting-started/app"
)

var (
	defaultPort = 8001
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		fmt.Println("error creating new app: ", err.Error())
		return
	}

	err = a.Start(defaultPort)
	if err != nil {
		fmt.Printf("error starting app on port%d: %s", defaultPort, err.Error())
	}
}
