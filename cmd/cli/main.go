package main

import (
	"ipserver-search/internal/app/cli"
	"log"
	"os"
)

func main() {
	application := cli.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
