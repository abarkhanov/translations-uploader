package main

import (
	"log"
	"os"

	app "github.com/abarkhanov/ttu/internal/app"
)

func main() {
	a := app.New()
	err := a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
