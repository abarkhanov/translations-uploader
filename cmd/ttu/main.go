package main

import (
	"log"
	"os"

	app "github.com/abarkhanov/ttu/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
