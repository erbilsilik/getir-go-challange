package main

import (
	server "github.com/erbilsilik/getir-go-challange/api"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func main() {
	server.Run()
	// cmd.Run()
}
