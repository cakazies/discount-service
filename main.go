package main

import (
	"discount-service/cmd"
	"log"

	"github.com/alexsasharegan/dotenv"
)

func main() {

	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal("load env error", err)
	}

	cmd.Run()
}
