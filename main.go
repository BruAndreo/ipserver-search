package main

import (
	"fmt"
	"ipserver-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting...")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
