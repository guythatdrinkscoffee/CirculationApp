package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully loaded the .env file")
}

func main() {

}
