package main

import (
	"log"

	"github.com/joho/godotenv"
	"ticket.narindo.com/delivery"
)

func main() {
	// Load config files
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatalln(err)
	}

	// Run the server
	delivery.Server().Run()
}
